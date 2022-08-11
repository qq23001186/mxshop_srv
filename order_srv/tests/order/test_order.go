package main

import (
	"context"
	"fmt"
	"nd/order_srv/proto"
	"nd/order_srv/tests"
)

func TestCreateOrder() {
	_, err := tests.OrderClient.CreateOrder(context.Background(), &proto.OrderRequest{
		UserId:  1,
		Address: "北京市",
		Name:    "bobby",
		Mobile:  "18787878787",
		Post:    "请尽快发货",
	})
	if err != nil {
		panic(err)
	}
}

func TestGetOrderDetail(orderId int32) {
	rsp, err := tests.OrderClient.OrderDetail(context.Background(), &proto.OrderRequest{
		Id: orderId,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.OrderInfo.OrderSn)
	for _, good := range rsp.Goods {
		fmt.Println(good.GoodsName)
	}
}

func TestOrderList() {
	//rsp, err := tests.OrderClient.OrderList(context.Background(), &proto.OrderFilterRequest{})
	rsp, err := tests.OrderClient.OrderList(context.Background(), &proto.OrderFilterRequest{
		UserId: 1,
	})
	if err != nil {
		panic(err)
	}
	for _, order := range rsp.Data {
		fmt.Println(order.OrderSn)
	}
}

func main() {
	tests.Init()

	//TestCreateOrder()
	//TestGetOrderDetail(1)
	TestOrderList()
	tests.Conn.Close()
}
