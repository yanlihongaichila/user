package api

import (
	"context"
	"github.com/JobNing/message/goods"
	"goods/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GoodsService struct {
	goods.UnimplementedGoodServer
}

func (GoodsService) GetGood(ctx context.Context, in *goods.GetGoodRequest) (*goods.GetGoodResponse, error) {
	if in == nil {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}
	if in.ID == 0 {
		return nil, status.Error(codes.InvalidArgument, "id is required and > 0")
	}

	info, err := service.GetGood(in.ID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &goods.GetGoodResponse{
		Info: info,
	}, nil
}

func (GoodsService) GetGoods(ctx context.Context, in *goods.GetGoodsRequest) (*goods.GetGoodsResponse, error) {
	if in == nil {
		return nil, status.Error(codes.InvalidArgument, "argument is required")
	}
	if in.Limit > 10000 {
		return nil, status.Error(codes.InvalidArgument, "limit must be less than 10000")
	}

	infos, total, err := service.GetGoods(in.Offset, in.Limit, in.Type)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &goods.GetGoodsResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (GoodsService) CreateGood(ctx context.Context, in *goods.CreateGoodRequest) (*goods.CreateGoodResponse, error) {
	if in == nil {
		return nil, status.Error(codes.InvalidArgument, "argument is required")
	}
	if in.Info == nil {
		return nil, status.Error(codes.InvalidArgument, "info is required")
	}
	if in.Info.GoodName == "" {
		return nil, status.Error(codes.InvalidArgument, "goods name is required")
	}
	if in.Info.Amount == "" {
		return nil, status.Error(codes.InvalidArgument, "amount is required")
	}
	info, err := service.CreateGood(in.Info)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &goods.CreateGoodResponse{
		Info: info,
	}, nil
}

func (GoodsService) UpdateGood(ctx context.Context, in *goods.UpdateGoodRequest) (*goods.UpdateGoodResponse, error) {
	if in == nil {
		return nil, status.Error(codes.InvalidArgument, "argument is required")
	}
	if in.Info == nil {
		return nil, status.Error(codes.InvalidArgument, "info is required")
	}
	if in.Info.ID == 0 {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}

	info, err := service.UpdateGood(in.Info)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &goods.UpdateGoodResponse{
		Info: info,
	}, nil
}

func (GoodsService) DeleteGood(ctx context.Context, in *goods.DeleteGoodRequest) (*goods.DeleteGoodResponse, error) {
	if in == nil {
		return nil, status.Error(codes.InvalidArgument, "argument is required")
	}
	if in.ID == 0 {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}

	err := service.DeleteGood(in.ID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &goods.DeleteGoodResponse{}, nil
}

func (GoodsService) UpdateStock(ctx context.Context, in *goods.UpdateStockRequest) (*goods.UpdateStockResponse, error) {
	if in == nil {
		return nil, status.Error(codes.InvalidArgument, "argument is required")
	}
	if in.ID == 0 {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}

	err := service.UpdateStock(in.ID, in.Num)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &goods.UpdateStockResponse{}, nil
}
