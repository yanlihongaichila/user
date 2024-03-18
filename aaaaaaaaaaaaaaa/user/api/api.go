package api

import (
	"github.com/JobNing/message/user"
	"google.golang.org/grpc"
)

func Register(r grpc.ServiceRegistrar) {
	user.RegisterUserServer(r, UserService{})
}
