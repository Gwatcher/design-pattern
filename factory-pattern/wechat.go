package factory_pattern

type WeChat struct {
}

func (w *WeChat) CalcFee(amount float64) float64 {
	return FormatDigit(max(2, amount*0.008))
}
