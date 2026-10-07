// Package pricing computes order totals, discounts and tax.
package pricing

import "errors"

// ErrNegative is returned when a caller passes a negative amount.
var ErrNegative = errors.New("pricing: amount must not be negative")

// Subtotal sums line amounts. Covered by the existing tests.
func Subtotal(amounts []float64) (float64, error) {
	var total float64
	for _, a := range amounts {
		if a < 0 {
			return 0, ErrNegative
		}
		total += a
	}
	return total, nil
}

// ApplyTax adds tax at the given rate. Covered by the existing tests.
func ApplyTax(amount, rate float64) (float64, error) {
	if amount < 0 || rate < 0 {
		return 0, ErrNegative
	}
	return amount + amount*rate, nil
}

// BulkDiscount returns the discount rate for a quantity.
// NOT covered by any test — three branches nobody verifies.
func BulkDiscount(qty int) float64 {
	switch {
	case qty >= 100:
		return 0.20
	case qty >= 50:
		return 0.10
	case qty >= 10:
		return 0.05
	default:
		return 0
	}
}

// LoyaltyPoints awards points per whole currency unit, doubled for large orders.
// NOT covered by any test.
func LoyaltyPoints(total float64, member bool) int {
	if !member || total <= 0 {
		return 0
	}
	pts := int(total)
	if total >= 500 {
		pts *= 2
	}
	return pts
}
