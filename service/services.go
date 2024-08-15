package service

import (
	pb "users_service/genproto/users"
	"users_service/grpc/clients"
	"users_service/pkg/logger"
	"users_service/storage"
)

type IServiceManager interface {
	AuthService() pb.AuthServiceServer
	// UsersService() pb.UsersServiceServer
	WorkersOfBranches() pb.WorkersOfBranchesServiceServer
}

type ServiceManager struct {
	storage storage.IStorage
	clients clients.IClientsMeneger
	log     logger.ILogger
}

func NewServiceManager(storage storage.IStorage, clients clients.IClientsMeneger, log logger.ILogger) IServiceManager {
	return &ServiceManager{
		storage: storage,
		clients: clients,
		log:     log,
	}
}

func (s *ServiceManager) AuthService() pb.AuthServiceServer {
	return NewAuthService(s.storage, s.log)
}

// func (s *ServiceManager) UsersService() pb.UsersServiceServer {
// 	return NewUsersService(s.storage, s.log)
// }

func (s *ServiceManager) WorkersOfBranches() pb.WorkersOfBranchesServiceServer {
	return NewWorkersOfBranchesService(s.storage, s.clients, s.log)
}
