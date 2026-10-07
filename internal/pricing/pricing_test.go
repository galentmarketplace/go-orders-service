package pricing

import "testing"

func TestSubtotal(t *testing.T) {
	got, err := Subtotal([]float64{10, 20, 5})
	if err != nil || got != 35 {
		t.Fatalf("Subtotal = %v, %v; want 35, nil", got, err)
	}
	if _, err := Subtotal([]float64{-1}); err == nil {
		t.Fatal("Subtotal(-1) should return an error")
	}
}

func TestApplyTax(t *testing.T) {
	got, err := ApplyTax(100, 0.2)
	if err != nil || got != 120 {
		t.Fatalf("ApplyTax = %v, %v; want 120, nil", got, err)
	}
}
