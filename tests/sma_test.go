package tests

import (
	"testing"

	"github.com/copyleftdev/indicator-libs/indicators"
)

func TestSMA(t *testing.T) {
	sma := indicators.NewSMA(3)
	data := []float64{1, 2, 3, 4, 5, 6}
	got, err := sma.Calculate(data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// With window=3, the first two results might be 0 (or 0.0) in this implementation.
	// From index=2 onwards:
	//   i=2 => average of (1,2,3) = 2
	//   i=3 => average of (2,3,4) = 3
	//   i=4 => average of (3,4,5) = 4
	//   i=5 => average of (4,5,6) = 5
	want := []float64{0, 0, 2, 3, 4, 5}

	for i := range want {
		if got[i] != want[i] {
			t.Errorf("index %d: got %v, want %v", i, got[i], want[i])
		}
	}
}
