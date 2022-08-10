package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"nd/inventory_srv/proto"
	"nd/inventory_srv/tests"
)

var invClient proto.InventoryClient
var conn *grpc.ClientConn

func Init() {
	var err error
	conn, err = grpc.Dial(tests.TargetAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	invClient = proto.NewInventoryClient(conn)
}

func TestSetInv(goodsId, Num int32) {
	_, err := invClient.SetInv(context.Background(), &proto.GoodsInvInfo{
		GoodsId: goodsId,
		Num:     Num,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("设置库存成功")
}

func TestInvDetail(goodsId int32) {
	rsp, err := invClient.InvDetail(context.Background(), &proto.GoodsInvInfo{
		GoodsId: goodsId,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Num)
}

func TestSell() {
	/*
		1. 第一件扣减成功： 第二件： 1. 没有库存信息 2. 库存不足
		2. 两件都扣减成功
	*/
	_, err := invClient.Sell(context.Background(), &proto.SellInfo{
		GoodsInfo: []*proto.GoodsInvInfo{
			{GoodsId: 1, Num: 1},
			{GoodsId: 2, Num: 70},
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("库存扣减成功")
}

func TestReback() {
	_, err := invClient.Reback(context.Background(), &proto.SellInfo{
		GoodsInfo: []*proto.GoodsInvInfo{
			{GoodsId: 1, Num: 10},
			{GoodsId: 100, Num: 30},
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("归还成功")
}

func main() {
	Init()
	//var i int32
	//for i = 1; i <= 9; i++ {
	//	TestSetInv(i, 90)
	//}

	//TestInvDetail(2)
	//TestSell()
	TestReback()
	conn.Close()
}
