package indicators

import (
	"errors"
	"math"
)

/*
SuperTrend Indicator:

1) Compute ATR over a chosen period.
2) Calculate "Basic Upper Band" (UB) and "Basic Lower Band" (LB) for each bar:
   midPrice = (High + Low) / 2
   UB = midPrice + multiplier * ATR
   LB = midPrice - multiplier * ATR

3) Compute "Final" UB and LB:
   finalUB[i] =
       if close[i-1] <= finalUB[i-1] then min(UB[i], finalUB[i-1])
       else UB[i]

   finalLB[i] =
       if close[i-1] >= finalLB[i-1] then max(LB[i], finalLB[i-1])
       else LB[i]

4) Determine SuperTrend direction:
   If superTrend was in "downtrend" (price < supertrend line) and close[i] > finalUB[i],
      then flip to "uptrend" (supertrend line = finalLB[i])
   If superTrend was in "uptrend" (price > supertrend line) and close[i] < finalLB[i],
      then flip to "downtrend" (supertrend line = finalUB[i])
   Otherwise, remain in the same trend.

For simplicity, we store a single line "superTrendLine" and
an integer "trendDirection": 1 (uptrend) or -1 (downtrend).

You need:
- period (for ATR)
- multiplier (commonly 3.0 or so)

Typical usage:
- superTrend flips below price for uptrend
- flips above price for downtrend
- can serve as a trailing stop
*/

type SuperTrend struct {
	Period     int
	Multiplier float64
}

// NewSuperTrend creates a SuperTrend with given ATR period and multiplier.
func NewSuperTrend(period int, multiplier float64) *SuperTrend {
	return &SuperTrend{
		Period:     period,
		Multiplier: multiplier,
	}
}

// Calculate returns three slices (superTrendLine, trendDirection, finalUB, finalLB):
//   - superTrendLine[i]: the main line for the supertrend at bar i
//   - trendDirection[i]:  1 = uptrend, -1 = downtrend
//   - finalUB[i]: final upper band
//   - finalLB[i]: final lower band
//
// The length = len(high). The user can use superTrendLine and trendDirection
// to see where price stands relative to the supertrend.
func (s *SuperTrend) Calculate(high, low, close []float64) (
	[]float64, []int, []float64, []float64, error,
) {
	length := len(high)
	if length != len(low) || length != len(close) {
		return nil, nil, nil, nil, errors.New("high, low, close must have same length")
	}
	if length < s.Period {
		return nil, nil, nil, nil, errors.New("not enough data for SuperTrend period")
	}

	// 1) Compute ATR over the same length
	// Reusing the existing ATR if you have one; else, implementing quickly:
	atrValues, err := computeATR(high, low, close, s.Period)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	superTrendLine := make([]float64, length)
	trendDirection := make([]int, length)
	finalUB := make([]float64, length)
	finalLB := make([]float64, length)

	// 2) Basic Upper/Lower Band
	basicUB := make([]float64, length)
	basicLB := make([]float64, length)

	for i := 0; i < length; i++ {
		midPrice := (high[i] + low[i]) / 2.0
		basicUB[i] = midPrice + s.Multiplier*atrValues[i]
		basicLB[i] = midPrice - s.Multiplier*atrValues[i]
	}

	// 3) Final UB and LB with smoothing
	// We need to set initial finalUB[0], finalLB[0],
	// but they're not fully meaningful until we have a prior bar. We'll just set them to the basic bands for the first bar.
	finalUB[0] = basicUB[0]
	finalLB[0] = basicLB[0]
	trendDirection[0] = 1          // arbitrarily assume uptrend at first bar
	superTrendLine[0] = finalLB[0] // or finalUB[0], up to your preference

	for i := 1; i < length; i++ {
		// finalUB
		if close[i-1] <= finalUB[i-1] {
			finalUB[i] = math.Min(basicUB[i], finalUB[i-1])
		} else {
			finalUB[i] = basicUB[i]
		}

		// finalLB
		if close[i-1] >= finalLB[i-1] {
			finalLB[i] = math.Max(basicLB[i], finalLB[i-1])
		} else {
			finalLB[i] = basicLB[i]
		}

		// 4) Decide trend
		// If the previous trend was up and close now < finalLB => flip to down
		// If the previous trend was down and close now > finalUB => flip to up
		if trendDirection[i-1] == 1 { // was up
			if close[i] <= finalLB[i] {
				trendDirection[i] = -1
				superTrendLine[i] = finalUB[i]
			} else {
				trendDirection[i] = 1
				superTrendLine[i] = finalLB[i]
			}
		} else { // was down
			if close[i] >= finalUB[i] {
				trendDirection[i] = 1
				superTrendLine[i] = finalLB[i]
			} else {
				trendDirection[i] = -1
				superTrendLine[i] = finalUB[i]
			}
		}
	}

	return superTrendLine, trendDirection, finalUB, finalLB, nil
}

// A quick internal ATR calculation to keep SuperTrend self-contained.
// If you already have an ATR function, you can reuse that instead.
func computeATR(high, low, close []float64, period int) ([]float64, error) {
	length := len(high)
	if length != len(low) || length != len(close) {
		return nil, errors.New("mismatched slice lengths for ATR")
	}
	if period < 1 || period > length {
		return nil, errors.New("invalid ATR period")
	}

	// TR array
	tr := make([]float64, length)
	// First TR is high[0] - low[0]
	tr[0] = high[0] - low[0]

	// Compute TR
	for i := 1; i < length; i++ {
		range1 := high[i] - low[i]
		range2 := math.Abs(high[i] - close[i-1])
		range3 := math.Abs(low[i] - close[i-1])
		tr[i] = math.Max(range1, math.Max(range2, range3))
	}

	// ATR values
	atr := make([]float64, length)

	// First ATR can be simple sum of first "period" TR / period or just TR[0].
	// We'll do Wilder's smoothing to be consistent with typical usage.
	// For i=0, it's just TR[0].
	// Then: ATR[i] = ((ATR[i-1] * (period - 1)) + TR[i]) / period
	sumTR := 0.0
	// sum up first period TR
	for i := 0; i < period; i++ {
		sumTR += tr[i]
	}
	atr[period-1] = sumTR / float64(period)

	// proceed with Wilder's smoothing
	for i := period; i < length; i++ {
		prev := atr[i-1]
		atr[i] = ((prev * float64(period-1)) + tr[i]) / float64(period)
	}

	// fill earliest bars (< period-1) with 0 or partial
	for i := 0; i < period-1; i++ {
		atr[i] = 0.0 // or math.NaN()
	}

	return atr, nil
}
