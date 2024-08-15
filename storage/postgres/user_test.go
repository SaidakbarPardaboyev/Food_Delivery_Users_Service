package postgres

import (
	"context"
	"testing"
	"users_service/configs"
	pb "users_service/genproto/users"
	"users_service/pkg/logger"

	"github.com/google/uuid"
)

func setupUserRepo() (*usersRepo, func(), error) {
	cfg := configs.Load()
	db, err := ConnectDB(context.Background(), cfg)
	if err != nil {
		return nil, nil, err
	}

	logger := logger.NewLogger(cfg.ServiceName, "debug", cfg.LogPath)
	repo := &usersRepo{
		db:  db,
		log: logger,
	}

	cleanup := func() {
		db.Exec(context.Background(), "TRUNCATE TABLE users CASCADE")
	}

	return repo, cleanup, nil
}

func TestGetById(t *testing.T) {
	repo, cleanup, err := setupUserRepo()
	if err != nil {
		t.Fatal(err)
	}
	defer cleanup()

	userID := uuid.New().String()
	_, err = repo.db.Exec(context.Background(), "INSERT INTO users(id, phone_number, full_name, user_role) VALUES($1, '1234567890', 'John Doe', 'user')", userID)
	if err != nil {
		t.Fatal(err)
	}

	req := &pb.PrimaryKey{Id: userID}
	user, err := repo.GetById(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}

	if user.Id != userID {
		t.Errorf("expected user ID %s, got %s", userID, user.Id)
	}
}

func TestGetAll(t *testing.T) {
	repo, cleanup, err := setupUserRepo()
	if err != nil {
		t.Fatal(err)
	}
	defer cleanup()

	userID1 := uuid.New().String()
	userID2 := uuid.New().String()
	_, err = repo.db.Exec(context.Background(), "INSERT INTO users(id, phone_number, full_name, user_role) VALUES($1, '1234567891', 'Jane Doe', 'admin')", userID1)
	if err != nil {
		t.Fatal(err)
	}
	_, err = repo.db.Exec(context.Background(), "INSERT INTO users(id, phone_number, full_name, user_role) VALUES($1, '1234567892', 'John Smith', 'user')", userID2)
	if err != nil {
		t.Fatal(err)
	}

	req := &pb.GetListRequest{
		Page:  1,
		Limit: 10,
	}
	users, err := repo.GetAll(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}

	if len(users.Users) != 2 {
		t.Errorf("expected 2 users, got %d", len(users.Users))
	}
}

func TestUpdate(t *testing.T) {
	repo, cleanup, err := setupUserRepo()
	if err != nil {
		t.Fatal(err)
	}
	defer cleanup()

	userID := uuid.New().String()
	_, err = repo.db.Exec(context.Background(), "INSERT INTO users(id, phone_number, full_name, user_role) VALUES($1, '1234567890', 'John Doe', 'user')", userID)
	if err != nil {
		t.Fatal(err)
	}

	req := &pb.UpdateUser{
		Id:       userID,
		FullName: "John Updated",
	}
	user, err := repo.Update(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}

	if user.FullName != "John Updated" || user.UserRole != "user" {
		t.Errorf("expected updated user details, got %v", user)
	}
}

func TestDelete(t *testing.T) {
	repo, cleanup, err := setupUserRepo()
	if err != nil {
		t.Fatal(err)
	}
	defer cleanup()

	userID := uuid.New().String()
	_, err = repo.db.Exec(context.Background(), "INSERT INTO users(id, phone_number, full_name, user_role) VALUES($1, '1234567890', 'John Doe', 'user')", userID)
	if err != nil {
		t.Fatal(err)
	}

	req := &pb.PrimaryKey{Id: userID}
	_, err = repo.Delete(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}

	var count int
	err = repo.db.QueryRow(context.Background(), "SELECT count(*) FROM users WHERE id = $1 AND deleted_at IS NOT NULL", userID).Scan(&count)
	if err != nil {
		t.Fatal(err)
	}
	if count != 1 {
		t.Errorf("expected 1 deleted user, got %d", count)
	}
}

func TestChangeUserRole(t *testing.T) {
	repo, cleanup, err := setupUserRepo()
	if err != nil {
		t.Fatal(err)
	}
	defer cleanup()

	userID := uuid.New().String()
	_, err = repo.db.Exec(context.Background(), "INSERT INTO users(id, phone_number, full_name, user_role) VALUES($1, '1234567890', 'John Doe', 'user')", userID)
	if err != nil {
		t.Fatal(err)
	}

	req := &pb.ChangeUserRole{
		Id:          userID,
		NewUserRole: "admin",
	}
	_, err = repo.ChangeUserRole(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}

	var userRole string
	err = repo.db.QueryRow(context.Background(), "SELECT user_role FROM users WHERE id = $1", userID).Scan(&userRole)
	if err != nil {
		t.Fatal(err)
	}
	if userRole != "admin" {
		t.Errorf("expected user role 'admin', got %s", userRole)
	}
}

func TestCheckUserIdExists(t *testing.T) {
	repo, cleanup, err := setupUserRepo()
	if err != nil {
		t.Fatal(err)
	}
	defer cleanup()

	userID := uuid.New().String()
	_, err = repo.db.Exec(context.Background(), "INSERT INTO users(id, phone_number, full_name, user_role) VALUES($1, '1234567890', 'John Doe', 'user')", userID)
	if err != nil {
		t.Fatal(err)
	}

	req := &pb.PrimaryKey{Id: userID}
	_, err = repo.CheckUserIdExists(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}

	nonExistentID := uuid.New().String()
	req = &pb.PrimaryKey{Id: nonExistentID}
	_, err = repo.CheckUserIdExists(context.Background(), req)
	if err == nil {
		t.Error("expected error for non-existent user ID")
	}
}
