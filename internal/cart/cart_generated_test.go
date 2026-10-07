package cart

import "testing"

func TestCount(t *testing.T) {
	cases := []struct {
		name  string
		items []Item
		want  int
	}{
		{
			name:  "empty",
			items: []Item{},
			want:  0,
		},
		{
			name: "multiple items",
			items: []Item{
				{SKU: "A", Qty: 2, Price: 10},
				{SKU: "B", Qty: 3, Price: 5},
			},
			want: 5,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := Count(tc.items)
			if got != tc.want {
				t.Errorf("Count(%v) = %d, want %d", tc.items, got, tc.want)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	cases := []struct {
		name  string
		items []Item
		sku   string
		want  []Item
	}{
		{
			name: "remove existing",
			items: []Item{
				{SKU: "A", Qty: 1, Price: 1},
				{SKU: "B", Qty: 2, Price: 2},
				{SKU: "C", Qty: 3, Price: 3},
			},
			sku: "B",
			want: []Item{
				{SKU: "A", Qty: 1, Price: 1},
				{SKU: "C", Qty: 3, Price: 3},
			},
		},
		{
			name: "sku not found",
			items: []Item{
				{SKU: "A", Qty: 1, Price: 1},
			},
			sku: "Z",
			want: []Item{
				{SKU: "A", Qty: 1, Price: 1},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := Remove(tc.items, tc.sku)
			if len(got) != len(tc.want) {
				t.Fatalf("Remove() length = %d, want %d", len(got), len(tc.want))
			}
			for i := range got {
				if got[i] != tc.want[i] {
					t.Errorf("Remove() item[%d] = %+v, want %+v", i, got[i], tc.want[i])
				}
			}
		})
	}
}
