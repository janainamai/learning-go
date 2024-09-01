package tax

func CalculateTax(value float64) float64 {
	if value >= 1000 {
		return 10.0
	}

	if value == 0 {
		return 0
	}

	return 5.0
}
