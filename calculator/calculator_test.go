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
	var want float64 = 4
	got := calculator.Add(2, 2)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestSubstract(t *testing.T) {
	t.Parallel()
	var want float64 = 2
	got := calculator.Subtract(4, 2)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}
