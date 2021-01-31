// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"errors"
	"math"
)

// Add takes two numbers and returns the result of adding them together.
func Add(a, b float64) float64 {
	return a + b
}

// Subtract takes two numbers and returns the result of subtracting the second from the first.
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply takes two numbers and returns the product
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide takes two numbers and returns the result of dividing the first number by the second. If either number input is zero, an error is returned.
func Divide(a, b float64) (float64, error) {
	if a == 0 || b == 0 {
		return 0, errors.New("invalid input, cannot divide by zero")
	}
	return a / b, nil
}

// Sqrt take s anumber and returns its suare root. If the input number is negative, and error is returned.
func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("invalid input, cannot square root a negative number")
	}
	return math.Sqrt(a), nil
}
