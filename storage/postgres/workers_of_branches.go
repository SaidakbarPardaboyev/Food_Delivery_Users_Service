package postgres

import (
	"context"
	"database/sql"
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
		offset    = int(request.GetPage()-1) * int(request.GetLimit())
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

	query += filter + ` deleted_at is null limit @limit offset @offset `
	params["limit"] = request.GetLimit()
	params["offset"] = offset

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

// func (w *workersOfBranchesRepo) Update(ctx context.Context, request *pb.Worker) (*pb.Worker, error) {

// }

// func (w *workersOfBranchesRepo) Delete(ctx context.Context, request *pb.WorkerId) (*pb.Void, error) {

// }

// func (w *workersOfBranchesRepo) CheckWorkerExists(ctx context.Context, request *pb.WorkerId) (*pb.Void, error) {

// }
