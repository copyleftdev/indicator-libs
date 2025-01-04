package tests

import (
	"math"
	"testing"

	"github.com/copyleftdev/indicator-libs/indicators"
)

func TestKeltnerChannels(t *testing.T) {
	// Small sample data for demonstration.
	// In reality, more bars would give a more reliable test.
	highs := []float64{10, 10.5, 11, 11.2, 11.1, 11.3, 11.7, 12.0}
	lows := []float64{9, 10.0, 10.2, 10.6, 10.8, 10.9, 11.1, 11.6}
	closes := []float64{9.5, 10.2, 10.5, 11.0, 10.9, 11.1, 11.4, 11.8}

	// Suppose we choose typical Keltner settings: EMA period=3, ATR period=3, multiplier=1.5
	// These are small for testing on a short data set.
	kc := indicators.NewKeltnerChannels(3, 3, 1.5)
	mid, up, low, err := kc.Calculate(highs, lows, closes)
	if err != nil {
		t.Fatalf("Calculate returned error: %v", err)
	}

	// All outputs should match input length
	if len(mid) != len(highs) || len(up) != len(highs) || len(low) != len(highs) {
		t.Errorf("output slice lengths must match input: got mid=%d, up=%d, low=%d, want=%d",
			len(mid), len(up), len(low), len(highs))
	}

	// Basic check: ensure no NaNs in final results (some early bars might be 0 or partial due to warm-up).
	for i := range mid {
		if math.IsNaN(mid[i]) {
			t.Errorf("mid[%d] is NaN, expected a valid float", i)
		}
		if math.IsNaN(up[i]) {
			t.Errorf("up[%d] is NaN, expected a valid float", i)
		}
		if math.IsNaN(low[i]) {
			t.Errorf("low[%d] is NaN, expected a valid float", i)
		}
	}

	// If you want to compare to known references, place them here.
	t.Logf("Final Keltner Channels => mid=%.3f, up=%.3f, low=%.3f",
		mid[len(mid)-1], up[len(up)-1], low[len(low)-1])
}
