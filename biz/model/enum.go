package model

type SendEmailType uint // 发送邮件类型

const (
	BindEmail   SendEmailType = 0
	UnBindEmail SendEmailType = 1
	ChangePwd   SendEmailType = 2
)

type FreightType uint // 取货方式

const (
	FreeFreight      FreightType = 0 // "包邮"
	PickUp           FreightType = 1 // "自提"
	NeedShippingCost FreightType = 2 // "按距离估算运费"
)

type OrderType uint // 订单状态

const (
	ToBePaid          OrderType = 0 // 下单待付款
	ToBeShipped       OrderType = 1 // 付款待发货
	ToBeReceived      OrderType = 2 // 发货待收货
	ToBeEvaluated     OrderType = 3 // 收货待评价
	ToBeClosed        OrderType = 4 // 评价待关闭
	TransactionClosed OrderType = 5 // 已经关闭

	RefundInProgress OrderType = 6 // 退款中
	CompletedRefund  OrderType = 7 // 退款完成
	RefundClosed     OrderType = 8 // 退款关闭
)
