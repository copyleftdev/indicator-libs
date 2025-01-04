package tests

import (
	"math"
	"testing"

	"github.com/copyleftdev/indicator-libs/indicators"
)

func TestUltimateOscillator(t *testing.T) {
	// Minimal example data
	highs := []float64{10, 12, 13, 13, 14, 15, 15, 16}
	lows := []float64{9, 11, 11, 12, 12, 13, 14, 15}
	closes := []float64{9.5, 11.5, 12.5, 12.5, 13.5, 14.5, 14.0, 15.0}

	uoCalc := indicators.NewUltimateOscillator(3, 5, 7) // small windows for the test
	results, err := uoCalc.Calculate(highs, lows, closes)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(results) != len(highs) {
		t.Errorf("output length mismatch, got %d, want %d", len(results), len(highs))
	}

	// We won't do a precise numeric match here, but we can check that the final result isn't NaN.
	finalVal := results[len(results)-1]
	if math.IsNaN(finalVal) {
		t.Error("expected final UO not to be NaN")
	}
}
