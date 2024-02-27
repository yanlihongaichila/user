package model

import (
	"github.com/yanlihongaichila/framework/mysql"
	"gorm.io/gorm"
)

// 创建表 并且对表进行操作
type User struct {
	gorm.Model
	Username string `gorm:"index"`
	Password string `gorm:"type:varchar(100)"`
	Mobile   string `gorm:"type:char(11)"`
	Age      int64  `gorm:"type:tinyint(3)"`
	Sex      int64  `gorm:"type:tinyint(1)"`
	Address  string `gorm:"type:varchar(1024)"`
}

// new User
func NewUser() *User {
	return new(User)
}

//对表的操作

func (u *User) GetUser(id int) (info *User, err error) {
	info = NewUser()
	err = mysql.Db.Where("id=?", id).First(info).Error
	return
}

func (u *User) GetUsers(offset int64, limit int64) (infos []User, total int64, err error) {

	err = mysql.Db.Offset(int(offset)).Limit(int(limit)).Find(&infos).Error
	if err != nil {
		return nil, 0, err
	}
	err = mysql.Db.Model(u).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	return infos, total, nil
}

func (u *User) GetUsersByUsername(name string) (info *User, err error) {
	info = NewUser()
	err = mysql.Db.Where("username=?", name).First(info).Error
	if err != nil {
		return nil, err
	}
	return info, err
}

func (u *User) CreatedUser(in *User) (info *User, err error) {
	info = NewUser()
	err = mysql.Db.Create(in).Error
	if err != nil {
		return nil, err
	}
	info = in
	return
}

func (u *User) UpdatedUser(in *User) (info *User, err error) {
	err = mysql.Db.Where("id = ?", in.ID).Updates(in).Error
	if err != nil {
		return nil, err
	}

	return in, nil

}

func (u *User) DeletedUser(id int64) error {
	err := mysql.Db.Where("id = ?", id).Delete(u).Error
	if err != nil {
		return err
	}
	return nil
}
