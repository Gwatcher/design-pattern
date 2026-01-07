package strategy

import "math"

type SFDeliver struct {
}

const (
	BaseFeeSF float64 = 15
)

func (s *SFDeliver) CalcFee(orderAmount, weight float64) float64 {
	extra := math.Max(0, weight-1)
	return BaseFeeSF + math.Ceil(extra)*5
}
