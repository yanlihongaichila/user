package api

import (
	"context"
	"github.com/yanlihongaichila/proto/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"user/server"
)

type UserService struct {
	user.UnimplementedUserServer
}

func (UserService) CreatedUser(ctx context.Context, in *user.CreatedUserRequest) (*user.CreatedUserResponse, error) {
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
