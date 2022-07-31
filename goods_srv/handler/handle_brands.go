package handler

import (
	"context"
	"nd/goods_srv/global"
	"nd/goods_srv/model"
	"nd/goods_srv/proto"
)

// BrandList 品牌和轮播图
func (s *GoodsServer) BrandList(ctx context.Context, req *proto.BrandFilterRequest) (*proto.BrandListResponse, error) {
	brandListResponse := proto.BrandListResponse{}

	var brands []model.Brands
	result := global.DB.Scopes(Paginate(int(req.Pages), int(req.PagePerNums))).Find(&brands)
	if result.Error != nil {
		return nil, result.Error
	}

	var total int64
	global.DB.Model(&model.Brands{}).Count(&total)
	brandListResponse.Total = int32(total)

	var brandResponses []*proto.BrandInfoResponse
	for _, brand := range brands {
		brandResponses = append(brandResponses, &proto.BrandInfoResponse{
			Id:   brand.ID,
			Name: brand.Name,
			Logo: brand.Logo,
		})
	}
	brandListResponse.Data = brandResponses
	return &brandListResponse, nil
}

//CreateBrand(ctx context.Context, in *BrandRequest, opts ...grpc.CallOption) (*BrandInfoResponse, error)
//DeleteBrand(ctx context.Context, in *BrandRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
//UpdateBrand(ctx context.Context, in *BrandRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
