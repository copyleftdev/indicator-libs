package tests

import (
	"math"
	"testing"

	"github.com/copyleftdev/indicator-libs/indicators"
)

func TestATR(t *testing.T) {
	// Minimal example data
	highs := []float64{10, 11, 13, 14, 15, 17, 17}
	lows := []float64{9, 9, 11, 13, 14, 15, 16}
	closes := []float64{9, 10, 12, 14, 14, 16, 17}

	// Window of 3 for demonstration
	atrCalc := indicators.NewATR(3)
	got, err := atrCalc.Calculate(highs, lows, closes)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(got) != len(highs) {
		t.Errorf("expected output length %d, got %d", len(highs), len(got))
	}

	// We won't do an exact numeric reference here, but let's ensure the final value isn't NaN.
	finalVal := got[len(got)-1]
	if math.IsNaN(finalVal) {
		t.Error("final ATR value should not be NaN")
	}
}
