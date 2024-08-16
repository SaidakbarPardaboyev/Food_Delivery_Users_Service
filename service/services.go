package service

import (
	pb "users_service/genproto/users"
	"users_service/pkg/logger"
	"users_service/storage"
)

type IServiceManager interface {
	AuthService() pb.AuthServiceServer
	UsersService() pb.UsersServiceServer
	UserLocation() pb.UserLocationServiceServer
}

type ServiceManager struct {
	storage storage.IStorage
	log     logger.ILogger
}

func NewServiceManager(storage storage.IStorage, log logger.ILogger) IServiceManager {
	return &ServiceManager{
		storage: storage,
		log:     log,
	}
}

func (s *ServiceManager) AuthService() pb.AuthServiceServer {
	return NewAuthService(s.storage, s.log)
}

func (s *ServiceManager) UsersService() pb.UsersServiceServer {
	return NewUsersService(s.storage, s.log)
}

func (s *ServiceManager) UserLocation() pb.UserLocationServiceServer {
	return NewUserLocationService(s.storage, s.log)
}
