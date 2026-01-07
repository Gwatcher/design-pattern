package strategy

import "testing"

func TestDeliveryStrategies(t *testing.T) {
	tests := []struct {
		name     string
		order    Order
		expected float64
	}{
		{
			name:     "local fixed fee",
			order:    Order{OrderAmount: 10, Weight: 5, DeliverType: LocalDeliverType},
			expected: 8,
		},
		{
			name:     "normal <=1kg",
			order:    Order{OrderAmount: 10, Weight: 1, DeliverType: NormalDeliverType},
			expected: 10,
		},
		{
			name:     "normal 1.2kg",
			order:    Order{OrderAmount: 10, Weight: 1.2, DeliverType: NormalDeliverType},
			expected: 12, // 基础10 + ceil(0.2)*2
		},
		{
			name:     "sf negative weight treated as 0",
			order:    Order{OrderAmount: 10, Weight: -1, DeliverType: SFDeliverType},
			expected: 15, // 取决于你SF策略的规则
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, ok := GetDeliverStrategy(tt.order.DeliverType)
			if !ok {
				t.Fatalf("unknown deliver type: %s", tt.order.DeliverType)
			}
			got := s.CalcFee(tt.order.OrderAmount, tt.order.Weight)
			if got != tt.expected {
				t.Fatalf("expected %v, got %v", tt.expected, got)
			}
		})
	}
}
