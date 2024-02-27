package api

import (
	"github.com/yanlihongaichila/proto/user"
	"google.golang.org/grpc"
)

func RegisterUser(g grpc.ServiceRegistrar) {
	user.RegisterUserServer(g, UserService{})
}
