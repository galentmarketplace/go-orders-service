package cart

import (
	"reflect"
	"testing"
)

func TestCount(t *testing.T) {
	cases := []struct {
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
				{SKU: "a", Qty: 3, Price: 1.0},
			},
			want: 3,
		},
		{
			name: "multiple items",
			items: []Item{
				{SKU: "a", Qty: 3, Price: 1.0},
				{SKU: "b", Qty: 5, Price: 2.0},
				{SKU: "c", Qty: 0, Price: 3.0},
			},
			want: 8,
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
			name: "sku not found",
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

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := Remove(tc.items, tc.sku)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Remove(%v, %q) = %v, want %v", tc.items, tc.sku, got, tc.want)
			}
		})
	}
}
