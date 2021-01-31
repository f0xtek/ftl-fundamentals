package calculator_test

import (
	"calculator"
	"math/rand"
	"testing"
)

type testCase = struct {
	name       string
	a, b, want float64
}

type testCaseWithError = struct {
	name        string
	a, b, want  float64
	errExpected bool
}

func TestAdd(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{name: "Two positive numbers which sum to a positive", a: 2, b: 2, want: 4},
		{name: "A positive number and zero which sum to the positive number", a: 5, b: 0, want: 5},
		{name: "A negative and a positive number which sum to a negative number", a: -10, b: 5, want: -5},
		{name: "A negative and a positive number whcih sum to a positive number", a: -10, b: 20, want: 10},
		{name: "A positive fractional number and a positive number which sum to a positive fractional number", a: 0.5, b: 10, want: 10.5},
		{name: "A negative fractional number and a positive number which sum to a negative fractional number", a: -20, b: 10, want: -10},
	}

	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("%s. Add(%f, %f): want %f, got %f", tc.name, tc.a, tc.b, tc.want, got)
		}
	}

}

func TestAddRandom(t *testing.T) {
	t.Parallel()

	for i := 0; i < 100; i++ {
		randomNumber1 := rand.Float64()
		randomNumber2 := rand.Float64()
		randomSum := randomNumber1 + randomNumber2
		got := calculator.Add(randomNumber1, randomNumber2)
		if got != randomSum {
			t.Errorf("Add(%f, %f): want %f, got %f", randomNumber1, randomNumber2, randomSum, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{name: "Two positive numbers that subtract to zero", a: 2, b: 2, want: 0},
		{name: "Two positive numbers which subtract to a positive number", a: 10, b: 5, want: 5},
		{name: "Two positive numbers which subtract to a negative number", a: 1, b: 10, want: -9},
		{name: "A positive and negative number which subtract to a negative", a: -1, b: 2, want: -3},
		{name: "Two negative numbers which subtract to 0", a: -1, b: -1, want: 0},
		{name: "Zero and a negative number which subtract to a positive", a: 0, b: -3, want: 3},
		{name: "A positive whole number and positive fractional number which subtract to a positive fractional number", a: 10, b: 0.5, want: 9.5},
	}

	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("%s. Subtract(%f, %f): want %f, got %f", tc.name, tc.a, tc.b, tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{name: "Two positive numbers which multiply to a positive", a: 2, b: 2, want: 4},
		{name: "Zero and a positive number which multiply to zero", a: 0, b: 10, want: 0},
		{name: "A positive whole number and a positive fractional number which multiply to a positive whole number", a: 10, b: 0.5, want: 5},
		{name: "Two negative numbers which multiple to a positive", a: -1, b: -1, want: 1},
	}

	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("%s. Multiply(%f, %f): want %f, got %f", tc.name, tc.a, tc.b, tc.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()

	testCases := []testCaseWithError{
		{name: "Divide by zero to return an error", a: 6, b: 0, want: 999, errExpected: true},
		{name: "Two positive integers which divide to a positive", a: 6, b: 2, want: 3, errExpected: false},
		{name: "A positive and a negative which divide to a negative", a: 6, b: -2, want: -3, errExpected: false},
		{name: "A whole number and fractional number whice divide to a positive", a: 6, b: 0.5, want: 12, errExpected: false},
		{name: "Two negative numbers which divide to a positive", a: -6, b: -2, want: 3, errExpected: false},
	}

	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)
		errReceived := err != nil
		if tc.errExpected != errReceived {
			// We do not need to check the data value because we are expetcing the test to fail when an error is not returned
			t.Fatalf("%s. Divide(%f, %f): Unexpected error status: %v", tc.name, tc.a, tc.b, errReceived)
		} else if !tc.errExpected && tc.want != got {
			t.Errorf("%s. Divide(%f, %f): want %f, got %f", tc.name, tc.a, tc.b, tc.want, got)
		}
	}
}
