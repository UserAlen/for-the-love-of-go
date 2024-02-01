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

	type testCases struct {
		fn         func(float64, float64) float64
		a, b, want float64
	}

	tcs := []testCases{
		{fn: calculator.Subtract, a: 5, b: 2, want: 3},
		{fn: calculator.Subtract, a: 55, b: 5, want: 50},
	}

	for _, tc := range tcs {
		got := tc.fn(tc.a, tc.b)

		if got != tc.want {
			t.Errorf("Subtract(%.1f, %.1f): got %.1f want %.1f", tc.a, tc.b, got, tc.want)
		}
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

	tcs := []struct {
		a, b        float64
		want        float64
		errExpected bool
	}{
		{a: 10, b: 5, want: 2, errExpected: false},
		{a: 2, b: 0, want: 9999, errExpected: true},
	}

	for _, tc := range tcs {
		got, err := calculator.Divide(tc.a, tc.b)

		errReceived := (err != nil)

		if errReceived != tc.errExpected {
			t.Fatalf("Divide(%.1f, %.1f): unexpected error value: %v", tc.a, tc.b, err)
		}

		if !errReceived && got != tc.want {
			t.Errorf("Divide(%.1f, %.1f): want %.1f got %.1f", tc.a, tc.b, tc.want, got)
		}
	}

}

func TestAddSubMul(t *testing.T) {
	t.Parallel()

	type testCase struct {
		fnName     string
		fn         func(float64, float64) float64
		a, b, want float64
	}

	tcs := []testCase{
		{fnName: "Add", fn: calculator.Add, a: 5, b: 5, want: 10},
		{fnName: "Subtract", fn: calculator.Subtract, a: 7, b: 2, want: 5},
		{fnName: "Multiply", fn: calculator.Multiply, a: 55, b: 5, want: 275},
	}

	for _, tc := range tcs {
		got := tc.fn(tc.a, tc.b)

		if got != tc.want {
			t.Errorf("%q(%f, %f): got %f, want %f", tc.fnName, tc.a, tc.b, got, tc.want)
		}
	}
}

func TestSqrt(t *testing.T) {
	t.Parallel()

	type testCase struct {
		fn          func(float64) (float64, error)
		fnName      string
		a           float64
		want        float64
		errExpected bool
	}

	tcs := []testCase{
		{fn: calculator.Sqrt, fnName: "Sqrt", a: 4, want: 2, errExpected: false},
		{fn: calculator.Sqrt, fnName: "Sqrt", a: 0, want: 0, errExpected: false},
		{fn: calculator.Sqrt, fnName: "Sqrt", a: -1, want: 9999, errExpected: true},
	}

	for _, tc := range tcs {
		got, err := tc.fn(tc.a)

		errReceived := (err != nil)

		if errReceived != tc.errExpected {
			t.Fatalf("%s(%.1f): unexpected error value: %v", tc.fnName, tc.a, err)
		}

		if got != tc.want {
			t.Errorf("%s(%.1f): got %.1f want %.1f", tc.fnName, tc.a, got, tc.want)
		}
	}
}
