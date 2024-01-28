package calculator_test

import (
	"calculator"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()

	got := calculator.Add(2, 2)

	//it's needed to be initialized explicitly otherwise "want := 4" is int
	var want float64 = 4

	if got != want {
		t.Errorf("got %f want %f", got, want)
	}
}
