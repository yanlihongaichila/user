package main

import (
	"flag"
	"github.com/JobNing/frameworkJ/app"
	"github.com/JobNing/frameworkJ/config"
	"github.com/JobNing/frameworkJ/grpc"
	"github.com/spf13/viper"
	"goods/api"
	"goods/consts"
	"goods/model"
	grpc2 "google.golang.org/grpc"
)

func main() {
	flag.Parse()
	if err := config.InitViper("config", "./config"); err != nil {
		panic(err)
	}

	err := app.Init(
		consts.ServiceName,
		viper.GetString("nacos.ip"),
		viper.GetString("nacos.port"),
		"mysql",
	)
	if err != nil {
		panic(err)
	}

	err = model.MigrateTable()
	if err != nil {
		panic(err)
	}

	err = grpc.RegisterGRPC(viper.GetString("nacos.group"), consts.ServiceName, func(s *grpc2.Server) {
		api.Register(s)
	})
	panic(err)
}
