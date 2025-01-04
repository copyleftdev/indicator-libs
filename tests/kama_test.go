package tests

import (
	"math"
	"testing"

	"github.com/copyleftdev/indicator-libs/indicators"
)

func TestKAMA(t *testing.T) {
	// A small sample of price data. Real usage would typically be many bars.
	// We'll need at least 'erPeriod' bars to test effectively.
	prices := []float64{
		10.0, 10.2, 10.5, 10.3, 10.4,
		10.8, 10.7, 10.9, 11.0, 11.2,
		11.5, 11.3, 11.1, 11.0, 10.9, 11.0, 11.2, 11.5, 11.7, 12.0,
	}

	// Typical defaults: ERPeriod=10, FastPeriod=2, SlowPeriod=30
	kamaCalc := indicators.NewKAMA(10, 2, 30)
	kamaVals, err := kamaCalc.Calculate(prices)
	if err != nil {
		t.Fatalf("KAMA calculation failed: %v", err)
	}

	// Ensure lengths match
	if len(kamaVals) != len(prices) {
		t.Errorf("KAMA output length (%d) != prices length (%d)", len(kamaVals), len(prices))
	}

	// Check for NaNs
	for i, v := range kamaVals {
		if math.IsNaN(v) {
			t.Errorf("KAMA[%d] is NaN; expected a valid number", i)
		}
	}

	// Optionally, compare the final KAMA to an expected reference
	// if you have one. For demonstration, we'll just log it.
	t.Logf("Final KAMA value: %.5f", kamaVals[len(kamaVals)-1])
}
