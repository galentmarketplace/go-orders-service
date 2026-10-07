package cart

import "testing"

func TestCount(t *testing.T) {
	tests := []struct {
		name  string
		items []Item
		want  int
	}{
		{
			name:  "empty slice",
			items: []Item{},
			want:  0,
		},
		{
			name: "single item",
			items: []Item{
				{SKU: "A1", Qty: 3, Price: 10.0},
			},
			want: 3,
		},
		{
			name: "multiple items",
			items: []Item{
				{SKU: "A1", Qty: 3, Price: 10.0},
				{SKU: "B2", Qty: 5, Price: 20.0},
				{SKU: "C3", Qty: 2, Price: 5.0},
			},
			want: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Count(tt.items)
			if got != tt.want {
				t.Errorf("Count(%v) = %d, want %d", tt.items, got, tt.want)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	tests := []struct {
		name    string
		items   []Item
		sku     string
		wantLen int
		wantSKUs []string
	}{
		{
			name: "remove existing item",
			items: []Item{
				{SKU: "A1", Qty: 1, Price: 10.0},
				{SKU: "B2", Qty: 2, Price: 20.0},
				{SKU: "C3", Qty: 3, Price: 30.0},
			},
			sku:     "B2",
			wantLen: 2,
			wantSKUs: []string{"A1", "C3"},
		},
		{
			name: "remove non-existing item",
			items: []Item{
				{SKU: "A1", Qty: 1, Price: 10.0},
				{SKU: "B2", Qty: 2, Price: 20.0},
			},
			sku:     "ZZZ",
			wantLen: 2,
			wantSKUs: []string{"A1", "B2"},
		},
		{
			name:    "remove from empty slice",
			items:   []Item{},
			sku:     "A1",
			wantLen: 0,
			wantSKUs: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Remove(tt.items, tt.sku)
			if len(got) != tt.wantLen {
				t.Fatalf("Remove() returned length %d, want %d", len(got), tt.wantLen)
			}
			for i, sku := range tt.wantSKUs {
				if got[i].SKU != sku {
					t.Errorf("Remove()[%d].SKU = %s, want %s", i, got[i].SKU, sku)
				}
			}
		})
	}
}
