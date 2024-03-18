package api

import (
	"github.com/JobNing/message/goods"
	"google.golang.org/grpc"
)

func Register(r grpc.ServiceRegistrar) {
	goods.RegisterGoodServer(r, GoodsService{})
}
