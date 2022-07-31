package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"nd/goods_srv/tests"

	"nd/goods_srv/proto"
)

var brandClient proto.GoodsClient
var conn *grpc.ClientConn

func TestGetCategoryList() {
	rsp, err := brandClient.GetAllCategorysList(context.Background(), &empty.Empty{})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Total)
	fmt.Println(rsp.JsonData)
}

func TestGetSubCategoryList() {
	rsp, err := brandClient.GetSubCategory(context.Background(), &proto.CategoryListRequest{
		Id: 135487,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.SubCategorys)
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
	TestGetSubCategoryList()
	TestGetCategoryList()

	conn.Close()
}
