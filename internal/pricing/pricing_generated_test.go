package pricing

import "testing"

func TestBulkDiscount(t *testing.T) {
	cases := []struct {
		name string
		qty  int
		want float64
	}{
		{name: "below threshold", qty: 5, want: 0},
		{name: "ten or more", qty: 10, want: 0.05},
		{name: "fifty or more", qty: 50, want: 0.10},
		{name: "hundred or more", qty: 100, want: 0.20},
		{name: "zero qty", qty: 0, want: 0},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := BulkDiscount(tc.qty)
			if got != tc.want {
				t.Errorf("BulkDiscount(%d) = %v, want %v", tc.qty, got, tc.want)
			}
		})
	}
}

func TestLoyaltyPoints(t *testing.T) {
	cases := []struct {
		name   string
		total  float64
		member bool
		want   int
	}{
		{name: "not a member", total: 100, member: false, want: 0},
		{name: "member but zero total", total: 0, member: true, want: 0},
		{name: "member negative total", total: -10, member: true, want: 0},
		{name: "member below 500", total: 100, member: true, want: 100},
		{name: "member at 500 doubles", total: 500, member: true, want: 1000},
		{name: "member above 500 doubles", total: 600, member: true, want: 1200},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := LoyaltyPoints(tc.total, tc.member)
			if got != tc.want {
				t.Errorf("LoyaltyPoints(%v, %v) = %d, want %d", tc.total, tc.member, got, tc.want)
			}
		})
	}
}
