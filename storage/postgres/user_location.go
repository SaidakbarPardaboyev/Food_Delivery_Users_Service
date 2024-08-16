package postgres

import (
	"context"
	"database/sql"
	"time"
	"users_service/pkg/helper"
	"users_service/pkg/logger"

	pb "users_service/genproto/users"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type userLocationRepo struct {
	db  *pgxpool.Pool
	log logger.ILogger
}

func NewUserLoactionRepo(db *pgxpool.Pool, log logger.ILogger) *userLocationRepo {
	return &userLocationRepo{
		db:  db,
		log: log,
	}
}

func (u *userLocationRepo) Create(ctx context.Context, req *pb.CreateUserLocation) (*pb.UserLocation, error) {
	var (
		query     string
		err       error
		res       = pb.UserLocation{}
		createdAt time.Time
	)

	query = `
	insert into user_locations(
		user_id,
		address,
		home_number,
		floor_number,
		apartment_number,
		entrance_number,
		latitute,
		longitute
	) values($1, $2, $3, $4, $5, $6, $7, $8) returning
	 	id,
		user_id,
		address,
		home_number,
		floor_number,
		apartment_number,
		entrance_number,
		latitute,
		longitute,
		created_at
	`

	err = u.db.QueryRow(ctx, query,
		req.UserId,
		req.Address,
		req.HomeNumber,
		req.FloorNumber,
		req.ApartmentNumber,
		req.EntranceNumber,
		req.Latitude,
		req.Longitude,
	).Scan(
		&res.Id,
		&res.UserId,
		&res.Address,
		&res.HomeNumber,
		&res.FloorNumber,
		&res.ApartmentNumber,
		&res.EntranceNumber,
		&res.Latitude,
		&res.Longitude,
		&createdAt,
	)
	if err != nil {
		u.log.Error("Error while creating user location in storage layer")
		return nil, err
	}

	res.CreatedAt = createdAt.Format(Layout)

	return &res, nil
}

func (u *userLocationRepo) GetById(ctx context.Context, req *pb.PrimaryKey) (*pb.UserLocation, error) {
	var (
		res       = pb.UserLocation{}
		err       error
		query     string
		createdAt sql.NullTime
		updatedAt sql.NullTime
	)

	query = `
	select
		id,
		user_id,
		address,
		home_number,
		floor_number,
		apartment_number,
		entrance_number,
		latitute,
		longitute,
		created_at,
		updated_at
	from
		user_locations
	where
		id = $1 and deleted_at is null
	`

	err = u.db.QueryRow(ctx, query,
		req.Id,
	).Scan(
		&res.Id,
		&res.UserId,
		&res.Address,
		&res.HomeNumber,
		&res.FloorNumber,
		&res.ApartmentNumber,
		&res.EntranceNumber,
		&res.Latitude,
		&res.Longitude,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		u.log.Error("Error while getting user location by id in storage layer")
		return nil, err
	}
	if createdAt.Valid {
		res.CreatedAt = createdAt.Time.Format(Layout)
	}
	if updatedAt.Valid {
		res.UpdatedAt = updatedAt.Time.Format(Layout)
	}

	return &res, nil
}

func (u *userLocationRepo) GetByUserId(ctx context.Context, req *pb.PrimaryKey) (*pb.UserLocation, error) {
	var (
		res       = pb.UserLocation{}
		err       error
		query     string
		createdAt sql.NullTime
		updatedAt sql.NullTime
	)

	query = `
	select
		id,
		user_id,
		address,
		home_number,
		floor_number,
		apartment_number,
		entrance_number,
		latitute,
		longitute,
		created_at,
		updated_at
	from
		user_locations
	where
		user_id = $1 and deleted_at is null
	`

	err = u.db.QueryRow(ctx, query,
		req.Id,
	).Scan(
		&res.Id,
		&res.UserId,
		&res.Address,
		&res.HomeNumber,
		&res.FloorNumber,
		&res.ApartmentNumber,
		&res.EntranceNumber,
		&res.Latitude,
		&res.Longitude,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		u.log.Error("Error while getting user location by id in storage layer")
		return nil, err
	}
	if createdAt.Valid {
		res.CreatedAt = createdAt.Time.Format(Layout)
	}
	if updatedAt.Valid {
		res.UpdatedAt = updatedAt.Time.Format(Layout)
	}

	return &res, nil
}

func (u *userLocationRepo) GetAll(ctx context.Context, req *pb.UserLocationFilter) (*pb.UserLocations, error) {
	var (
		res          = pb.UserLocations{}
		userLocation = pb.UserLocation{}
		err          error
		query        string
		filter       string
		params       = map[string]interface{}{}
		rows         pgx.Rows
		createdAt    sql.NullTime
		updatedAt    sql.NullTime
	)

	query = `
	select
		id,
		user_id,
		address,
		home_number,
		floor_number,
		apartment_number,
		entrance_number,
		latitute,
		longitute,
		created_at,
		updated_at
	from
		user_locations
	where
		deleted_at is null	 
	`

	if req.Address != "" {
		filter += " and address = @address "
		params["address"] = req.Address
	}
	if req.HomeNumber != "" {
		filter += " and home_number = @home_number "
		params["home_number"] = req.HomeNumber
	}
	if req.FloorNumber != 0 {
		filter += " and floor_number = @floor_number "
		params["floor_number"] = req.FloorNumber
	}
	if req.ApartmentNumber != 0 {
		filter += " and apartment_number = @apartment_number "
		params["apartment_number"] = req.ApartmentNumber
	}
	if req.EntranceNumber != 0 {
		filter += " and entrance_number = @entrance_number "
		params["entrance_number"] = req.EntranceNumber
	}
	f, args := helper.ReplaceQueryParams(filter, params)

	query += f
	rows, err = u.db.Query(ctx, query, args...)
	if err != nil {
		u.log.Error("Error while getting user location by id in storage layer")
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&userLocation.Id,
			&userLocation.UserId,
			&userLocation.Address,
			&userLocation.HomeNumber,
			&userLocation.FloorNumber,
			&userLocation.ApartmentNumber,
			&userLocation.EntranceNumber,
			&userLocation.Latitude,
			&userLocation.Longitude,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			u.log.Error("Error while getting user location by id in storage layer")
			return nil, err
		}
		if createdAt.Valid {
			userLocation.CreatedAt = createdAt.Time.Format(Layout)
		}
		if updatedAt.Valid {
			userLocation.UpdatedAt = updatedAt.Time.Format(Layout)
		}

		res.UserLocations = append(res.UserLocations, &userLocation)
	}
	if err = rows.Err(); err != nil {
		u.log.Error("Error while getting user location by id in storage layer")
		return nil, err
	}

	return &res, nil
}

func (u *userLocationRepo) Update(ctx context.Context, req *pb.UpdateUserLocation) (*pb.UserLocation, error) {
	var (
		query     string
		err       error
		res       = pb.UserLocation{}
		params    = map[string]interface{}{}
		createdAt sql.NullTime
		updatedAt sql.NullTime
	)
	query = `
	update
		user_locations
	set
		updated_at = now()
	`
	params["id"] = req.Id

	if req.Address != "" {
		query += " , address = @address"
		params["address"] = req.Address
	}
	if req.HomeNumber != "" {
		query += " , home_number = @home_number"
		params["home_number"] = req.HomeNumber
	}
	if req.FloorNumber != 0 {
		query += " , floor_number = @floor_number "
		params["floor_number"] = req.FloorNumber
	}
	if req.ApartmentNumber != 0 {
		query += " , apartment_number = @apartment_number "
		params["apartment_number"] = req.ApartmentNumber
	}
	if req.EntranceNumber != 0 {
		query += " , entrance_number = @entrance_number "
		params["entrance_number"] = req.EntranceNumber
	}
	if req.EntranceNumber != 0 {
		query += " , entrance_number = @entrance_number "
		params["entrance_number"] = req.EntranceNumber
	}
	if req.Latitude != 0 {
		query += " , latitute = @latitute "
		params["latitute"] = req.Latitude
	}
	if req.Longitude != 0 {
		query += " , longitute = @longitute "
		params["longitute"] = req.Longitude
	}

	query += `
	where id = @id and deleted_at is null
	returning
		id,
		user_id,
		address,
		home_number,
		floor_number,
		apartment_number,
		entrance_number,
		latitute,
		longitute,
		created_at,
		updated_at
	`

	filter, args := helper.ReplaceQueryParams(query, params)

	err = u.db.QueryRow(ctx, filter, args...).Scan(
		&res.Id,
		&res.UserId,
		&res.Address,
		&res.HomeNumber,
		&res.FloorNumber,
		&res.ApartmentNumber,
		&res.EntranceNumber,
		&res.Latitude,
		&res.Longitude,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		u.log.Error("Error while updating user location by id in storage layer")
		return nil, err
	}

	if createdAt.Valid {
		res.CreatedAt = createdAt.Time.Format(Layout)
	}
	if updatedAt.Valid {
		res.UpdatedAt = updatedAt.Time.Format(Layout)
	}

	return &res, nil
}

func (u *userLocationRepo) Delete(ctx context.Context, req *pb.PrimaryKey) (*pb.Void, error) {
	_, err := u.db.Exec(ctx, ` update user_locations set deleted_at = now() where id = $1`, req.GetId())

	return &pb.Void{}, err
}
