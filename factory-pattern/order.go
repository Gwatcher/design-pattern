package factory_pattern

import "fmt"

type PayType string
type Order struct {
	OrderID string
	Amount  float64
	PayType PayType
}

func (o Order) String() string {
	return fmt.Sprintf("OrderID: %v,\nAmount: %v,\nPayType:%v", o.OrderID, o.Amount, o.PayType)
}

const (
	PayTypeAliPay    PayType = "alipay"
	PayTypeWeChatPay PayType = "wechat"
	PayTypeBankPay   PayType = "bank"
)
