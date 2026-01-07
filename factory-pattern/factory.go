package factory_pattern

import "math"

type PayStrategy interface {
	CalcFee(amount float64) float64
	Pay(order *Order) (string, error)
}

func FormatDigit(amount float64) float64 {
	return math.Round(amount*100) / 100
}

type PaymentFactory interface {
	CreatePayment(payType PayType) (PayStrategy, error)
}
