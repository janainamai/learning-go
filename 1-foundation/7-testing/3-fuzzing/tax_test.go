package tax

import "testing"

func FuzzCalculateTax(f *testing.F) {
	// dando exemplos do que quero testar
	seed := []float64{-1, -2, -2.5, 500.0, 1000.0, 1500.0}

	for _, amount := range seed {
		f.Add(amount)
	}

	f.Fuzz(func(t *testing.T, amount float64) {
		result := CalculateTax(amount)

		// validando comportamento esperado
		if amount <= 0 && result != 0 {
			t.Errorf("Received %f but expected 0", result)
		}

		if amount >= 10000 && amount < 20000 && result != 10.0 {
			t.Errorf("Received %f but expected 10", result)
		}

		if amount >= 20000 && result != 20.0 {
			t.Errorf("Received %f but expected 20", result)
		}

		if amount < 10000 && amount > 0 && result != 5 {
			t.Errorf("Received %f but expected 5", result)
		}
	})

}

// go test -fuzz=. -run=^#

// determinando tempo a ser executado
// go test -fuzz=. -fuzztime=5s -run=^#
