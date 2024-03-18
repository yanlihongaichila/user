package service

import (
	"encoding/json"
	"github.com/JobNing/message/goods"
	"goods/model"
	"gorm.io/gorm"
)

func GetGood(id int64) (*goods.GoodInfo, error) {
	goodsMod := model.NewGoods()
	info, err := goodsMod.Get(id)
	if err != nil {
		return nil, err
	}

	return mysqlToPb(info)
}

func GetGoods(offset, limit, goodType int64) (infos []*goods.GoodInfo, total int64, err error) {
	goodsMod := model.NewGoods()
	goodsInfos, total, err := goodsMod.GetGoodss(offset, limit, goodType)
	if err != nil {
		return nil, 0, err
	}

	for _, val := range goodsInfos {
		info, _ := mysqlToPb(&val)
		infos = append(infos, info)
	}

	return
}

func CreateGood(in *goods.GoodInfo) (*goods.GoodInfo, error) {
	modInfo, err := pbToMysql(in)
	if err != nil {
		return nil, err
	}

	goodsMod := model.NewGoods()
	info, err := goodsMod.Create(modInfo)
	if err != nil {
		return nil, err
	}

	return mysqlToPb(info)
}

func UpdateGood(in *goods.GoodInfo) (*goods.GoodInfo, error) {
	modInfo, err := pbToMysql(in)
	if err != nil {
		return nil, err
	}

	goodsMod := model.NewGoods()
	info, err := goodsMod.Update(modInfo)
	if err != nil {
		return nil, err
	}

	return mysqlToPb(info)
}

func DeleteGood(id int64) error {
	goodsMod := model.NewGoods()
	return goodsMod.Delete(id)
}

func UpdateStock(id, num int64) error {
	goodsMod := model.NewGoods()
	return goodsMod.UpdateStock(id, num)
}

func mysqlToPb(info *model.Goods) (*goods.GoodInfo, error) {
	var images []string
	err := json.Unmarshal([]byte(info.Image), &images)
	if err != nil {
		return nil, err
	}

	return &goods.GoodInfo{
		ID:          int64(info.ID),
		GoodName:    info.GoodName,
		Amount:      info.Amount,
		Stock:       info.Stock,
		Image:       images,
		GoodContent: info.GoodContent,
	}, nil
}

func pbToMysql(in *goods.GoodInfo) (*model.Goods, error) {
	img, err := json.Marshal(in.Image)
	if err != nil {
		return nil, err
	}

	return &model.Goods{
		Model: gorm.Model{
			ID: uint(in.ID),
		},
		GoodName:    in.GoodName,
		Amount:      in.Amount,
		Stock:       in.Stock,
		Image:       string(img),
		GoodContent: in.GoodContent,
	}, nil
}
