package main

import (
	"context"
	"fmt"
	"nd/order_srv/proto"
	"nd/order_srv/tests"
)

func TestCartItemList(userId int32) {
	rsp, err := tests.OrderClient.CartItemList(context.Background(), &proto.UserInfo{
		Id: userId,
	})
	if err != nil {
		panic(err)
	}
	for _, item := range rsp.Data {
		fmt.Println(item.Id, item.GoodsId, item.Nums)
	}
}

func TestCreateCartItem(userId, nums, goodsId int32) {
	rsp, err := tests.OrderClient.CreateCartItem(context.Background(), &proto.CartItemRequest{
		UserId:  userId,
		Nums:    nums,
		GoodsId: goodsId,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Id)
}

func TestUpdateCartItem(id int32) {
	_, err := tests.OrderClient.UpdateCartItem(context.Background(), &proto.CartItemRequest{
		Id:      id,
		Checked: true,
	})
	if err != nil {
		panic(err)
	}
}

func main() {
	tests.Init()
	TestCreateCartItem(1, 1, 26)
	TestCreateCartItem(1, 1, 27)
	TestCartItemList(1)
	TestUpdateCartItem(1)
	tests.Conn.Close()
}
