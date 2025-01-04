package indicators

import (
	"errors"
	"math"
)

// WilliamsR calculates the Williams %R indicator over a specified window.
// Formula (for day i):
//
//	HighestHigh = max of High over last 'window' bars
//	LowestLow   = min of Low  over last 'window' bars
//	%R = (HighestHigh - Close[i]) / (HighestHigh - LowestLow) * (-100)
//
// The resulting values typically range from 0 to -100.
type WilliamsR struct {
	Window int
}

// NewWilliamsR returns a new instance with the desired lookback window (e.g., 14).
func NewWilliamsR(window int) *WilliamsR {
	return &WilliamsR{Window: window}
}

// Calculate returns a slice of Williams %R values.
// The first (window-1) values may be set to 0 or math.NaN().
func (w *WilliamsR) Calculate(highs, lows, closes []float64) ([]float64, error) {
	if len(highs) != len(lows) || len(lows) != len(closes) {
		return nil, errors.New("highs, lows, and closes must have the same length")
	}
	if len(highs) < w.Window {
		return nil, errors.New("not enough data for Williams %R")
	}

	length := len(highs)
	wprValues := make([]float64, length)

	// For the first (window-1) elements, we can't compute W%R.
	for i := 0; i < w.Window-1; i++ {
		wprValues[i] = 0 // or math.NaN(), depending on your convention
	}

	// Calculate Williams %R for i >= (window-1)
	for i := w.Window - 1; i < length; i++ {
		// Determine the highest high and lowest low over the last 'window' bars
		highestHigh := -math.MaxFloat64
		lowestLow := math.MaxFloat64

		start := i - w.Window + 1
		for j := start; j <= i; j++ {
			if highs[j] > highestHigh {
				highestHigh = highs[j]
			}
			if lows[j] < lowestLow {
				lowestLow = lows[j]
			}
		}

		// Compute %R
		denominator := highestHigh - lowestLow
		if denominator == 0 {
			// Avoid division by zero; if range is zero, price hasn't moved
			wprValues[i] = 0
		} else {
			wprValues[i] = (highestHigh - closes[i]) / denominator * -100.0
		}
	}

	return wprValues, nil
}
