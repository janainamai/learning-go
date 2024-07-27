package tax

import "errors"

func CalculateTax(value float64) (float64, error) {
	if value <= 0 {
		return 0, errors.New("amount must be greater than 0")
	}

	if value >= 10000 && value < 20000 {
		return 10.0, nil
	}

	if value >= 20000 {
		return 20.0, nil
	}

	return 5.0, nil
}
