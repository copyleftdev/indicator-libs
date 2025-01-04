package tests

import (
	"testing"

	"github.com/copyleftdev/indicator-libs/indicators"
)

func TestOBV(t *testing.T) {
	closes := []float64{10, 11, 11, 12, 11, 12}
	volumes := []float64{100, 200, 150, 300, 250, 400}

	obvCalc := indicators.NewOBV()
	obvVals, err := obvCalc.Calculate(closes, volumes)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(obvVals) != len(closes) {
		t.Errorf("expected OBV length %d, got %d", len(closes), len(obvVals))
	}

	// We'll do a rough check:
	//
	// obv[0] = volumes[0] = 100
	// close[1] > close[0], so obv[1] = 100 + 200 = 300
	// close[2] == close[1], so obv[2] = 300
	// close[3] > close[2], so obv[3] = 300 + 300 = 600
	// close[4] < close[3], so obv[4] = 600 - 250 = 350
	// close[5] > close[4], so obv[5] = 350 + 400 = 750
	//
	// So the final array should be: [100, 300, 300, 600, 350, 750]

	want := []float64{100, 300, 300, 600, 350, 750}
	for i := range want {
		if obvVals[i] != want[i] {
			t.Errorf("index %d: got %.2f, want %.2f", i, obvVals[i], want[i])
		}
	}
}
