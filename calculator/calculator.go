// Package calculator does simple calculations.
package calculator

import "errors"

// Add takes two numbers and return the sum.
func Add(a, b float64) float64 {
	return a + b
}

// Subtract takes two numbers and return the difference.
func Subtract(a, b float64) float64 {
	return a - b
}

func Multiply(a, b float64) float64 {
	return a * b
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero not allowed")
	}
	return a / b, nil
}
