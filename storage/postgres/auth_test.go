package postgres

import (
	"context"
	"testing"
	"time"
	"users_service/configs"
	pb "users_service/genproto/users"
	"users_service/pkg/logger"
)

func setupRepo() (*authRepo, error) {
	cfg := configs.Load()
	db, err := ConnectDB(context.Background(), cfg)
	if err != nil {
		return nil, err
	}

	logger := logger.NewLogger(cfg.ServiceName, "debug", cfg.LogPath)

	return &authRepo{
		db:  db,
		log: logger,
	}, nil
}

func TestCreateUser(t *testing.T) {
	repo, err := setupRepo()
	if err != nil {
		t.Error(err)
		return
	}
	defer repo.db.Exec(context.Background(), "DELETE FROM users WHERE phone_number = $1", "1234567890")

	phoneNumber := "1234567890"
	req := pb.CreateUser{
		PhoneNumber: phoneNumber,
		FullName:    "John Doe",
	}

	user, err := repo.Create(context.Background(), &req)
	if err != nil {
		t.Error(err)
		return
	}

	if user.PhoneNumber != req.PhoneNumber {
		t.Error("user phone number does not match")
	}
}

func TestGetByPhone(t *testing.T) {
	repo, err := setupRepo()
	if err != nil {
		t.Error(err)
		return
	}
	repo.db.Exec(context.Background(), "DELETE FROM users WHERE phone_number = $1", "12345634890")
	defer repo.db.Exec(context.Background(), "DELETE FROM users WHERE phone_number = $1", "12345634890")

	phoneNumber := "12345634890"
	createReq := pb.CreateUser{
		PhoneNumber: phoneNumber,
		FullName:    "John Doe",
	}
	user, err := repo.Create(context.Background(), &createReq)
	if err != nil {
		t.Error(err)
		return
	}

	getReq := pb.Phone{PhoneNumber: phoneNumber}
	retrievedUser, err := repo.GetByPhone(context.Background(), &getReq)
	if err != nil {
		t.Error(err)
		return
	}

	if retrievedUser.Id != user.Id {
		t.Error("retrieved user ID does not match")
	}
}

func TestCheckRefreshTokenExists(t *testing.T) {
	repo, err := setupRepo()
	if err != nil {
		t.Error(err)
	}
	defer repo.db.Exec(context.Background(), "DELETE FROM refresh_tokens WHERE refresh_token = $1", "test-refresh-token")
	defer repo.db.Exec(context.Background(), "DELETE FROM users WHERE phone_number = $1", "12345634890")

	phoneNumber := "12345634890"
	createReq := pb.CreateUser{
		PhoneNumber: phoneNumber,
		FullName:    "John Doe",
	}
	user, err := repo.Create(context.Background(), &createReq)
	if err != nil {
		t.Error(err)
	}

	storeReq := pb.RefreshToken{
		UserId:       user.Id,
		RefreshToken: "test-refresh-token",
		ExpiresAt:    time.Now().Add(time.Hour).Format(time.RFC3339),
	}

	_, err = repo.StoreRefreshToken(context.Background(), &storeReq)
	if err != nil {
		t.Error(err)
	}

	checkReq := pb.RequestRefreshToken{RefreshToken: "test-refresh-token"}
	_, err = repo.CheckRefreshTokenExists(context.Background(), &checkReq)
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteRefreshTokenByUserId(t *testing.T) {
	repo, err := setupRepo()
	if err != nil {
		t.Error(err)
		return
	}
	repo.db.Exec(context.Background(), "DELETE FROM users WHERE phone_number = $1", "12345634890")
	defer repo.db.Exec(context.Background(), "DELETE FROM refresh_tokens WHERE refresh_token = $1", "test-refresh-token")
	defer repo.db.Exec(context.Background(), "DELETE FROM users WHERE phone_number = $1", "12345634890")

	phoneNumber := "12345634890"
	createReq := pb.CreateUser{
		PhoneNumber: phoneNumber,
		FullName:    "John Doe",
	}
	user, err := repo.Create(context.Background(), &createReq)
	if err != nil {
		t.Error(err)
		return
	}

	storeReq := pb.RefreshToken{
		UserId:       user.Id,
		RefreshToken: "test-refresh-token",
		ExpiresAt:    time.Now().Add(time.Hour).Format(time.RFC3339),
	}
	_, err = repo.StoreRefreshToken(context.Background(), &storeReq)
	if err != nil {
		t.Error(err)
		return
	}

	deleteReq := pb.PrimaryKey{Id: user.Id}
	_, err = repo.DeleteRefreshTokenByUserId(context.Background(), &deleteReq)
	if err != nil {
		t.Error(err)
		return
	}

	checkReq := pb.RequestRefreshToken{RefreshToken: "test-refresh-token"}
	_, err = repo.CheckRefreshTokenExists(context.Background(), &checkReq)
	if err == nil {
		t.Error("refresh token should not exist")
	}
}

func TestStoreRefreshToken(t *testing.T) {
	repo, err := setupRepo()
	if err != nil {
		t.Error(err)
		return
	}
	repo.db.Exec(context.Background(), "DELETE FROM users WHERE phone_number = $1", "12345634890")
	defer repo.db.Exec(context.Background(), "DELETE FROM refresh_tokens WHERE refresh_token = $1", "test-refresh-token")
	defer repo.db.Exec(context.Background(), "DELETE FROM users WHERE phone_number = $1", "12345634890")

	phoneNumber := "12345634890"
	createReq := pb.CreateUser{
		PhoneNumber: phoneNumber,
		FullName:    "John Doe",
	}
	user, err := repo.Create(context.Background(), &createReq)
	if err != nil {
		t.Error(err)
		return
	}

	req := pb.RefreshToken{
		UserId:       user.Id,
		RefreshToken: "test-refresh-token",
		ExpiresAt:    time.Now().Add(time.Hour).Format(time.RFC3339),
	}

	_, err = repo.StoreRefreshToken(context.Background(), &req)
	if err != nil {
		t.Error(err)
	}
}
