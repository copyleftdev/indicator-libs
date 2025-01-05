package tests

import (
	"math"
	"testing"

	"github.com/copyleftdev/indicator-libs/indicators"
)

func TestT3(t *testing.T) {
	// Some synthetic price data
	prices := []float64{
		10.0, 10.1, 9.9, 10.2, 10.5,
		10.3, 10.6, 10.8, 10.7, 11.0,
		11.2, 10.9, 10.8, 10.7, 10.9,
		11.1, 11.2, 11.5, 11.4, 11.3,
	}

	// T3 with period=5, volumeFactor=0.7, for demonstration
	t3Calc := indicators.NewT3(5, 0.7)
	t3vals, err := t3Calc.Calculate(prices)
	if err != nil {
		t.Fatalf("T3 calculation error: %v", err)
	}

	// Basic checks
	if len(t3vals) != len(prices) {
		t.Errorf("length mismatch: got %d, want %d", len(t3vals), len(prices))
	}

	// Ensure no NaNs
	for i := 0; i < len(t3vals); i++ {
		if math.IsNaN(t3vals[i]) {
			t.Errorf("index %d: T3 is NaN", i)
		}
	}

	// Optionally compare final T3 vs. an expected reference from a known source
	t.Logf("Final T3 value = %.4f", t3vals[len(t3vals)-1])
}
