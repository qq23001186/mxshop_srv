package main

import (
	"context"
	"fmt"
	"nd/goods_srv/tests"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"nd/goods_srv/proto"
)

var brandClient proto.GoodsClient
var conn *grpc.ClientConn

func TestGetBrandList() {
	rsp, err := brandClient.BrandList(context.Background(), &proto.BrandFilterRequest{})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Total)
	for _, brand := range rsp.Data {
		fmt.Println(brand.Name)
	}
}

func Init() {
	var err error
	conn, err = grpc.Dial(tests.TargetAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	brandClient = proto.NewGoodsClient(conn)
}

func main() {
	Init()
	TestGetBrandList()

	conn.Close()
}
