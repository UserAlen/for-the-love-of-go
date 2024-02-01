package calculator

import (
	"errors"
	"math"
)

func Add(a, b float64) float64 {
	return a + b
}

func Subtract(a, b float64) float64 {
	return a - b
}

func Multiply(a, b float64) float64 {
	return a * b
}

func Divide(a, b float64) (float64, error) {

	if b == 0 {
		return 0, errors.New("division by zero is not allowed")
	}

	return a / b, nil
}

func Sqrt(a float64) (float64, error) {

	if a < 0 {
		return 0, errors.New("values below 0 are not allowed by function Sqrt")
	}

	return math.Sqrt(a), nil
}
