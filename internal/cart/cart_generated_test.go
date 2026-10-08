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
			name:  "empty slice",
			items: []Item{},
			want:  0,
		},
		{
			name: "single item",
			items: []Item{
				{SKU: "a", Qty: 3, Price: 1.5},
			},
			want: 3,
		},
		{
			name: "multiple items",
			items: []Item{
				{SKU: "a", Qty: 3, Price: 1.5},
				{SKU: "b", Qty: 5, Price: 2.0},
				{SKU: "c", Qty: 0, Price: 0},
			},
			want: 8,
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
			name: "remove existing item",
			items: []Item{
				{SKU: "a", Qty: 1, Price: 1.0},
				{SKU: "b", Qty: 2, Price: 2.0},
				{SKU: "c", Qty: 3, Price: 3.0},
			},
			sku: "b",
			want: []Item{
				{SKU: "a", Qty: 1, Price: 1.0},
				{SKU: "c", Qty: 3, Price: 3.0},
			},
		},
		{
			name: "remove non-existing item",
			items: []Item{
				{SKU: "a", Qty: 1, Price: 1.0},
				{SKU: "b", Qty: 2, Price: 2.0},
			},
			sku: "z",
			want: []Item{
				{SKU: "a", Qty: 1, Price: 1.0},
				{SKU: "b", Qty: 2, Price: 2.0},
			},
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
				t.Errorf("Remove() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
