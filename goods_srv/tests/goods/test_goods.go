package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"nd/goods_srv/proto"
	"nd/goods_srv/tests"
)

var brandClient proto.GoodsClient
var conn *grpc.ClientConn

func Init() {
	var err error
	conn, err = grpc.Dial(tests.TargetAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	brandClient = proto.NewGoodsClient(conn)
}

func TestGetGoodsList() {
	rsp, err := brandClient.GoodsList(context.Background(), &proto.GoodsFilterRequest{
		//TopCategory: 2007,
		PriceMin: 10,
		KeyWords: "水",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Total)
	for _, good := range rsp.Data {
		fmt.Println(good.Name, good.ShopPrice)
	}
}

func TestBatchGetGoods() {
	rsp, err := brandClient.BatchGetGoods(context.Background(), &proto.BatchGoodsIdInfo{
		Id: []int32{29, 30, 31},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Total)
	for _, good := range rsp.Data {
		fmt.Println(good.Name, good.ShopPrice)
	}
}

func TestGetGoodsDetail() {
	rsp, err := brandClient.GetGoodsDetail(context.Background(), &proto.GoodInfoRequest{
		Id: 28,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Name)
	fmt.Println(rsp.ShopPrice)
	fmt.Println(rsp.MarketPrice)
}

func main() {
	Init()
	//TestGetGoodsList()
	TestBatchGetGoods()
	fmt.Println("----------------------------")
	TestGetGoodsDetail()
	conn.Close()
}
