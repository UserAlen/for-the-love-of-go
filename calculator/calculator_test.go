package calculator_test

import (
	"calculator"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()

	var a, b float64

	a, b = 2, 2

	got := calculator.Add(a, b)

	//it's needed to be initialized explicitly otherwise "want := 4" is int
	var want float64 = 4

	if got != want {
		t.Errorf("Add(%.1f, %.1f): got %.1f want %.1f", a, b, got, want)
	}
}

func TestSubstract(t *testing.T) {
	t.Parallel()

	var a, b float64

	a = 4
	b = 2

	got := calculator.Substract(a, b)
	var want float64 = 2

	if got != want {
		t.Errorf("Substract(%.1f, %.1f): got %.1f want %.1f", a, b, got, want)
	}

}
