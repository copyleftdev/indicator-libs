package indicators

import (
	"errors"
	"math"

	"gonum.org/v1/gonum/stat"
)

/*
CCI (Commodity Channel Index):
-----------------------------------------
Typical Price (TP) = (High + Low + Close) / 3
Moving Average of TP (SMA or sometimes EMA)
Mean Deviation = average( |TP[i] - movingAvgTP| ) over the window
CCI = (TP - movingAvgTP) / (0.015 * meanDeviation)
-----------------------------------------
Default window is often 14, but can vary.
*/
type CCI struct {
	Window int
}

func NewCCI(window int) *CCI {
	return &CCI{Window: window}
}

// Calculate returns a slice of CCI values. It expects three slices: highs, lows, closes.
func (c *CCI) Calculate(highs, lows, closes []float64) ([]float64, error) {
	if len(highs) != len(lows) || len(lows) != len(closes) {
		return nil, errors.New("highs, lows, and closes must have the same length")
	}
	if len(highs) < c.Window {
		return nil, errors.New("not enough data for CCI")
	}

	length := len(highs)

	// Compute Typical Price for each bar.
	typicalPrices := make([]float64, length)
	for i := 0; i < length; i++ {
		typicalPrices[i] = (highs[i] + lows[i] + closes[i]) / 3.0
	}

	// We’ll compute a Simple Moving Average of TP over Window.
	// Then we need the Mean Deviation of each bar’s TP from that average.
	cciValues := make([]float64, length)

	for i := 0; i < c.Window-1; i++ {
		cciValues[i] = 0 // or math.NaN(), up to you
	}

	for i := c.Window - 1; i < length; i++ {
		windowTP := typicalPrices[i-c.Window+1 : i+1]
		avgTP := stat.Mean(windowTP, nil)

		// Calculate mean deviation: average of abs(TP - avgTP)
		var sumDev float64
		for _, tp := range windowTP {
			sumDev += math.Abs(tp - avgTP)
		}
		meanDev := sumDev / float64(c.Window)

		// If meanDev is 0, CCI is 0 or could be NaN, but typically 0 is returned.
		if meanDev == 0 {
			cciValues[i] = 0
		} else {
			cciValues[i] = (typicalPrices[i] - avgTP) / (0.015 * meanDev)
		}
	}

	return cciValues, nil
}
