package model

type CreateOrderRequest struct {
	ProductId uint      `json:"product_id"` // 商品
	AddressId uint      `json:"address_id"` // 地址
	BossId    uint      `json:"boss_id"`    // 卖家
	Money     float64   `json:"money"`      // 金额
	Type      OrderType `json:"type"`       // 订单状态

}

type CreateOrderResponse struct {
}
