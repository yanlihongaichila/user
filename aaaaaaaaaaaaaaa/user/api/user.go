package api

import (
	"context"
	"fmt"
	"github.com/JobNing/message/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"user/service"
)

type UserService struct {
	user.UnimplementedUserServer
}

func (UserService) GetUser(ctx context.Context, in *user.GetUserRequest) (*user.GetUserResponse, error) {
	fmt.Println("**************************")
	if in == nil {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}
	if in.ID == 0 {
		return nil, status.Error(codes.InvalidArgument, "id is required and > 0")
	}

	info, err := service.GetUser(in.ID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &user.GetUserResponse{
		Info: info,
	}, nil
}

func (UserService) GetUsers(ctx context.Context, in *user.GetUsersRequest) (*user.GetUsersResponse, error) {
	if in == nil {
		return nil, status.Error(codes.InvalidArgument, "argument is required")
	}
	if in.Limit > 10000 {
		return nil, status.Error(codes.InvalidArgument, "limit must be less than 10000")
	}

	infos, total, err := service.GetUsers(in.Offset, in.Limit)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &user.GetUsersResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (UserService) GetUserByUsername(ctx context.Context, in *user.GetUserByUsernameRequest) (*user.GetUserByUsernameResponse, error) {
	if in == nil {
		return nil, status.Error(codes.InvalidArgument, "argument is required")
	}
	if in.Username == "" {
		return nil, status.Error(codes.InvalidArgument, "username is required")
	}

	info, err := service.GetByUsername(in.Username)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &user.GetUserByUsernameResponse{
		Info: info,
	}, nil
}

func (UserService) GetUserByMobile(ctx context.Context, in *user.GetUserByMobileRequest) (*user.GetUserByMobileResponse, error) {
	if in == nil {
		return nil, status.Error(codes.InvalidArgument, "argument is required")
	}
	if in.Mobile == "" {
		return nil, status.Error(codes.InvalidArgument, "username is required")
	}

	info, err := service.GetByMobile(in.Mobile)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &user.GetUserByMobileResponse{
		Info: info,
	}, nil
}

func (UserService) CreateUser(ctx context.Context, in *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	if in == nil {
		return nil, status.Error(codes.InvalidArgument, "argument is required")
	}
	if in.Info == nil {
		return nil, status.Error(codes.InvalidArgument, "info is required")
	}
	if in.Info.Mobile == "" {
		return nil, status.Error(codes.InvalidArgument, "mobile is required")
	}
	if in.Info.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "username is required")
	}

	info, err := service.CreateUser(in.Info)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &user.CreateUserResponse{
		Info: info,
	}, nil
}

func (UserService) UpdateUser(ctx context.Context, in *user.UpdateUserRequest) (*user.UpdateUserResponse, error) {
	if in == nil {
		return nil, status.Error(codes.InvalidArgument, "argument is required")
	}
	if in.Info == nil {
		return nil, status.Error(codes.InvalidArgument, "info is required")
	}
	if in.Info.ID == 0 {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}

	info, err := service.UpdateUser(in.Info)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &user.UpdateUserResponse{
		Info: info,
	}, nil
}

func (UserService) DeleteUser(ctx context.Context, in *user.DeleteUserRequest) (*user.DeleteUserResponse, error) {
	if in == nil {
		return nil, status.Error(codes.InvalidArgument, "argument is required")
	}
	if in.ID == 0 {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}

	err := service.DeleteUser(in.ID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &user.DeleteUserResponse{}, nil
}
