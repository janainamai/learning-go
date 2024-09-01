package tax

func CalculateTax(value float64) float64 {
	if value <= 0 {
		return 0
	}

	if value >= 10000 && value < 20000 {
		return 10.0
	}

	if value >= 20000 {
		return 20.0
	}

	return 5.0
}
