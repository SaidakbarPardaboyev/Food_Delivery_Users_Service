package service

import (
	"context"
	pbo "users_service/genproto/orders_service"
	pb "users_service/genproto/users"
	"users_service/grpc/clients"
	"users_service/pkg/logger"
	"users_service/storage"
)

type workersOfBranchesService struct {
	storage storage.IStorage
	clients clients.IClientsMeneger
	log     logger.ILogger
	pb.UnimplementedWorkersOfBranchesServiceServer
}

func NewWorkersOfBranchesService(storage storage.IStorage, clients clients.IClientsMeneger, log logger.ILogger) *workersOfBranchesService {
	return &workersOfBranchesService{
		storage: storage,
		log:     log,
	}
}

func (w *workersOfBranchesService) Create(ctx context.Context, request *pb.CreateWorker) (*pb.WorkerId, error) {

	if _, err := w.clients.BranchesService().CheckBranchExist(ctx, &pbo.PrimaryKey{Id: request.BranchId}); err != nil {
		w.log.Error("error while checking branch is exists in service layer", logger.Error(err))
		return &pb.WorkerId{}, err
	}

	resp, err := w.storage.WorkersOfBranches().Create(ctx, request)
	if err != nil {
		w.log.Error("error while creating new worker in service layer", logger.Error(err))
		return &pb.WorkerId{}, err
	}

	return resp, nil
}

func (w *workersOfBranchesService) GetByUUID(ctx context.Context, request *pb.PrimaryKey) (*pb.Worker, error) {

	worker, branchId, err := w.storage.WorkersOfBranches().GetByUUID(ctx, request)
	if err != nil {
		w.log.Error("error while getting worker info from workers_of_branches table by uuid in service layer", logger.Error(err))
		return &pb.Worker{}, err
	}

	user, err := w.storage.Users().GetById(ctx, &pb.PrimaryKey{Id: worker.Id})
	if err != nil {
		w.log.Error("error while getting worker info from users table in service layer", logger.Error(err))
		return &pb.Worker{}, err
	}

	branch, err := w.clients.BranchesService().GetBranchTileById(ctx, &pbo.PrimaryKey{Id: branchId})
	if err != nil {
		w.log.Error("error while getting branch title from branches tablein service layer", logger.Error(err))
		return &pb.Worker{}, err
	}

	return &pb.Worker{
		Id:          worker.Id,
		WorkerId:    worker.WorkerId,
		BranchTitle: branch.Title,
		PhoneNumber: user.PhoneNumber,
		FullName:    user.FullName,
		Birthday:    user.Birthday,
		UserRole:    user.UserRole,
		CreatedAt:   worker.CreatedAt,
		UpdatedAt:   worker.UpdatedAt,
	}, nil
}

func (w *workersOfBranchesService) GetByWorkerId(ctx context.Context, request *pb.WorkerId) (*pb.Worker, error) {

	worker, branchId, err := w.storage.WorkersOfBranches().GetByWorkerId(ctx, request)
	if err != nil {
		w.log.Error("error while getting worker info from workers_of_branches table by worker id in service layer", logger.Error(err))
		return &pb.Worker{}, err
	}

	user, err := w.storage.Users().GetById(ctx, &pb.PrimaryKey{Id: worker.Id})
	if err != nil {
		w.log.Error("error while getting worker info from users table in service layer", logger.Error(err))
		return &pb.Worker{}, err
	}

	branch, err := w.clients.BranchesService().GetBranchTileById(ctx, &pbo.PrimaryKey{Id: branchId})
	if err != nil {
		w.log.Error("error while getting branch title from branches table in service layer", logger.Error(err))
		return &pb.Worker{}, err
	}

	return &pb.Worker{
		Id:          worker.Id,
		WorkerId:    worker.WorkerId,
		BranchTitle: branch.Title,
		PhoneNumber: user.PhoneNumber,
		FullName:    user.FullName,
		Birthday:    user.Birthday,
		UserRole:    user.UserRole,
		CreatedAt:   worker.CreatedAt,
		UpdatedAt:   worker.UpdatedAt,
	}, nil
}

func (w *workersOfBranchesService) GetAll(ctx context.Context, request *pb.WorkerFilter) (*pb.Workers, error) {

	workers, branches, err := w.storage.WorkersOfBranches().GetAll(ctx, request)
	if err != nil {
		w.log.Error("error while getting all workers info from workers_of_branches table in service layer", logger.Error(err))
		return &pb.Workers{}, err
	}

	resp := pb.Workers{}

	for ind, val := range workers.Workers {
		user, err := w.storage.Users().GetById(ctx, &pb.PrimaryKey{Id: val.Id})
		if err != nil {
			w.log.Error("error while getting worker info from users table in service layer", logger.Error(err))
			return &pb.Workers{}, err
		}

		if request.UserRole == "" {
			continue
		}

		branch, err := w.clients.BranchesService().GetBranchTileById(ctx, &pbo.PrimaryKey{Id: branches[ind]})
		if err != nil {
			w.log.Error("error while getting branch title from branches table in service layer", logger.Error(err))
			return &pb.Workers{}, err
		}

		if len(resp.Workers) < (int(request.GetPage()-1) * int(request.GetLimit())) {
			worker := &pb.Worker{
				Id:          val.Id,
				WorkerId:    val.WorkerId,
				BranchTitle: branch.Title,
				PhoneNumber: user.PhoneNumber,
				FullName:    user.FullName,
				Birthday:    user.Birthday,
				UserRole:    user.UserRole,
				CreatedAt:   val.CreatedAt,
				UpdatedAt:   val.UpdatedAt,
			}

			resp.Workers = append(resp.Workers, worker)
		}
		resp.Count++
	}

	return &resp, nil
}

func (w *workersOfBranchesService) Update(ctx context.Context, request *pb.Worker) (*pb.Worker, error) {

}

func (w *workersOfBranchesService) Delete(ctx context.Context, request *pb.WorkerId) (*pb.Void, error) {

}

func (w *workersOfBranchesService) CheckWorkerExists(ctx context.Context, request *pb.WorkerId) (*pb.Void, error) {

}
