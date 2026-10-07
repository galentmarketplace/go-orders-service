package pricing

import "testing"

func TestBulkDiscount(t *testing.T) {
	cases := []struct {
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
		{name: "non-member gets zero points", total: 100, member: false, want: 0},
		{name: "member with zero total gets zero points", total: 0, member: true, want: 0},
		{name: "member with negative total gets zero points", total: -50, member: true, want: 0},
		{name: "member under 500 gets base points", total: 100, member: true, want: 100},
		{name: "member exactly 500 gets double points", total: 500, member: true, want: 1000},
		{name: "member above 500 gets double points", total: 600, member: true, want: 1200},
		{name: "member with fractional total truncates", total: 100.9, member: true, want: 100},
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
