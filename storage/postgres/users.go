package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"users_service/pkg/helper"
	"users_service/pkg/logger"

	"github.com/jackc/pgx/v5/pgxpool"

	pb "users_service/genproto/users"
)

type usersRepo struct {
	db  *pgxpool.Pool
	log logger.ILogger
}

func NewUsersRepo(db *pgxpool.Pool, log logger.ILogger) *usersRepo {
	return &usersRepo{
		db:  db,
		log: log,
	}
}

func (u *usersRepo) GetById(ctx context.Context, request *pb.PrimaryKey) (*pb.User, error) {

	var (
		user      = pb.User{}
		query     string
		err       error
		createdAt sql.NullTime
		updatedAt sql.NullTime
	)

	query = `select
		id,
		full_name,
		phone_number,
		user_role,
		created_at,
		updated_at
	from
		users
	where
		id = $1 and
		deleted_at is null
	`

	if err = u.db.QueryRow(ctx, query,
		request.GetId()).
		Scan(
			&user.Id,
			&user.FullName,
			&user.PhoneNumber,
			&user.UserRole,
			&createdAt,
			&updatedAt,
		); err != nil {
		u.log.Error("error while getting user info in storage layer", logger.Error(err))
		return nil, err
	}

	if createdAt.Valid {
		user.CreatedAt = createdAt.Time.Format(Layout)
	}
	if updatedAt.Valid {
		user.UpdatedAt = updatedAt.Time.Format(Layout)
	}

	return &user, nil
}

func (u *usersRepo) GetAll(ctx context.Context, request *pb.GetListRequest) (*pb.Users, error) {

	var (
		users             = []*pb.User{}
		query, countQuery string
		offset            = (request.GetPage() - 1) * int32(request.GetLimit())
		filter            = ""
		params            = make(map[string]interface{})
		err               error
		count             int
		createdAt         sql.NullTime
		updatedAt         sql.NullTime
	)

	if request.GetFullName() != "" {
		filter += " full_name = @full_name and "
		params["full_name"] = request.GetFullName()
	}

	if request.GetPhoneNumber() != "" {
		filter += " phone_number = @phone_number and "
		params["phone_number"] = request.GetPhoneNumber()
	}

	if request.GetUserRole() != "" {
		filter += " user_role = @user_role and "
		params["user_role"] = request.GetUserRole()
	}
	where, args := helper.ReplaceQueryParams(filter, params)

	countQuery = `select count(*) from users where deleted_at is null `
	if len(args) > 0 {
		countQuery += " and " + where[:len(where)-4]
	}

	if err := u.db.QueryRow(ctx, countQuery, args...).Scan(&count); err != nil {
		u.log.Error("error while taking count of users in storage layer", logger.Error(err))
		return &pb.Users{}, err
	}

	query = `select
		id,
		phone_number,
		full_name,
		user_role,
		created_at,
		updated_at
	from
		users
	where 
		deleted_at is null 
	`
	if len(args) > 0 {
		query += " and " + where[:len(where)-4]
	}

	query += fmt.Sprintf(` order by created_at desc LIMIT %d OFFSET %d`, request.Limit, offset)

	rows, err := u.db.Query(ctx, query, args...)
	if err != nil {
		u.log.Error("error while taking rows to get all users in storage layer", logger.Error(err))
		return &pb.Users{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var user pb.User
		if err = rows.Scan(
			&user.Id,
			&user.PhoneNumber,
			&user.FullName,
			&user.UserRole,
			&createdAt,
			&updatedAt,
		); err != nil {
			u.log.Error("error while getting user info in storage layer", logger.Error(err))
			return nil, err
		}
		if createdAt.Valid {
			user.CreatedAt = createdAt.Time.Format(Layout)
		}
		if updatedAt.Valid {
			user.UpdatedAt = updatedAt.Time.Format(Layout)
		}

		users = append(users, &user)
	}
	if rows.Err() != nil {
		u.log.Error("error while iterating rows in storage layer", logger.Error(err))
		return nil, err
	}

	return &pb.Users{
		Users: users,
		Page:  request.Page,
		Limit: request.Limit,
		Count: int32(count),
	}, nil
}

func (u *usersRepo) Update(ctx context.Context, request *pb.UpdateUser) (*pb.User, error) {

	var (
		user      = pb.User{}
		params    = make(map[string]interface{})
		filter    = ""
		query     = ` update users set `
		err       error
		createdAt sql.NullTime
		updatedAt sql.NullTime
	)

	params["id"] = request.GetId()

	if request.GetFullName() != "" {
		filter += ` full_name = @full_name, `
		params["full_name"] = request.GetFullName()
	}

	if request.GetBirthday() != "" {
		filter += ` birthday = @birthday, `
		params["birthday"] = request.GetBirthday()
	}

	query += filter + ` updated_at = now() where id = @id returning 
		id,
		full_name,
		phone_number,
		user_role,
		created_at,
		updated_at
	`
	fullQuery, args := helper.ReplaceQueryParams(query, params)
	if err = u.db.QueryRow(ctx, fullQuery, args...).Scan(
		&user.Id,
		&user.FullName,
		&user.PhoneNumber,
		&user.UserRole,
		&createdAt,
		&updatedAt,
	); err != nil {
		u.log.Error("error while updating user info in storage layer", logger.Error(err))
		return nil, err
	}

	if createdAt.Valid {
		user.CreatedAt = createdAt.Time.Format(Layout)
	}
	if updatedAt.Valid {
		user.UpdatedAt = updatedAt.Time.Format(Layout)
	}

	return &user, nil
}

func (u *usersRepo) Delete(ctx context.Context, request *pb.PrimaryKey) (*pb.Void, error) {

	_, err := u.db.Exec(ctx, ` update users set deleted_at = now() where id = $1`, request.GetId())

	return &pb.Void{}, err
}

func (u *usersRepo) ChangeUserRole(ctx context.Context, request *pb.ChangeUserRole) (*pb.Void, error) {

	var (
		query string
		err   error
	)

	query = `
		update 
			users 
		set 
			user_role = $1
		where
			id = $2 and 
			deleted_at is null
	`

	if _, err = u.db.Exec(ctx, query,
		request.GetNewUserRole(),
		request.GetId(),
	); err != nil {
		u.log.Error("error while changing user role in storage layer", logger.Error(err))
		return nil, err
	}

	return &pb.Void{}, nil
}

func (u *usersRepo) CheckUserIdExists(ctx context.Context, request *pb.PrimaryKey) (*pb.Void, error) {

	var (
		exist int
		query string
		err   error
	)

	query = `
		select
			1
		from
			users
		where 
			id = $1 and 
			deleted_at is null `

	err = u.db.QueryRow(ctx, query, request.GetId()).Scan(&exist)

	if err != nil && err.Error() == "no rows in result set" {
		return &pb.Void{}, fmt.Errorf("error: user not found")
	}

	return &pb.Void{}, err
}
