package model

type Inventory struct {
	BaseModel
	Goods   int32 `gorm:"type:int;index"` // 商品id
	Stocks  int32 `gorm:"type:int"`       // 库存
	Version int32 `gorm:"type:int"`       //分布式锁的乐观锁
}

//type InventoryHistory struct {
//	user int32
//	goods int32
//	nums int32
//	order int32
//	status int32 //1. 表示库存是预扣减， 幂等性， 2. 表示已经支付
//}
