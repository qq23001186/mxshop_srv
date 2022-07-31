package handler

import (
	"nd/goods_srv/proto"
)

type GoodsServer struct {
	proto.UnimplementedGoodsServer
}

//// GoodsList 商品接口
//func (s *GoodsServer) GoodsList(ctx context.Context, req *proto.GoodsFilterRequest) (*proto.GoodsListResponse, error) {
//	return nil, nil
//}
//
//// BatchGetGoods 现在用户提交订单有多个商品，你得批量查询商品的信息吧
//func (s *GoodsServer) BatchGetGoods(ctx context.Context, req *proto.BatchGoodsIdInfo) (*proto.GoodsListResponse, error) {
//	return nil, nil
//}
//func (s *GoodsServer) CreateGoods(ctx context.Context, req *proto.CreateGoodsInfo) (*proto.GoodsInfoResponse, error) {
//	return nil, nil
//}
//func (s *GoodsServer) DeleteGoods(ctx context.Context, req *proto.DeleteGoodsInfo) (*emptypb.Empty, error) {
//	return nil, nil
//}
//func (s *GoodsServer) UpdateGoods(ctx context.Context, req *proto.CreateGoodsInfo) (*emptypb.Empty, error) {
//	return nil, nil
//}
//func (s *GoodsServer) GetGoodsDetail(ctx context.Context, req *proto.GoodInfoRequest) (*proto.GoodsInfoResponse, error) {
//	return nil, nil
//}
