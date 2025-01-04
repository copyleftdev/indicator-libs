// tests/ema_test.go
package tests

import (
	"math"
	"testing"

	"github.com/copyleftdev/indicator-libs/indicators"
)

func TestEMA(t *testing.T) {
	ema := indicators.NewEMA(3)
	data := []float64{1, 2, 3, 4, 5}
	got, err := ema.Calculate(data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// We'll do a rough check of final values.
	// For a 3-day EMA with data {1,2,3,4,5}, let's manually compute:
	//
	//  smoothing factor k = 2/(3+1) = 0.5
	//  out[0] = 1
	//  out[1] = (2 * 0.5) + (1 * 0.5) = 1.5
	//  out[2] = (3 * 0.5) + (1.5 * 0.5) = 2.25
	//  out[3] = (4 * 0.5) + (2.25 * 0.5) = 3.125
	//  out[4] = (5 * 0.5) + (3.125 * 0.5) = 4.0625

	want := []float64{1.0, 1.5, 2.25, 3.125, 4.0625}
	eps := 1e-8

	for i, w := range want {
		if math.Abs(got[i]-w) > eps {
			t.Errorf("index %d: got %.5f, want %.5f", i, got[i], w)
		}
	}
}
