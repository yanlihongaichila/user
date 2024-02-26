package main

import (
	"flag"
	"github.com/JobNing/framework/grpc"
	"gorm.io/gorm"
)

func main() {
	//连接数据库
	//flag.Parse()
	//err := app.Init("mysql")
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = model.MigrateTable()
	//if err != nil {
	//	panic(err)
	//}
	//
	//gprc.ConcentGrpc(8077, func(s *grpc.Server) {
	//
	//})
	//if err != nil {
	//	panic(err)
	//}
	flag.Parse()
	err := app.Init("mysql")
	if err != nil {
		panic(err)
	}

	mysql.WithTX(func(tx *gorm.DB) error {
		err := tx.Table("order").Where("id = ?", 63).Update("user_id", 90).Error
		if err != nil {
			return err
		}

		return tx.Table("good9s").Where("id = ?", 1).Update("name", "西瓜").Error
	})

	err = grpc.RegisterGRPC(*port, func(s *grpc2.Server) {
		api.Register(s)
	})
	if err != nil {
		panic(err)
	}
}
