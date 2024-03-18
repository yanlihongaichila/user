package api

import (
	"context"
	"fmt"
	"github.com/yanlihongaichila/proto/user"
	"github.com/yanlihongaichila/user/server"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	user.UnimplementedUserServer
}

func (u UserService) GetUser(ctx context.Context, request *user.GetUserRequest) (*user.GetUserResponse, error) {
	//校验参数
	if request == nil {
		return nil, status.Error(codes.InvalidArgument, "argument is required")
	}

	if request.ID < 1 {
		return nil, status.Error(codes.InvalidArgument, "id cannot smaller 1")
	}

	getUser, err := server.GetUser(request.ID)
	if err != nil {
		return nil, err
	}

	return &user.GetUserResponse{User: getUser}, nil
}

func (u UserService) GetUsers(ctx context.Context, request *user.GetUsersRequest) (*user.GetUsersResponse, error) {

	if request.Offset < 1 {
		request.Offset = 0
	}

	if request.Limit > 1000 {
		return nil, status.Error(codes.FailedPrecondition, "limit must not be greater than 1000")
	}

	users, total, err := server.GetUsers(request.Offset, request.Limit)
	if err != nil {
		return nil, err
	}

	return &user.GetUsersResponse{
		Users: users,
		Total: total,
	}, nil
}

func (u UserService) CreatedUser(ctx context.Context, in *user.CreatedUserRequest) (*user.CreatedUserResponse, error) {
	//对传入的值进行校验
	if in == nil {
		return nil, status.Error(codes.InvalidArgument, "argument is required")
	}

	if in.User == nil {
		return nil, status.Error(codes.InvalidArgument, "user is required")
	}

	if in.User.Username == "" {
		return nil, status.Error(codes.InvalidArgument, "username is required")
	}

	if in.User.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "password is required")
	}

	createdUser, err := server.CreatedUser(in.User)
	if err != nil {
		return nil, err
	}

	return &user.CreatedUserResponse{User: createdUser}, nil
}

func (u UserService) UpdatedUser(ctx context.Context, request *user.UpdatedUserRequest) (*user.UpdatedUserResponse, error) {
	if request == nil {
		return nil, status.Error(codes.InvalidArgument, "argument is required")
	}

	if request.User.ID < 1 {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}

	updatedUser, err := server.UpdatedUser(request.User)
	if err != nil {
		return nil, err
	}

	return &user.UpdatedUserResponse{User: updatedUser}, nil
}

func (u UserService) DeletedUser(ctx context.Context, request *user.DeletedUserRequest) (*user.DeletedUserResponse, error) {
	if request.ID < 1 {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}

	err := server.DeletedUser(request.ID)
	if err != nil {
		return nil, err
	}

	return &user.DeletedUserResponse{}, nil
}

func (u UserService) GetUserByUsername(ctx context.Context, request *user.GetUserByUsernameRequest) (*user.GetUserByUsernameResponse, error) {
	fmt.Println("1111111111111111111")
	fmt.Println(request.Username)
	if request.Username == "" {
		return nil, status.Error(codes.InvalidArgument, "username is required")
	}

	username, err := server.GetUsersByUsername(request.Username)
	if err != nil {
		return nil, err
	}
	return &user.GetUserByUsernameResponse{User: username}, nil
}

func (u UserService) mustEmbedUnimplementedUserServer() {
	//TODO implement me
	panic("implement me")
}
