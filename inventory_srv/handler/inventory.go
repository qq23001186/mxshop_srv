package handler

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"nd/inventory_srv/global"
	"nd/inventory_srv/model"
	"nd/inventory_srv/proto"
)

type InventoryServer struct {
	proto.UnimplementedInventoryServer
}

func (*InventoryServer) SetInv(ctx context.Context, req *proto.GoodsInvInfo) (*emptypb.Empty, error) {
	//设置库存， 如果我要更新库存
	var inv model.Inventory
	global.DB.Where(&model.Inventory{Goods: req.GoodsId}).First(&inv)
	inv.Goods = req.GoodsId
	inv.Stocks = req.Num

	global.DB.Save(&inv)
	return &emptypb.Empty{}, nil
}

func (*InventoryServer) InvDetail(ctx context.Context, req *proto.GoodsInvInfo) (*proto.GoodsInvInfo, error) {
	var inv model.Inventory
	if result := global.DB.Where(&model.Inventory{Goods: req.GoodsId}).First(&inv); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "没有库存信息")
	}
	return &proto.GoodsInvInfo{
		GoodsId: inv.Goods,
		Num:     inv.Stocks,
	}, nil
}

func (*InventoryServer) Sell(ctx context.Context, req *proto.SellInfo) (*emptypb.Empty, error) {
	// 扣减库存，本地事务
	// 数据库基本的一个应用场景：数据库事务
	// 并发情况之下 可能会出现超卖 1
	tx := global.DB.Begin()
	for _, goodInfo := range req.GoodsInfo {
		var inv model.Inventory
		if result := global.DB.Where(&model.Inventory{Goods: goodInfo.GoodsId}).First(&inv); result.RowsAffected == 0 {
			tx.Rollback() // 回滚之前的操作
			return nil, status.Errorf(codes.InvalidArgument, "没有库存信息")
		}
		// 判断库存是否充足
		if inv.Stocks < goodInfo.Num {
			tx.Rollback() // 回滚之前的操作
			return nil, status.Errorf(codes.ResourceExhausted, "库存不足")
		}
		// 扣减，这里会出现数据不一致的问题
		inv.Stocks -= goodInfo.Num
		tx.Save(&inv) // 一旦使用了事务的，保存修改数据库的操作就需要使用事务的tx，而不能使用db
	}
	tx.Commit() // 需要自己手动提交操作
	return &emptypb.Empty{}, nil
}

func (*InventoryServer) Reback(ctx context.Context, req *proto.SellInfo) (*emptypb.Empty, error) {
	//库存归还： 1：订单超时归还 2. 订单创建失败，归还之前扣减的库存 3. 手动归还
	tx := global.DB.Begin()
	for _, goodInfo := range req.GoodsInfo {
		var inv model.Inventory
		if result := global.DB.Where(&model.Inventory{Goods: goodInfo.GoodsId}).First(&inv); result.RowsAffected == 0 {
			tx.Rollback() //回滚之前的操作
			return nil, status.Errorf(codes.InvalidArgument, "没有库存信息")
		}

		//扣减， 会出现数据不一致的问题 - 锁，分布式锁
		inv.Stocks += goodInfo.Num
		tx.Save(&inv)
	}
	tx.Commit() // 需要自己手动提交操作
	return &emptypb.Empty{}, nil
}
