// tests/helpers.go
package tests

import "math"

func nearlyEqual(a, b, epsilon float64) bool {
	return math.Abs(a-b) < epsilon
}
