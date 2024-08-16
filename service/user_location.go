package service

import (
	"context"
	"users_service/pkg/logger"
	"users_service/storage"

	pb "users_service/genproto/users"
)

type userLocationService struct {
	storage storage.IStorage
	log     logger.ILogger
	pb.UnimplementedUserLocationServiceServer
}

func NewUserLocationService(storage storage.IStorage, log logger.ILogger) *userLocationService {
	return &userLocationService{
		storage: storage,
		log:     log,
	}
}

func (s *userLocationService) Create(ctx context.Context, req *pb.CreateUserLocation) (*pb.UserLocation, error) {
	res, err := s.storage.UserLocation().Create(ctx, req)
	if err != nil {
		s.log.Error("Error creating user location", logger.Error(err))
		return nil, err
	}
	return res, nil
}

func (s *userLocationService) GetById(ctx context.Context, req *pb.PrimaryKey) (*pb.UserLocation, error) {
	res, err := s.storage.UserLocation().GetById(ctx, req)
	if err != nil {
		s.log.Error("Error getting user location by ID", logger.Error(err))
		return nil, err
	}
	return res, nil
}

func (s *userLocationService) GetByUserId(ctx context.Context, req *pb.PrimaryKey) (*pb.UserLocation, error) {
	res, err := s.storage.UserLocation().GetById(ctx, req)
	if err != nil {
		s.log.Error("Error getting user location by ID", logger.Error(err))
		return nil, err
	}
	return res, nil
}

func (s *userLocationService) GetAll(ctx context.Context, req *pb.UserLocationFilter) (*pb.UserLocations, error) {
	res, err := s.storage.UserLocation().GetAll(ctx, req)
	if err != nil {
		s.log.Error("Error getting all user locations", logger.Error(err))
		return nil, err
	}
	return res, nil
}

func (s *userLocationService) Update(ctx context.Context, req *pb.UpdateUserLocation) (*pb.UserLocation, error) {
	res, err := s.storage.UserLocation().Update(ctx, req)
	if err != nil {
		s.log.Error("Error updating user location", logger.Error(err))
		return nil, err
	}
	return res, nil
}

func (s *userLocationService) Delete(ctx context.Context, req *pb.PrimaryKey) (*pb.Void, error) {
	_, err := s.storage.UserLocation().Delete(ctx, req)
	if err != nil {
		s.log.Error("Error deleting user location", logger.Error(err))
		return nil, err
	}
	return &pb.Void{}, nil
}
