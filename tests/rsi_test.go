// tests/rsi_test.go
package tests

import (
	"math"
	"testing"

	"github.com/copyleftdev/indicator-libs/indicators"
)

func TestRSI(t *testing.T) {
	// Very short dataset to see if RSI runs without error.
	data := []float64{10, 11, 11, 9, 8, 12, 15}
	rsi := indicators.NewRSI(3)
	got, err := rsi.Calculate(data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(got) != len(data) {
		t.Errorf("expected rsi length %d, got %d", len(data), len(got))
	}

	// Check that there's no error on early values, though they might be 0 or near 0.
	// We'll just check final value for a sanity check (not an exact reference).
	final := got[len(got)-1]
	if math.IsNaN(final) {
		t.Error("expected final RSI not to be NaN")
	}
}
