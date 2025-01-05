package tests

import (
	"math"
	"testing"

	"github.com/copyleftdev/indicator-libs/indicators"
)

func TestSuperTrend(t *testing.T) {
	// Small synthetic dataset for demonstration.
	// For real usage, you'd have more bars and possibly cross-check known supertrend values.
	high := []float64{10, 10.2, 10.3, 10.1, 10.5, 10.8, 11.0, 10.9}
	low := []float64{9, 9.8, 10.0, 9.7, 10.0, 10.4, 10.7, 10.5}
	close := []float64{9.5, 10.0, 10.1, 10.0, 10.3, 10.6, 10.9, 10.7}

	// Example: period=3, multiplier=2.0 (small for testing).
	st := indicators.NewSuperTrend(3, 2.0)
	superTrendLine, trendDir, finalUB, finalLB, err := st.Calculate(high, low, close)
	if err != nil {
		t.Fatalf("SuperTrend calculation failed: %v", err)
	}

	if len(superTrendLine) != len(close) || len(trendDir) != len(close) ||
		len(finalUB) != len(close) || len(finalLB) != len(close) {
		t.Errorf("Output slices must match input length.")
	}

	// Basic check for NaNs.
	for i := range close {
		if math.IsNaN(superTrendLine[i]) {
			t.Errorf("superTrendLine[%d] is NaN", i)
		}
		if trendDir[i] != 1 && trendDir[i] != -1 {
			t.Errorf("trendDir[%d] should be 1 or -1, got %d", i, trendDir[i])
		}
		if math.IsNaN(finalUB[i]) {
			t.Errorf("finalUB[%d] is NaN", i)
		}
		if math.IsNaN(finalLB[i]) {
			t.Errorf("finalLB[%d] is NaN", i)
		}
	}

	// Log some final values for manual inspection.
	t.Logf("Last bar: superTrend=%.3f, trendDir=%d, finalUB=%.3f, finalLB=%.3f",
		superTrendLine[len(close)-1],
		trendDir[len(close)-1],
		finalUB[len(close)-1],
		finalLB[len(close)-1])
}
