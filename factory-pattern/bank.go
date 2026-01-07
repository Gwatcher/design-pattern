package factory_pattern

type Bank struct {
}

func (b *Bank) CalcFee(amount float64) float64 {
	if amount < 1000 {
		return FormatDigit(3)
	}
	return FormatDigit(2)
}
