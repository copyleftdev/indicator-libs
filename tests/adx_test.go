package tests

import (
	"math"
	"testing"

	"github.com/copyleftdev/indicator-libs/indicators"
)

func TestADX(t *testing.T) {
	// Minimal example data
	highs := []float64{10, 11, 13, 14, 15, 18, 20}
	lows := []float64{9, 10, 11, 13, 14, 15, 16}
	closes := []float64{9, 10, 12, 14, 14, 16, 19}

	adxCalc := indicators.NewADX(3)
	adxVals, plusDI, minusDI, err := adxCalc.Calculate(highs, lows, closes)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(adxVals) != len(highs) || len(plusDI) != len(highs) || len(minusDI) != len(highs) {
		t.Errorf("all returned slices must match input length")
	}

	// Just a quick sanity check that final values aren't NaN.
	idx := len(adxVals) - 1
	if math.IsNaN(adxVals[idx]) || math.IsNaN(plusDI[idx]) || math.IsNaN(minusDI[idx]) {
		t.Error("final ADX / +DI / -DI is NaN, expected a valid value")
	}
}
