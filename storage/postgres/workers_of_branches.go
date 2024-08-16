package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"users_service/pkg/helper"
	"users_service/pkg/logger"

	"github.com/jackc/pgx/v5/pgxpool"

	pb "users_service/genproto/users"
)

type workersOfBranchesRepo struct {
	db  *pgxpool.Pool
	log logger.ILogger
}

func NewWorkersOfBranchesRepo(db *pgxpool.Pool, log logger.ILogger) *workersOfBranchesRepo {
	return &workersOfBranchesRepo{
		db:  db,
		log: log,
	}
}

func (w *workersOfBranchesRepo) Create(ctx context.Context, request *pb.CreateWorker) (*pb.WorkerId, error) {

	var (
		workerId string
		query    string
		err      error
		userId   string
		timeNow  = time.Now()
	)

	tx, err := w.db.Begin(ctx)
	if err != nil {
		w.log.Error("error while creating transaction in storage layer", logger.Error(err))
		return &pb.WorkerId{}, err
	}
	defer tx.Commit(ctx)

	query = `
		insert into users (
			phone_number,
			full_name,
			birthday,
			user_role,
			created_at
		) values ($1, $2, $3, $4, $5) returning id `

	if err := tx.QueryRow(ctx, query,
		request.GetPhoneNumber(),
		request.GetFullName(),
		request.GetBirthday(),
		request.GetUserRole(),
		timeNow,
	).Scan(&userId); err != nil {
		tx.Rollback(ctx)
		w.log.Error("error while creating workers in users table in storage layer", logger.Error(err))
		return &pb.WorkerId{}, err
	}

	query = `
		insert into workers_of_branches (
			user_id,
			branch_id,
			created_at
		) values ($1, $2, $3) returning worker_id `

	if err := tx.QueryRow(ctx, query,
		userId,
		request.GetBranchId(),
		timeNow,
	).Scan(&workerId); err != nil {
		tx.Rollback(ctx)
		w.log.Error("error while creating workers in workers_of_branches table in storage layer", logger.Error(err))
		return &pb.WorkerId{}, err
	}

	return &pb.WorkerId{WorkerId: workerId}, nil
}

func (w *workersOfBranchesRepo) GetByUUID(ctx context.Context, request *pb.PrimaryKey) (*pb.Worker, string, error) {

	var (
		worker    pb.Worker
		query     string
		err       error
		createdAt time.Time
		updatedAt sql.NullTime
		branchId  string
	)

	query = `
		select 
			user_id,
			worker_id,
			branch_id,
			created_at,
			updated_at
		from 
			workers_of_branches
		where
			id = $1 and
			deleted_at is null `

	if err = w.db.QueryRow(ctx, query, request.GetId()).Scan(
		&worker.Id,
		&worker.WorkerId,
		&branchId,
		&createdAt,
		&updatedAt,
	); err != nil {
		w.log.Error("error while getting workers from workers_of_branches table in storage layer", logger.Error(err))
		return &pb.Worker{}, "", err
	}

	worker.CreatedAt = createdAt.Format(Layout)

	if updatedAt.Valid {
		worker.UpdatedAt = updatedAt.Time.Format(Layout)
	}

	return &worker, branchId, nil
}

func (w *workersOfBranchesRepo) GetByWorkerId(ctx context.Context, request *pb.WorkerId) (*pb.Worker, string, error) {

	var (
		worker    pb.Worker
		query     string
		err       error
		createdAt time.Time
		updatedAt sql.NullTime
		branchId  string
	)

	query = `
		select 
			user_id,
			worker_id,
			branch_id,
			created_at,
			updated_at
		from 
			workers_of_branches
		where
			worker_id = $1 and
			deleted_at is null `

	if err = w.db.QueryRow(ctx, query, request.GetWorkerId()).Scan(
		&worker.Id,
		&worker.WorkerId,
		&branchId,
		&createdAt,
		&updatedAt,
	); err != nil {
		w.log.Error("error while getting workers from workers_of_branches table by worker id in storage layer", logger.Error(err))
		return &pb.Worker{}, "", err
	}

	worker.CreatedAt = createdAt.Format(Layout)

	if updatedAt.Valid {
		worker.UpdatedAt = updatedAt.Time.Format(Layout)
	}

	return &worker, branchId, nil
}

func (w *workersOfBranchesRepo) GetAll(ctx context.Context, request *pb.WorkerFilter) (*pb.Workers, []string, error) {

	var (
		workers   pb.Workers
		branchIds []string
		query     string
		err       error
		filter    = ""
		params    = make(map[string]interface{})
		createdAt time.Time
		updatedAt sql.NullTime
	)

	query = `
	select 
		user_id,
		worker_id,
		branch_id,
		created_at,
		updated_at
	from 
		workers_of_branches
	where `

	if request.GetBranchId() != "" {
		filter += ` branch_id = @branch_id and `
		params["branch_id"] = request.GetBranchId()
	}

	query += filter + ` deleted_at is null `

	fullQuery, args := helper.ReplaceQueryParams(query, params)

	rows, err := w.db.Query(ctx, fullQuery, args...)
	if err != nil {
		w.log.Error("error while taking rows in storage layer", logger.Error(err))
		return &pb.Workers{}, []string{}, err
	}

	for rows.Next() {
		var worker pb.Worker
		var branchId string

		if err = rows.Scan(
			&worker.Id,
			&worker.WorkerId,
			&branchId,
			&createdAt,
			&updatedAt,
		); err != nil {
			w.log.Error("error while getting worker from workers_of_branches table in storage layer", logger.Error(err))
			return &pb.Workers{}, []string{}, err
		}

		worker.CreatedAt = createdAt.Format(Layout)

		if updatedAt.Valid {
			worker.UpdatedAt = updatedAt.Time.Format(Layout)
		}

		branchIds = append(branchIds, branchId)
		workers.Workers = append(workers.Workers, &worker)
	}

	if err = rows.Err(); err != nil {
		w.log.Error("error while iterating rows in storage layer", logger.Error(err))
		return &pb.Workers{}, []string{}, err
	}

	return &workers, branchIds, nil
}

func (w *workersOfBranchesRepo) Update(ctx context.Context, request *pb.UpdateWorker) (*pb.Worker, string, error) {

	var (
		worker      pb.Worker
		queryUser   = ` update users set `
		queryWorker = ` update workers_of_branches set `
		branchId    string
		err         error
		birthday    time.Time
		filter      = ""
		params      = make(map[string]interface{})
		createdAt   time.Time
		updatedAt   sql.NullTime
	)

	tx, err := w.db.Begin(ctx)
	if err != nil {
		return &pb.Worker{}, "", err
	}
	defer tx.Commit(ctx)

	params["id"] = request.GetId()

	if request.GetPhoneNumber() != "" {
		filter += ` phone_number = @phone_number, `
		params["phone_number"] = request.GetPhoneNumber()
	}

	if request.GetFullName() != "" {
		filter += ` full_name = @full_name, `
		params["full_name"] = request.GetFullName()
	}

	if request.GetBirthday() != "" {
		filter += ` birthday = @birthday, `
		params["birthday"] = request.GetBirthday()
	}

	if request.GetUserRole() != "" {
		filter += ` user_role = @user_role, `
		params["user_role"] = request.GetUserRole()
	}

	queryUser += filter + ` updated_at = now() where id = @id and deleted_at is null returning 
		id,
		phone_number,
		full_name,
		birthday,
		user_role
	`

	fullQuery, args := helper.ReplaceQueryParams(queryUser, params)

	if err = tx.QueryRow(ctx, fullQuery, args...).Scan(
		&worker.Id,
		&worker.PhoneNumber,
		&worker.FullName,
		&birthday,
		&worker.UserRole,
	); err != nil {
		tx.Rollback(ctx)
		w.log.Error("error while updating worker info in users table", logger.Error(err))
		return &pb.Worker{}, "", err
	}
	worker.Birthday = birthday.Format("02.01.2006")

	filter = ""
	params = make(map[string]interface{})

	if request.GetWorkerId() == "" {
		tx.Rollback(ctx)
		w.log.Error("error: worker id is requered")
		return &pb.Worker{}, "", fmt.Errorf("error: worker id is requered")
	}

	params["worker_id"] = request.GetWorkerId()

	if request.GetBranchId() != "" {
		filter += ` branch_id = @branch_id, `
		params["branch_id"] = request.GetBranchId()
	}

	queryWorker += filter + ` updated_at = now() where worker_id = @worker_id and deleted_at is null returning
		worker_id,
		branch_id,
		created_at,
		updated_at `

	fullQuery, args = helper.ReplaceQueryParams(queryWorker, params)
	if err = tx.QueryRow(ctx, fullQuery, args...).Scan(
		&worker.WorkerId,
		&branchId,
		&createdAt,
		&updatedAt,
	); err != nil {
		tx.Rollback(ctx)
		w.log.Error("error while updating worker info in workers_of_branches table", logger.Error(err))
		return &pb.Worker{}, "", err
	}

	worker.CreatedAt = createdAt.Format(Layout)

	if updatedAt.Valid {
		worker.UpdatedAt = updatedAt.Time.Format(Layout)
	}

	return &worker, branchId, nil
}

func (w *workersOfBranchesRepo) Delete(ctx context.Context, request *pb.WorkerId) (*pb.Void, error) {

	var (
		queryUser   string
		queryWorker string
		err         error
		userId      string
	)

	tx, err := w.db.Begin(ctx)
	if err != nil {
		tx.Rollback(ctx)
		w.log.Error("error while creating transaction", logger.Error(err))
		return &pb.Void{}, err
	}
	defer tx.Commit(ctx)

	queryWorker = `
		update
			workers_of_branches
		set
			deleted_at = now()
		where
			worker_id = $1 and
			deleted_at is null 
		returning
			user_id `

	if err = tx.QueryRow(ctx, queryWorker, request.GetWorkerId()).Scan(&userId); err != nil {
		tx.Rollback(ctx)
		w.log.Error("error while deleting worker info from workers_of_branches table", logger.Error(err))
		return &pb.Void{}, err
	}

	queryUser = `
		update
			users
		set
			deleted_at = now()
		where
			id = $1 and
			deleted_at is null `

	if _, err = tx.Exec(ctx, queryUser, userId); err != nil {
		tx.Rollback(ctx)
		w.log.Error("error while deleting worker info from users table", logger.Error(err))
		return &pb.Void{}, err
	}

	return &pb.Void{}, nil
}

func (w *workersOfBranchesRepo) CheckWorkerExists(ctx context.Context, request *pb.WorkerId) (*pb.Void, error) {

	var (
		exist int
		query string
		err   error
	)

	query = `
		select
			1
		from
			workers_of_branches
		where 
			worker_id = $1 and 
			deleted_at is null `

	err = w.db.QueryRow(ctx, query, request.GetWorkerId()).Scan(&exist)

	if err != nil && err.Error() == "no rows in result set" {
		return &pb.Void{}, fmt.Errorf("error: worker not found")
	}

	return &pb.Void{}, nil
}
