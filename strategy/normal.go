package strategy

import "math"

type NormalDeliver struct {
}

const (
	BaseFeeNormal float64 = 10
)

func (n *NormalDeliver) CalcFee(orderAmount, weight float64) float64 {
	extra := math.Max(0, weight-1)
	return BaseFeeNormal + math.Ceil(extra)*2
}
