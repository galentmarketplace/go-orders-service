package pricing

import "testing"

func TestBulkDiscount(t *testing.T) {
	tests := []struct {
		name string
		qty  int
		want float64
	}{
		{name: "qty 0 - no discount", qty: 0, want: 0},
		{name: "qty 9 - no discount", qty: 9, want: 0},
		{name: "qty 10 - 5% discount", qty: 10, want: 0.05},
		{name: "qty 49 - 5% discount", qty: 49, want: 0.05},
		{name: "qty 50 - 10% discount", qty: 50, want: 0.10},
		{name: "qty 99 - 10% discount", qty: 99, want: 0.10},
		{name: "qty 100 - 20% discount", qty: 100, want: 0.20},
		{name: "qty 1000 - 20% discount", qty: 1000, want: 0.20},
		{name: "negative qty - no discount", qty: -5, want: 0},
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
		{name: "non-member", total: 100, member: false, want: 0},
		{name: "member with zero total", total: 0, member: true, want: 0},
		{name: "member with negative total", total: -50, member: true, want: 0},
		{name: "member below 500", total: 100, member: true, want: 100},
		{name: "member at 500 - doubled", total: 500, member: true, want: 1000},
		{name: "member above 500 - doubled", total: 750, member: true, want: 1500},
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
