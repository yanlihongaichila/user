package main

import (
	"flag"
	"github.com/yanlihongaichila/framework/app"
	"github.com/yanlihongaichila/framework/gprc"
	"github.com/yanlihongaichila/user/api"
	"github.com/yanlihongaichila/user/model"
	"google.golang.org/grpc"
)

func main() {
	//连接数据库
	flag.Parse()
	err := app.Init("config", "./config", "nacos", "mysql")
	if err != nil {
		panic("failed to Mysql")
	}
	//自动建表
	err = model.MigrateTable()
	if err != nil {
		panic(err)
	}

	//gprc.ConcentGrpc(8077, func(s *grpc.Server) {
	//	api.RegisterUser(s)
	//})

	err = gprc.ConcentGrpc("grpc", func(s *grpc.Server) {
		api.RegisterUser(s)
	})
	if err != nil {
		return
	}

}
