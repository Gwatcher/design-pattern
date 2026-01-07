package strategy

type Order struct {
	OrderAmount float64     // 订单金额
	Weight      float64     // 订单重量
	DeliverType DeliverType // 配送方式
}

type DeliverType string

const (
	NormalDeliverType DeliverType = "normal" // 普通
	SFDeliverType     DeliverType = "sf"     // 顺丰
	LocalDeliverType  DeliverType = "local"  // 同城
)
