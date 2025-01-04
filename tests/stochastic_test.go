// tests/stochastic_test.go
package tests

import (
	"math"
	"testing"

	"github.com/copyleftdev/indicator-libs/indicators"
)

func TestStochastic(t *testing.T) {
	// Example: 7 data points
	highs := []float64{5, 6, 7, 8, 9, 11, 12}
	lows := []float64{1, 2, 3, 3, 4, 5, 7}
	closes := []float64{3, 5, 6, 7, 8, 10, 12}

	s := indicators.NewStochasticOscillator(3, 2)
	kVals, dVals, err := s.Calculate(highs, lows, closes)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(kVals) != len(highs) || len(dVals) != len(highs) {
		t.Errorf("stochastic outputs must match input length")
	}

	// Check final (or near-final) for a non-NaN
	if math.IsNaN(kVals[len(kVals)-1]) || math.IsNaN(dVals[len(dVals)-1]) {
		t.Error("expected valid final %K and %D values, got NaN")
	}
}
