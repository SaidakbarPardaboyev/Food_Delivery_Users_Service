package service

import (
	"context"
	pb "users_service/genproto/users"
	"users_service/pkg/logger"
	"users_service/storage"
)

type userService struct {
	storage storage.IStorage
	log     logger.ILogger
	pb.UnimplementedUsersServiceServer
}

func NewUsersService(storage storage.IStorage, log logger.ILogger) *userService {
	return &userService{
		storage: storage,
		log:     log,
	}
}

func (u *userService) GetById(ctx context.Context, request *pb.PrimaryKey) (*pb.User, error) {

	resp, err := u.storage.Users().GetById(ctx, request)
	if err != nil {
		u.log.Error("error while getting user info in service layer", logger.Error(err))
		return &pb.User{}, err
	}

	return resp, nil
}

func (u *userService) GetAll(ctx context.Context, request *pb.GetListRequest) (*pb.Users, error) {

	resp, err := u.storage.Users().GetAll(ctx, request)
	if err != nil {
		u.log.Error("error while getting all users info in service layer", logger.Error(err))
		return &pb.Users{}, err
	}

	return resp, nil
}

func (u *userService) Update(ctx context.Context, request *pb.UpdateUser) (*pb.User, error) {

	resp, err := u.storage.Users().Update(ctx, request)
	if err != nil {
		u.log.Error("error while updating user info in service layer", logger.Error(err))
		return &pb.User{}, err
	}

	return resp, nil
}

func (u *userService) Delete(ctx context.Context, request *pb.PrimaryKey) (*pb.Void, error) {

	resp, err := u.storage.Users().Delete(ctx, request)
	if err != nil {
		u.log.Error("error while deleting user info in service layer", logger.Error(err))
		return &pb.Void{}, err
	}

	return resp, nil
}

func (u *userService) ChangeUserRole(ctx context.Context, request *pb.ChangeUserRole) (*pb.Void, error) {

	resp, err := u.storage.Users().ChangeUserRole(ctx, request)
	if err != nil {
		u.log.Error("error while changing user role in service layer", logger.Error(err))
		return &pb.Void{}, err
	}

	return resp, nil
}

func (u *userService) CheckUserIdExists(ctx context.Context, in *pb.PrimaryKey) (*pb.Void, error) {

	resp, err := u.storage.Users().CheckUserIdExists(ctx, in)
	if err != nil {
		u.log.Error("error while check user id in service layer", logger.Error(err))
		return &pb.Void{}, err
	}

	return resp, nil
}
