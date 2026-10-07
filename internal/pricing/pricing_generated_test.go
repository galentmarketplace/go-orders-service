package pricing

import "testing"

func TestBulkDiscount(t *testing.T) {
	cases := []struct {
		name string
		qty  int
		want float64
	}{
		{"below threshold", 5, 0},
		{"zero qty", 0, 0},
		{"exactly 10", 10, 0.05},
		{"between 10 and 50", 20, 0.05},
		{"exactly 50", 50, 0.10},
		{"between 50 and 100", 75, 0.10},
		{"exactly 100", 100, 0.20},
		{"above 100", 500, 0.20},
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
		{"not a member", 100, false, 0},
		{"zero total", 0, true, 0},
		{"negative total", -50, true, 0},
		{"member below 500", 100, true, 100},
		{"member exactly 500", 500, true, 1000},
		{"member above 500", 600, true, 1200},
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
