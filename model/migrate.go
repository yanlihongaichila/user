package model

import "github.com/yanlihongaichila/framework/mysql"

// 自动创建表
func MigrateTable() error {
	return mysql.Db.AutoMigrate(&User{})
}
