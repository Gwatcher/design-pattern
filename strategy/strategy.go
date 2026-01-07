package strategy

type DeliverStrategy interface {
	CalcFee(orderAmount, weight float64) float64
}

var DeliverStrategyMap = map[DeliverType]DeliverStrategy{
	SFDeliverType:     &SFDeliver{},
	LocalDeliverType:  &LocalDeliver{},
	NormalDeliverType: &NormalDeliver{},
}

func GetDeliverStrategy(deliverType DeliverType) (DeliverStrategy, bool) {
	s, ok := DeliverStrategyMap[deliverType]
	return s, ok
}
