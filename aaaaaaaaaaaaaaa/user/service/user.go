package service

import (
	"github.com/JobNing/message/user"
	"gorm.io/gorm"
	"user/model"
)

func GetUser(id int64) (*user.UserInfo, error) {
	userMod := model.NewUser()
	info, err := userMod.Get(id)
	if err != nil {
		return nil, err
	}

	return mysqlToPb(info)
}

func GetByUsername(username string) (*user.UserInfo, error) {
	userMod := model.NewUser()
	info, err := userMod.GetByUsername(username)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}
	return mysqlToPb(info)
}

func GetByMobile(username string) (*user.UserInfo, error) {
	userMod := model.NewUser()
	info, err := userMod.GetByMobile(username)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}
	return mysqlToPb(info)
}

func GetUsers(offset, limit int64) (infos []*user.UserInfo, total int64, err error) {
	userMod := model.NewUser()
	userInfos, total, err := userMod.GetUsers(offset, limit)
	if err != nil {
		return nil, 0, err
	}

	for _, val := range userInfos {
		info, _ := mysqlToPb(&val)
		infos = append(infos, info)
	}

	return
}

func CreateUser(in *user.UserInfo) (*user.UserInfo, error) {
	userMod := model.NewUser()
	info, err := userMod.Create(pbToMysql(in))
	if err != nil {
		return nil, err
	}

	return mysqlToPb(info)
}

func UpdateUser(in *user.UserInfo) (*user.UserInfo, error) {
	userMod := model.NewUser()
	info, err := userMod.Update(pbToMysql(in))
	if err != nil {
		return nil, err
	}

	return mysqlToPb(info)
}

func DeleteUser(id int64) error {
	userMod := model.NewUser()
	return userMod.Delete(id)
}

func mysqlToPb(info *model.User) (*user.UserInfo, error) {
	return &user.UserInfo{
		ID:       int64(info.ID),
		Username: info.Username,
		Mobile:   info.Mobile,
		Age:      int64(info.Age),
		Sex:      user.Sex(info.Sex),
		Address:  info.Address,
	}, nil
}

func pbToMysql(in *user.UserInfo) *model.User {
	return &model.User{
		Model: gorm.Model{
			ID: uint(in.ID),
		},
		Username: in.Username,
		Password: in.Password,
		Mobile:   in.Mobile,
		Sex:      int(in.Sex.Number()),
		Age:      int(in.Age),
		Address:  in.Address,
	}
}
