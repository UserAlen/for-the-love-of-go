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

func TestSubtract(t *testing.T) {
	t.Parallel()

	var a, b float64

	a = 4
	b = 2

	got := calculator.Subtract(a, b)
	var want float64 = 2

	if got != want {
		t.Errorf("Subtract(%.1f, %.1f): got %.1f want %.1f", a, b, got, want)
	}

}

func TestMultiply(t *testing.T) {

	t.Parallel()

	type testCases struct {
		a, b, want float64
	}

	tcs := []testCases{
		{a: 5, b: 75, want: 375},
		{a: 6, b: 348, want: 2088},
		{a: 7, b: 89, want: 623},
		{a: 12, b: 574, want: 6888},
		{a: 11, b: 53549, want: 589039},
	}

	for _, tc := range tcs {
		got := calculator.Multiply(tc.a, tc.b)

		if got != tc.want {
			t.Errorf("Multiply(%.1f, %.1f): got %.1f want %.1f", tc.a, tc.b, got, tc.want)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()

	type testCase struct {
		a, b, want float64
	}

	tcs := []testCase{
		{a: 5, b: 2, want: 2.5},
		{a: 6, b: 3, want: 2},
		{a: 75, b: 5, want: 15}, //always put trailing comma
		{a: 8, b: 2, want: 8 / 2},
	}

	for _, tc := range tcs {
		got := calculator.Divide(tc.a, tc.b)

		if got != tc.want {
			t.Errorf("Divide(%f, %f): want %f got %f", tc.a, tc.b, got, tc.want)
		}
	}

}
