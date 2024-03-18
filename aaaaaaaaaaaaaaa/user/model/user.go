package model

import (
	"fmt"
	"github.com/JobNing/frameworkJ/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"index"`
	Password string `gorm:"type:varchar(100)"`
	Mobile   string `gorm:"type:char(11)"`
	Sex      int    `gorm:"type:tinyint(1)"`
	Age      int    `gorm:"type:tinyint(3)"`
	Address  string `gorm:"type:varchar(1024)"`
}

func NewUser() *User {
	return new(User)
}

func (m *User) Get(id int64) (info *User, err error) {
	info = new(User)
	err = mysql.DB.Where("id = ?", id).First(info).Error
	return
}

func (m *User) GetByUsername(username string) (info *User, err error) {
	info = new(User)
	err = mysql.DB.Where("username = ?", username).First(info).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return
}

func (m *User) GetByMobile(mobile string) (info *User, err error) {
	info = new(User)
	err = mysql.DB.Where("mobile = ?", mobile).First(info).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return
}

func (m *User) GetUsers(offset, limit int64) ([]User, int64, error) {
	var infos []User
	var total int64
	err := mysql.DB.Offset(int(offset)).Limit(int(limit)).Find(&infos).Error
	if err != nil {
		return nil, 0, err
	}
	err = mysql.DB.Model(m).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	return infos, total, nil
}

func (m *User) Count(offset, limit int64) ([]User, error) {
	var infos []User
	fmt.Println(mysql.DB)
	err := mysql.DB.Offset(int(offset)).Limit(int(limit)).Find(&infos).Error
	if err != nil {
		return nil, err
	}
	return infos, nil
}

func (m *User) Create(in *User) (info *User, err error) {
	err = mysql.DB.Create(in).Error
	return in, err
}

func (m *User) Update(in *User) (info *User, err error) {
	err = mysql.DB.Model(m).Where("id =?", in.ID).Updates(in).Error
	return in, err
}

func (m *User) Delete(id int64) error {
	return mysql.DB.Where("id = ?", id).Delete(m).Error
}
