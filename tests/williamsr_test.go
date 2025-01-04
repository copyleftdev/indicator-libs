package tests

import (
	"math"
	"testing"

	"github.com/copyleftdev/indicator-libs/indicators"
)

func TestWilliamsR(t *testing.T) {
	// Minimal sample data
	highs := []float64{10, 12, 13, 14, 15, 16, 17}
	lows := []float64{9, 10, 11, 12, 13, 14, 15}
	closes := []float64{9, 11, 12, 13, 14, 15, 16}

	wpr := indicators.NewWilliamsR(3)
	got, err := wpr.Calculate(highs, lows, closes)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(got) != len(highs) {
		t.Errorf("expected length %d, got %d", len(highs), len(got))
	}

	// Check final value for a non-NaN
	final := got[len(got)-1]
	if math.IsNaN(final) {
		t.Error("expected last Williams %R to be a valid number, got NaN")
	}
}
