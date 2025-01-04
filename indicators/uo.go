package indicators

import (
	"errors"
	"math"
)

/*
Ultimate Oscillator (UO):
-------------------------------------------
Inputs: High, Low, Close arrays, and 3 lookback periods (commonly 7, 14, 28).
Steps:
 1. Compute True Low  = min(Low[i],  Close[i-1])
    Compute True High = max(High[i], Close[i-1])
 2. BP (Buying Pressure) = Close[i] - True Low
 3. TR (True Range)      = True High - True Low
 4. Average each BP & TR over short, medium, long windows
    AvgShort  = sum(BP over shortWindow)  / sum(TR over shortWindow)
    AvgMedium = sum(BP over mediumWindow) / sum(TR over mediumWindow)
    AvgLong   = sum(BP over longWindow)   / sum(TR over longWindow)
 5. UltimateOscillator = 100 * [ (w1 * AvgShort) + (w2 * AvgMedium) + (w3 * AvgLong ) ] / (w1 + w2 + w3)
    Usually, w1=4, w2=2, w3=1

-------------------------------------------
*/
type UltimateOscillator struct {
	ShortPeriod  int
	MediumPeriod int
	LongPeriod   int
	Weight1      float64
	Weight2      float64
	Weight3      float64
}

// NewUltimateOscillator constructs the UO with typical default periods (7, 14, 28)
// and default weights (4, 2, 1).
func NewUltimateOscillator(shortP, mediumP, longP int) *UltimateOscillator {
	return &UltimateOscillator{
		ShortPeriod:  shortP,
		MediumPeriod: mediumP,
		LongPeriod:   longP,
		Weight1:      4.0,
		Weight2:      2.0,
		Weight3:      1.0,
	}
}

// Calculate returns a slice of UO values, each in [0..100] (typically).
// The first (longPeriod-1) data points might be 0 or NaN, since we need
// at least 'longPeriod' bars to compute the full UO.
func (u *UltimateOscillator) Calculate(highs, lows, closes []float64) ([]float64, error) {
	n := len(highs)
	if n != len(lows) || n != len(closes) {
		return nil, errors.New("highs, lows, closes must have the same length")
	}
	if n < u.LongPeriod {
		return nil, errors.New("not enough data for Ultimate Oscillator")
	}

	uoValues := make([]float64, n)

	// For each bar, compute the components we need: True Low, True High, BP, TR
	trueLow := make([]float64, n)
	trueHigh := make([]float64, n)
	bp := make([]float64, n) // Buying Pressure
	tr := make([]float64, n) // True Range

	// For bar 0, we can’t look back, so do a minimal initialization
	trueLow[0] = lows[0]
	trueHigh[0] = highs[0]
	bp[0] = closes[0] - lows[0]
	tr[0] = highs[0] - lows[0]

	for i := 1; i < n; i++ {
		trueLow[i] = math.Min(lows[i], closes[i-1])
		trueHigh[i] = math.Max(highs[i], closes[i-1])
		bp[i] = closes[i] - trueLow[i]
		tr[i] = trueHigh[i] - trueLow[i]
	}

	// Helper function for sums
	sumArray := func(arr []float64, start, end int) float64 {
		var s float64
		for i := start; i <= end; i++ {
			s += arr[i]
		}
		return s
	}

	for i := 0; i < u.LongPeriod-1; i++ {
		// Not enough bars for full calculation
		uoValues[i] = 0 // or math.NaN()
	}

	// from i = (longPeriod - 1) onward, we can compute UO
	for i := u.LongPeriod - 1; i < n; i++ {
		// short window range: [i - shortPeriod+1 .. i]
		startShort := i - u.ShortPeriod + 1
		if startShort < 0 {
			startShort = 0
		}
		// medium window range: [i - mediumPeriod+1 .. i]
		startMed := i - u.MediumPeriod + 1
		if startMed < 0 {
			startMed = 0
		}
		// long window range: [i - longPeriod+1 .. i]
		startLong := i - u.LongPeriod + 1
		if startLong < 0 {
			startLong = 0
		}

		sumBPShort := sumArray(bp, startShort, i)
		sumTRShort := sumArray(tr, startShort, i)

		sumBPMed := sumArray(bp, startMed, i)
		sumTRMed := sumArray(tr, startMed, i)

		sumBPLong := sumArray(bp, startLong, i)
		sumTRLong := sumArray(tr, startLong, i)

		var avgShort, avgMed, avgLong float64
		if sumTRShort == 0 {
			avgShort = 0
		} else {
			avgShort = sumBPShort / sumTRShort
		}
		if sumTRMed == 0 {
			avgMed = 0
		} else {
			avgMed = sumBPMed / sumTRMed
		}
		if sumTRLong == 0 {
			avgLong = 0
		} else {
			avgLong = sumBPLong / sumTRLong
		}

		numer := (u.Weight1 * avgShort) + (u.Weight2 * avgMed) + (u.Weight3 * avgLong)
		denom := (u.Weight1 + u.Weight2 + u.Weight3)

		uoValues[i] = 100.0 * (numer / denom)
	}

	return uoValues, nil
}
