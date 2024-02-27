package main

import (
	"flag"
	"github.com/yanlihongaichila/framework/app"
	"github.com/yanlihongaichila/framework/gprc"
	"google.golang.org/grpc"
	"user/api"
	"user/model"
)

func main() {
	//连接数据库
	flag.Parse()
	err := app.Init("mysql")
	if err != nil {
		panic("failed to Mysql")
	}
	err = model.MigrateTable()
	if err != nil {
		panic(err)
	}

	gprc.ConcentGrpc(8077, func(s *grpc.Server) {
		api.RegisterUser(s)
	})
	if err != nil {
		panic(err)
	}

}
