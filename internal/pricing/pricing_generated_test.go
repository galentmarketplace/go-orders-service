package pricing

import "testing"

func TestBulkDiscount(t *testing.T) {
	tests := []struct {
		name string
		qty  int
		want float64
	}{
		{name: "qty 100 gives 0.20", qty: 100, want: 0.20},
		{name: "qty above 100 gives 0.20", qty: 150, want: 0.20},
		{name: "qty 50 gives 0.10", qty: 50, want: 0.10},
		{name: "qty 10 gives 0.05", qty: 10, want: 0.05},
		{name: "qty below 10 gives 0", qty: 5, want: 0},
		{name: "qty 0 gives 0", qty: 0, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BulkDiscount(tt.qty)
			if got != tt.want {
				t.Errorf("BulkDiscount(%d) = %v, want %v", tt.qty, got, tt.want)
			}
		})
	}
}

func TestLoyaltyPoints(t *testing.T) {
	tests := []struct {
		name   string
		total  float64
		member bool
		want   int
	}{
		{name: "non member returns 0", total: 600, member: false, want: 0},
		{name: "member with zero total returns 0", total: 0, member: true, want: 0},
		{name: "member with negative total returns 0", total: -10, member: true, want: 0},
		{name: "member below 500 returns points", total: 100, member: true, want: 100},
		{name: "member at 500 doubles points", total: 500, member: true, want: 1000},
		{name: "member above 500 doubles points", total: 600, member: true, want: 1200},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LoyaltyPoints(tt.total, tt.member)
			if got != tt.want {
				t.Errorf("LoyaltyPoints(%v, %v) = %d, want %d", tt.total, tt.member, got, tt.want)
			}
		})
	}
}
