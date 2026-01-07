package factory_pattern

import "errors"

type AliPay struct {
}

func (p *AliPay) CalcFee(amount float64) float64 {
	return FormatDigit(max(1, amount*0.006))
}

func (p *AliPay) Pay(order *Order) (string, error) {
	if order.Amount <= 0 {
		return "", errors.New("amount is wrong")
	}
	return "success", nil
}
