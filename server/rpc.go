package server

import (
	"github.com/yanlihongaichila/proto/user"
	"gorm.io/gorm"
	"user/model"
)

//添加

func CreatedUser(in *user.UserInfo) (*user.UserInfo, error) {
	newUser := model.NewUser()
	createdUser, err := newUser.CreatedUser(pbToMysql(in))
	if err != nil {
		return nil, err
	}

	return mysqlToPb(createdUser), nil
}

func pbToMysql(in *user.UserInfo) *model.User {

	return &model.User{
		Model: gorm.Model{
			ID: uint(in.ID),
		},
		Username: in.Username,
		Password: in.Password,
		Mobile:   in.Mobile,
		Age:      in.Age,
		Sex:      int64(in.Sex),
		Address:  in.Address,
	}
}

func mysqlToPb(in *model.User) *user.UserInfo {
	return &user.UserInfo{
		ID:       int64(in.ID),
		Username: in.Username,
		Password: in.Password,
		Mobile:   in.Mobile,
		Age:      in.Age,
		Sex:      user.Sex(in.Sex),
		Address:  in.Address,
	}
}
