package tests

import (
	"math"
	"testing"

	"github.com/copyleftdev/indicator-libs/indicators"
)

func TestParabolicSAR(t *testing.T) {
	// Example data (High & Low) for ~8 bars.
	// In real usage, you'd have many more bars.
	highs := []float64{10, 10.5, 10.8, 11, 10.7, 10.4, 10.6, 10.9}
	lows := []float64{9.5, 10.0, 10.2, 10.5, 10.3, 10.1, 10.3, 10.6}

	// Typical defaults: start=0.02, increment=0.02, max=0.2
	psar := indicators.NewParabolicSAR(0.02, 0.02, 0.2)
	sarValues, err := psar.Calculate(highs, lows)
	if err != nil {
		t.Fatalf("unexpected error from Parabolic SAR calculation: %v", err)
	}

	if len(sarValues) != len(highs) {
		t.Errorf("output length mismatch: got %d, expected %d", len(sarValues), len(highs))
	}

	// Basic sanity checks: no NaNs in the final output
	for i, v := range sarValues {
		if math.IsNaN(v) {
			t.Errorf("index %d: got NaN, expected a valid float", i)
		}
	}

	// Optionally, you might compare specific values to a known reference.
	// For demonstration, we'll just print the final SAR for visual inspection.
	t.Logf("Final Parabolic SAR value: %.4f", sarValues[len(sarValues)-1])
}
