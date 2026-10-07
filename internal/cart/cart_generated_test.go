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
			name:  "empty",
			items: []Item{},
			want:  0,
		},
		{
			name: "multiple items",
			items: []Item{
				{SKU: "a", Qty: 2, Price: 10},
				{SKU: "b", Qty: 3, Price: 5},
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
				{SKU: "a", Qty: 1, Price: 1},
				{SKU: "b", Qty: 2, Price: 2},
				{SKU: "c", Qty: 3, Price: 3},
			},
			sku: "b",
			want: []Item{
				{SKU: "a", Qty: 1, Price: 1},
				{SKU: "c", Qty: 3, Price: 3},
			},
		},
		{
			name: "sku not found",
			items: []Item{
				{SKU: "a", Qty: 1, Price: 1},
			},
			sku: "zzz",
			want: []Item{
				{SKU: "a", Qty: 1, Price: 1},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := Remove(tc.items, tc.sku)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Remove() = %v, want %v", got, tc.want)
			}
		})
	}
}
