package strategy

type LocalDeliver struct {
}

const (
	BaseFeeLocal float64 = 8
)

func (l *LocalDeliver) CalcFee(orderAmount, weight float64) float64 {
	return BaseFeeLocal
}
