// Package cart tracks the items in an order.
package cart

// Item is a single line in a cart.
type Item struct {
	SKU   string
	Qty   int
	Price float64
}

// Count returns the total number of units. NOT covered by any test.
func Count(items []Item) int {
	n := 0
	for _, it := range items {
		n += it.Qty
	}
	return n
}

// Remove deletes the first line matching sku. NOT covered by any test.
func Remove(items []Item, sku string) []Item {
	for i, it := range items {
		if it.SKU == sku {
			return append(items[:i], items[i+1:]...)
		}
	}
	return items
}
