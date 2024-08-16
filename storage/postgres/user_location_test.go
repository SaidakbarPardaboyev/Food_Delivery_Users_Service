package postgres

import (
	"context"
	"testing"
	"users_service/configs"
	"users_service/pkg/logger"

	pb "users_service/genproto/users"

	"github.com/stretchr/testify/require"
)

func setupUserLocationRepo() (*userLocationRepo, func(), error) {
	cfg := configs.Load()
	db, err := ConnectDB(context.Background(), cfg)
	if err != nil {
		return nil, nil, err
	}

	logger := logger.NewLogger(cfg.ServiceName, "debug", cfg.LogPath)
	repo := NewUserLoactionRepo(db, logger)

	cleanup := func() {
		db.Exec(context.Background(), "TRUNCATE TABLE user_locations CASCADE")
	}

	return repo, cleanup, nil
}

func TestCreateUserLocation(t *testing.T) {
	repo, cleanup, err := setupUserLocationRepo()
	require.NoError(t, err)
	defer cleanup()

	req := &pb.CreateUserLocation{
		UserId:          "024fcac1-b077-4d5c-b22a-671e01904025",
		Address:         "123 Test St",
		HomeNumber:      "1A",
		FloorNumber:     2,
		ApartmentNumber: 3,
		EntranceNumber:  4,
		Latitude:        40.7128,
		Longitude:       -74.0060,
	}

	location, err := repo.Create(context.Background(), req)
	if err != nil {
		t.Fatalf("Error creating user location: %v", err)
	}

	if location == nil {
		t.Fatal("Created location is nil")
	}

	if location.UserId != req.UserId {
		t.Errorf("Expected UserId %v, got %v", req.UserId, location.UserId)
	}
}

func TestGetUserLocationById(t *testing.T) {
	repo, cleanup, err := setupUserLocationRepo()
	require.NoError(t, err)
	defer cleanup()

	createReq := &pb.CreateUserLocation{ UserId: "024fcac1-b077-4d5c-b22a-671e01904025" }
	location, err := repo.Create(context.Background(), createReq)
	if err != nil {
		t.Error(err)
	}

	getReq := &pb.PrimaryKey{Id: location.Id}
	result, err := repo.GetById(context.Background(), getReq)
	require.NoError(t, err)

	if result == nil {
		t.Fatal("Result is nil")
	}
	if result.Id != location.Id {
		t.Errorf("Expected Id %v, got %v", location.Id, result.Id)
	}
}

func TestGetAllUserLocations(t *testing.T) {
	repo, cleanup, err := setupUserLocationRepo()
	require.NoError(t, err)
	defer cleanup()

	_, err = repo.Create(context.Background(), &pb.CreateUserLocation{
		UserId:          "024fcac1-b077-4d5c-b22a-671e01904025",
		Address:         "Address ",
		HomeNumber:      "Home ",
		FloorNumber:     1,
		ApartmentNumber: 1,
		EntranceNumber:  1,
		Latitude:        40.7128,
		Longitude:       -74.0060,
	})
	if err != nil {
		t.Fatalf("Error creating user location: %v", err)
	}

	filter := &pb.UserLocationFilter{}
	result, err := repo.GetAll(context.Background(), filter)
	if err != nil {
		t.Fatalf("Error getting all user locations: %v", err)
	}

	if result == nil {
		t.Fatal("Result is nil")
	}
	if len(result.UserLocations) < 1 {
		t.Errorf("Expected 1+ user locations, got %v", len(result.UserLocations))
	}
}

func TestUpdateUserLocation(t *testing.T) {
	repo, cleanup, err := setupUserLocationRepo()
	require.NoError(t, err)
	defer cleanup()

	createReq := &pb.CreateUserLocation{ UserId: "024fcac1-b077-4d5c-b22a-671e01904025"}
	loc, err := repo.Create(context.Background(), createReq)
	require.NoError(t, err)

	updateReq := &pb.UpdateUserLocation{ Id: loc.Id }
	updatedLocation, err := repo.Update(context.Background(), updateReq)
	require.NoError(t, err)

	if updatedLocation == nil {
		t.Fatal("Updated location is nil")
	}
}

func TestDeleteUserLocation(t *testing.T) {
	repo, cleanup, err := setupUserLocationRepo()
	require.NoError(t, err)
	defer cleanup()

	createReq := &pb.CreateUserLocation{ UserId: "024fcac1-b077-4d5c-b22a-671e01904025" }
	location, err := repo.Create(context.Background(), createReq)
	require.NoError(t, err)

	deleteReq := &pb.PrimaryKey{Id: location.Id}
	_, err = repo.Delete(context.Background(), deleteReq)
	require.NoError(t, err)

}
