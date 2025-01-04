package tests

import (
	"math"
	"testing"

	"github.com/copyleftdev/indicator-libs/indicators"
)

func TestIchimoku(t *testing.T) {
	// A small dataset for demonstration. Typically, you'd want more bars
	// to see valid forward/backward shifts.
	highs := []float64{10, 12, 13, 13, 14, 15, 15, 16}
	lows := []float64{9, 10, 11, 12, 12, 13, 14, 15}
	closes := []float64{9.5, 11.0, 12.0, 12.5, 13.0, 14.0, 14.5, 15.0}

	// Typical defaults: Tenkan=9, Kijun=26, Senkou=52, shift=26
	// but we'll use smaller values for a short dataset:
	ichimoku := indicators.NewIchimoku(3, 5, 7, 3)
	tenkan, kijun, spanA, spanB, chikou, err := ichimoku.Calculate(highs, lows, closes)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Basic checks
	if len(tenkan) != len(highs) ||
		len(kijun) != len(highs) ||
		len(spanA) != len(highs) ||
		len(spanB) != len(highs) ||
		len(chikou) != len(highs) {
		t.Errorf("all output slices must match input length")
	}

	// Quick sanity check for final indexes
	if math.IsNaN(tenkan[len(tenkan)-1]) {
		t.Log("Tenkan-sen might be NaN if there's insufficient data at the end.")
	}
	if math.IsNaN(spanA[len(spanA)-1]) {
		t.Log("SpanA might be NaN due to forward shift.")
	}
}
