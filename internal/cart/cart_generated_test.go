package cart

import (
	"reflect"
	"testing"
)

func TestCount(t *testing.T) {
	tests := []struct {
		name  string
		items []Item
		want  int
	}{
		{
			name:  "happy path multiple items",
			items: []Item{{SKU: "a", Qty: 2}, {SKU: "b", Qty: 3}},
			want:  5,
		},
		{
			name:  "empty slice",
			items: []Item{},
			want:  0,
		},
		{
			name:  "nil slice",
			items: nil,
			want:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Count(tt.items)
			if got != tt.want {
				t.Errorf("Count() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	tests := []struct {
		name  string
		items []Item
		sku   string
		want  []Item
	}{
		{
			name:  "remove existing sku",
			items: []Item{{SKU: "a", Qty: 1}, {SKU: "b", Qty: 2}, {SKU: "c", Qty: 3}},
			sku:   "b",
			want:  []Item{{SKU: "a", Qty: 1}, {SKU: "c", Qty: 3}},
		},
		{
			name:  "sku not found",
			items: []Item{{SKU: "a", Qty: 1}, {SKU: "b", Qty: 2}},
			sku:   "z",
			want:  []Item{{SKU: "a", Qty: 1}, {SKU: "b", Qty: 2}},
		},
		{
			name:  "empty slice",
			items: []Item{},
			sku:   "a",
			want:  []Item{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Remove(tt.items, tt.sku)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Remove() = %v, want %v", got, tt.want)
			}
		})
	}
}
