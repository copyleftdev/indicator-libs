package tests

import (
	"math"
	"testing"

	"github.com/copyleftdev/indicator-libs/indicators"
)

func TestCCI(t *testing.T) {
	highs := []float64{10, 11, 12, 13, 14, 16, 17, 19}
	lows := []float64{9, 9, 10, 12, 12, 14, 15, 17}
	closes := []float64{9, 10, 11, 12, 13, 15, 16, 18}

	cciCalc := indicators.NewCCI(3)
	cciVals, err := cciCalc.Calculate(highs, lows, closes)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(cciVals) != len(highs) {
		t.Errorf("expected output length %d, got %d", len(highs), len(cciVals))
	}

	// Basic sanity check: final value shouldn't be NaN if there's enough data.
	finalCCI := cciVals[len(cciVals)-1]
	if math.IsNaN(finalCCI) {
		t.Error("final CCI value is NaN, expected a valid float")
	}
}
