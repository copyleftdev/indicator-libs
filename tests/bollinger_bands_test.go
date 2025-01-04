// tests/bollinger_bands_test.go
package tests

import (
	"testing"

	"github.com/copyleftdev/indicator-libs/indicators"
)

func TestBollingerBands(t *testing.T) {
	data := []float64{10, 11, 12, 13, 14, 15}
	bb := indicators.NewBollingerBands(3, 2.0)
	mid, up, low, err := bb.Calculate(data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(mid) != len(data) || len(up) != len(data) || len(low) != len(data) {
		t.Errorf("bollinger slices must match input length")
	}

	// Check that the final band is not obviously incorrect
	finalIndex := len(data) - 1
	if up[finalIndex] < mid[finalIndex] || low[finalIndex] > mid[finalIndex] {
		t.Error("upper band should be >= mid, and lower band <= mid")
	}
}
