package model

import (
	"github.com/JobNing/frameworkJ/mysql"
)

func MigrateTable() error {
	return mysql.DB.AutoMigrate(&User{})
}
