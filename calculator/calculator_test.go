package calculator_test

import (
	"testing"

	"github.com/28604/for-the-love-of-go/calculator"
)

func TestAdd(t *testing.T) {
	// The *testing.T controls the test execution.

	// Parallel is a standard prelude for tests,
	// telling Go to run the test concurrently with other tests
	// and make the testing process faster.
	t.Parallel()
	// For different test cases, we create a testCase struct
	// to store a pair of input and a desired output.
	// So, rather than writing multiple tests, we write multiple test cases.
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 1, b: 1, want: 2},
		{a: 5, b: 0, want: 5},
	}
	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Add(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()

	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 4, b: 2, want: 2},
		{a: 6, b: 1, want: 5},
		{a: -1, b: -3, want: 2},
	}

	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Subtract(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()

	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: -2, b: 3, want: -6},
		{a: 0, b: 10, want: 0},
	}

	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Multiply(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 2, b: 1, want: 2},
		{a: -6, b: 2, want: -3},
		{a: 10, b: 2, want: 5},
	}
	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)
		if err != nil {
			// t.Fatalf checks validity,
			// if not, exits the test immediately.
			t.Fatalf("Want no error for valid input, got %v", err)
		}
		if tc.want != got {
			// t.Errorf checks correctness,
			// if not, marks the test as failed, but continues executing the test.
			t.Errorf("Divide(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestInvalidInput(t *testing.T) {
	t.Parallel()
	_, err := calculator.Divide(1, 0)
	if err == nil {
		t.Errorf("Want error for invalid input, got nil.")
	}
}
