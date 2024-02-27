package server

import (
	"github.com/yanlihongaichila/proto/user"
	"gorm.io/gorm"
	"user/model"
)

/*
************************************************************
*******************---Search---*****************************
************************************************************
 */
func GetUser(id int64) (*user.UserInfo, error) {
	newUser := model.NewUser()
	getUser, err := newUser.GetUser(int(id))
	if err != nil {
		return nil, err
	}
	pb := mysqlToPb(getUser)
	return pb, nil
}

func GetUsers(offset, limit int64) ([]*user.UserInfo, int64, error) {
	newUser := model.NewUser()
	searchUsers := []*user.UserInfo{}
	users, total, err := newUser.GetUsers(offset, limit)
	if err != nil {
		return nil, 0, err
	}

	for _, val := range users {
		searchUsers = append(searchUsers, mysqlToPb(&val))
	}
	return searchUsers, total, nil

}

func GetUsersByUsername(username string) (*user.UserInfo, error) {
	newUser := model.NewUser()
	byUsername, err := newUser.GetUsersByUsername(username)
	if err != nil {
		return nil, err
	}

	pb := mysqlToPb(byUsername)

	return pb, nil
}

/*
************************************************************
*******************---Created---****************************
************************************************************
 */
func CreatedUser(in *user.UserInfo) (*user.UserInfo, error) {
	newUser := model.NewUser()
	createdUser, err := newUser.CreatedUser(pbToMysql(in))
	if err != nil {
		return nil, err
	}

	return mysqlToPb(createdUser), nil
}

/*
************************************************************
*******************---Updated---*****************************
************************************************************
 */

func UpdatedUser(in *user.UserInfo) (*user.UserInfo, error) {
	newUser := model.NewUser()
	updatedUser, err := newUser.UpdatedUser(pbToMysql(in))
	if err != nil {
		return nil, err
	}

	return mysqlToPb(updatedUser), err
}

/*
************************************************************
*******************---Deleted---*****************************
************************************************************
 */
func DeletedUser(id int64) error {
	newUser := model.NewUser()
	err := newUser.DeletedUser(id)
	if err != nil {
		return err
	}
	return nil
}

/*
************************************************************
*******************---Other---*****************************
************************************************************
 */
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
