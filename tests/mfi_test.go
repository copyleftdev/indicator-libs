package tests

import (
	"math"
	"testing"

	"github.com/copyleftdev/indicator-libs/indicators"
)

func TestMFI(t *testing.T) {
	high := []float64{10, 11, 12, 11, 13, 14, 13}
	low := []float64{9, 10, 11, 10, 12, 12, 12}
	close := []float64{9.5, 10.5, 11.5, 10.5, 12.5, 13.5, 12.5}
	volume := []float64{1000, 1200, 1300, 800, 1500, 2000, 1600}

	mfiCalc := indicators.NewMFI(3)
	mfiVals, err := mfiCalc.Calculate(high, low, close, volume)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(mfiVals) != len(high) {
		t.Errorf("output length mismatch, got %d, want %d", len(mfiVals), len(high))
	}

	// Basic checks:
	// 1) First 2 values (because window=3) might be 0 or NaN
	// 2) Final value shouldn't be NaN if there's enough data
	final := mfiVals[len(mfiVals)-1]
	if math.IsNaN(final) {
		t.Error("final MFI is NaN, expected a valid float")
	}
}
