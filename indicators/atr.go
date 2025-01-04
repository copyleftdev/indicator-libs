package indicators

import (
	"errors"
	"math"
)

// ATR calculates the Average True Range for a given window using Wilder's smoothing.
type ATR struct {
	Window int
}

func NewATR(window int) *ATR {
	return &ATR{Window: window}
}

// Calculate expects three slices: highs, lows, closes. Returns a slice of ATR values.
// For the first (window-1) data points, ATR is set to 0. Then Wilder smoothing is applied.
func (a *ATR) Calculate(highs, lows, closes []float64) ([]float64, error) {
	if len(highs) != len(lows) || len(lows) != len(closes) {
		return nil, errors.New("highs, lows, and closes must have the same length")
	}
	if len(highs) < a.Window {
		return nil, errors.New("not enough data for ATR")
	}

	length := len(highs)
	tr := make([]float64, length)

	// First TR can simply be high - low, or incorporate previous day close if needed
	tr[0] = highs[0] - lows[0]

	// True Range from day 1 onward
	// TR = max(high[i] - low[i], |high[i] - close[i-1]|, |low[i] - close[i-1]|)
	for i := 1; i < length; i++ {
		range1 := highs[i] - lows[i]
		range2 := math.Abs(highs[i] - closes[i-1])
		range3 := math.Abs(lows[i] - closes[i-1])
		tr[i] = max(range1, max(range2, range3)) // uses max from utils.go
	}

	// Allocate ATR output
	atr := make([]float64, length)

	// First (window-1) are set to 0 (or could be math.NaN())
	for i := 0; i < a.Window-1; i++ {
		atr[i] = 0
	}

	// Initial average TR for index window-1
	var sumTR float64
	for i := 0; i < a.Window; i++ {
		sumTR += tr[i]
	}
	atr[a.Window-1] = sumTR / float64(a.Window)

	// Wilder's smoothing
	for i := a.Window; i < length; i++ {
		prev := atr[i-1]
		atr[i] = ((prev * float64(a.Window-1)) + tr[i]) / float64(a.Window)
	}

	return atr, nil
}
