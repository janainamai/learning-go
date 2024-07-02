package tax

import (
	"testing"
)

func TestCalculateTax(t *testing.T) {
	amount := 900.0
	expectedTax := 5.0

	result := CalculateTax(amount)

	if result != expectedTax {
		t.Errorf("Expected %f but got %f", expectedTax, result)
	}
}
