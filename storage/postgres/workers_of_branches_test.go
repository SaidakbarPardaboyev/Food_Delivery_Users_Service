package postgres

import (
	"context"
	"testing"
	"users_service/configs"
	pb "users_service/genproto/users"
	"users_service/pkg/logger"
)

func NewWorkersOfBranchesRepoTest() (*workersOfBranchesRepo, error) {
	db, err := ConnectDB(context.Background(), configs.Load())
	if err != nil {
		return nil, err
	}

	log := logger.NewLogger("", "debug", "./../../app.log")

	return NewWorkersOfBranchesRepo(db, log), nil
}

func TestCreateWorker(t *testing.T) {
	repo, err := NewWorkersOfBranchesRepoTest()
	if err != nil {
		t.Error(err)
	}

	request := pb.CreateWorker{
		PhoneNumber: "+998997693735",
		FullName:    "Saida",
		Birthday:    "11-06-2005",
		UserRole:    "admin",
		BranchId:    "64a5471a-8d82-4589-a2cc-20c8dfe5e7ab",
	}

	_, err = repo.Create(context.Background(), &request)
	if err != nil {
		t.Error(err)
	}
}

func TestGetWorkerByUUID(t *testing.T) {
	repo, err := NewWorkersOfBranchesRepoTest()
	if err != nil {
		t.Error(err)
	}

	request := pb.PrimaryKey{Id: "2be6efb8-83ac-4984-a3f9-d11c06b681c8"}

	_, _, err = repo.GetByUUID(context.Background(), &request)
	if err != nil {
		t.Error(err)
	}
}

func TestGetWorkerByWorkerId(t *testing.T) {
	repo, err := NewWorkersOfBranchesRepoTest()
	if err != nil {
		t.Error(err)
	}

	request := pb.WorkerId{WorkerId: "1000"}

	_, _, err = repo.GetByWorkerId(context.Background(), &request)
	if err != nil {
		t.Error(err)
	}
}

func TestGetAllWorker(t *testing.T) {
	repo, err := NewWorkersOfBranchesRepoTest()
	if err != nil {
		t.Error(err)
	}

	request := pb.WorkerFilter{
		BranchId: "632f717f-0d32-48e6-9c71-0abd463521f9",
		UserRole: "",
		Page:     1,
		Limit:    10,
	}

	_, _, err = repo.GetAll(context.Background(), &request)
	if err != nil {
		t.Error(err)
	}
}

func TestUpdateWorker(t *testing.T) {
	repo, err := NewWorkersOfBranchesRepoTest()
	if err != nil {
		t.Error(err)
	}

	request := pb.UpdateWorker{
		Id:          "4f4d8d9c-8f30-425e-917a-c7d8d7449959",
		PhoneNumber: "+998997693733",
		FullName:    "Akbar",
		Birthday:    "18-08-2006",
		UserRole:    "chief",
		BranchId:    "64a5471a-8d82-4589-a2cc-20c8dfe5e7ab",
		WorkerId:    "1001",
	}

	_, _, err = repo.Update(context.Background(), &request)
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteWorker(t *testing.T) {
	repo, err := NewWorkersOfBranchesRepoTest()
	if err != nil {
		t.Error(err)
	}

	request := pb.WorkerId{
		WorkerId: "1002",
	}

	_, err = repo.Delete(context.Background(), &request)
	if err != nil {
		t.Error(err)
	}
}

func TestCheckWorkerExists(t *testing.T) {
	repo, err := NewWorkersOfBranchesRepoTest()
	if err != nil {
		t.Error(err)
	}

	request := pb.WorkerId{
		WorkerId: "1000",
	}

	_, err = repo.CheckWorkerExists(context.Background(), &request)
	if err != nil {
		t.Error(err)
	}
}
