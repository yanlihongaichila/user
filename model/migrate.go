package model

import (
	"fmt"
	"github.com/yanlihongaichila/framework/mysql"
)

// 自动创建表
func MigrateTable() error {
	fmt.Println(mysql.Db)
	return mysql.Db.AutoMigrate(new(User))
}
