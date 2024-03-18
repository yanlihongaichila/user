package model

import (
	"context"
	"fmt"
	"github.com/JobNing/frameworkJ/mysql"
	"github.com/JobNing/frameworkJ/redis"
	"gorm.io/gorm"
	"time"
)

type Goods struct {
	gorm.Model
	GoodName    string `gorm:"index"`
	GoodContent string `gorm:"index"`
	Amount      string `gorm:"type:decimal(10,2)"`
	Stock       int64  `gorm:"type:int(11)"`
	Image       string `gorm:"type:text(0)"`
	Type        int64  `gorm:"type:tinyint(1)"`
}

func NewGoods() *Goods {
	return new(Goods)
}

func (m *Goods) Get(id int64) (info *Goods, err error) {
	info = new(Goods)
	err = mysql.DB.Where("id = ?", id).First(info).Error
	return
}

func (m *Goods) GetGoodss(offset, limit, goodType int64) ([]Goods, int64, error) {
	var infos []Goods
	var total int64
	mod := mysql.DB.Model(m).Debug()
	if goodType > 0 {
		mod = mod.Where("type = ?", goodType)
	}
	err := mod.Offset(int(offset)).Limit(int(limit)).Find(&infos).Error
	if err != nil {
		return nil, 0, err
	}

	err = mysql.DB.Model(m).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	return infos, total, nil
}

func (m *Goods) Count(offset, limit int64) ([]Goods, error) {
	var infos []Goods
	fmt.Println(mysql.DB)
	err := mysql.DB.Offset(int(offset)).Limit(int(limit)).Find(&infos).Error
	if err != nil {
		return nil, err
	}
	return infos, nil
}

func (m *Goods) Create(in *Goods) (info *Goods, err error) {
	err = mysql.DB.Create(in).Error
	return in, err
}

func (m *Goods) Update(in *Goods) (info *Goods, err error) {
	err = mysql.DB.Model(m).Where("id =?", in.ID).Updates(in).Error
	return in, err
}

func (m *Goods) Delete(id int64) error {
	return mysql.DB.Where("id = ?", id).Delete(m).Error
}

//var mutex = sync.Mutex{}
//
//func (m *Goods) UpdateStock(id, num int64) error {
//	mutex.Lock()
//	fmt.Println("**************123")
//	//time.Sleep(time.Second * 5)
//	defer mutex.Unlock()
//	goodInfo, err := m.Get(id)
//	if err != nil {
//		return err
//	}
//	num = goodInfo.Stock + num
//
//	return mysql.DB.Model(m).Where("id = ?", id).Update("stock", num).Error
//}

func (m *Goods) UpdateStock(id, num int64) error {
	_, err := redis.Lock(context.Background(), "goods", "update:stock", 1, time.Second*10, false)
	if err != nil {
		return err
	}

	defer redis.UnLock(context.Background(), "goods", "update:stock")
	goodInfo, err := m.Get(id)
	if err != nil {
		return err
	}
	num = goodInfo.Stock + num

	return mysql.DB.Model(m).Where("id = ?", id).Update("stock", num).Error
}
