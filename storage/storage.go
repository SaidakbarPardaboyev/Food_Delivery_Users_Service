package storage

import (
	"context"
	"users_service/configs"
	"users_service/pkg/logger"
	"users_service/storage/postgres"

	pb "users_service/genproto/users"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	dbPostgres *pgxpool.Pool
	log        logger.ILogger
}

type IStorage interface {
	Close()
	Auth() IAuthStorage
	Users() IUsersStorage
	WorkersOfBranches() IWorkersOfBranches
	UserLocation() IUserLocationRepo
}

type IAuthStorage interface {
	Create(context.Context, *pb.CreateUser) (*pb.User, error)
	GetByPhone(context.Context, *pb.Phone) (*pb.User, error)
	CheckRefreshTokenExists(context.Context, *pb.RequestRefreshToken) (*pb.Void, error)
	StoreRefreshToken(context.Context, *pb.RefreshToken) (*pb.Void, error)
	DeleteRefreshTokenByUserId(context.Context, *pb.PrimaryKey) (*pb.Void, error)
}

type IUsersStorage interface {
	GetById(context.Context, *pb.PrimaryKey) (*pb.User, error)
	GetAll(context.Context, *pb.GetListRequest) (*pb.Users, error)
	Update(context.Context, *pb.UpdateUser) (*pb.User, error)
	Delete(context.Context, *pb.PrimaryKey) (*pb.Void, error)
	ChangeUserRole(context.Context, *pb.ChangeUserRole) (*pb.Void, error)
	CheckUserIdExists(context.Context, *pb.PrimaryKey) (*pb.Void, error)
}

type IWorkersOfBranches interface {
	Create(ctx context.Context, request *pb.CreateWorker) (*pb.WorkerId, error)
	GetByUUID(ctx context.Context, request *pb.PrimaryKey) (*pb.Worker, string, error)
	GetByWorkerId(ctx context.Context, request *pb.WorkerId) (*pb.Worker, string, error)
	GetAll(ctx context.Context, request *pb.WorkerFilter) (*pb.Workers, []string, error)
	Update(ctx context.Context, request *pb.UpdateWorker) (*pb.Worker, string, error)
	Delete(ctx context.Context, request *pb.WorkerId) (*pb.Void, error)
	CheckWorkerExists(ctx context.Context, request *pb.WorkerId) (*pb.Void, error)
}

type IUserLocationRepo interface {
	Create(context.Context, *pb.CreateUserLocation) (*pb.UserLocation, error)
	GetById(context.Context, *pb.PrimaryKey) (*pb.UserLocation, error)
	GetByUserId(context.Context, *pb.PrimaryKey) (*pb.UserLocation, error)
	GetAll(context.Context, *pb.UserLocationFilter) (*pb.UserLocations, error)
	Update(context.Context, *pb.UpdateUserLocation) (*pb.UserLocation, error)
	Delete(context.Context, *pb.PrimaryKey) (*pb.Void, error)
}

func New(ctx context.Context, cfg *configs.Config, log *logger.ILogger) (IStorage, error) {
	dbPostgres, err := postgres.ConnectDB(ctx, cfg)
	if err != nil {
		return nil, err
	}

	return &Storage{
		dbPostgres: dbPostgres,
		log:        *log,
	}, nil
}

func (s *Storage) Close() {
	s.dbPostgres.Close()
}

func (s *Storage) Auth() IAuthStorage {
	return postgres.NewAuthRepo(s.dbPostgres, s.log)
}

func (s *Storage) Users() IUsersStorage {
	return postgres.NewUsersRepo(s.dbPostgres, s.log)
}

func (s *Storage) WorkersOfBranches() IWorkersOfBranches {
	return postgres.NewWorkersOfBranchesRepo(s.dbPostgres, s.log)
}

func (s *Storage) UserLocation() IUserLocationRepo {
	return postgres.NewUserLoactionRepo(s.dbPostgres, s.log)
}
