package indicators

import (
	"errors"
	"math"
)

// ParabolicSAR holds configuration for the Parabolic SAR calculations.
type ParabolicSAR struct {
	// StartAF is the initial acceleration factor, typically around 0.02.
	StartAF float64
	// IncrementAF is how much the AF increments each time a new extreme is found, e.g. 0.02.
	IncrementAF float64
	// MaxAF is the maximum acceleration factor, e.g. 0.20.
	MaxAF float64
}

// NewParabolicSAR creates a ParabolicSAR struct with given parameters.
func NewParabolicSAR(startAF, incrementAF, maxAF float64) *ParabolicSAR {
	return &ParabolicSAR{
		StartAF:     startAF,
		IncrementAF: incrementAF,
		MaxAF:       maxAF,
	}
}

// Calculate computes the Parabolic SAR for each bar.
// It expects high and low slices of equal length.
// The returned slice has the same length; early bars may be less accurate.
func (p *ParabolicSAR) Calculate(high, low []float64) ([]float64, error) {
	length := len(high)
	if length != len(low) {
		return nil, errors.New("high and low must have the same length")
	}
	if length == 0 {
		return nil, errors.New("no data provided")
	}

	sar := make([]float64, length)

	// 1) Determine initial trend by comparing first two bars
	//    If second bar's low is higher than first bar's low => likely uptrend, etc.
	//    Or you can manually set an initial direction. We'll do an auto-detect:
	upTrend := false
	if length == 1 {
		// With only one bar, can't detect direction; just set SAR to that bar's low or high
		sar[0] = low[0]
		return sar, nil
	}

	if (high[1]+low[1])/2 > (high[0]+low[0])/2 {
		upTrend = true
	}

	// 2) Initialize EP (extreme point) and SAR
	var ep float64
	var af = p.StartAF

	if upTrend {
		sar[0] = low[0] // starting SAR below the first bar if uptrend
		ep = high[0]    // extreme point is the highest high
	} else {
		sar[0] = high[0] // starting SAR above the first bar if downtrend
		ep = low[0]      // extreme point is the lowest low
	}

	// 3) Iterate through each bar to compute SAR
	for i := 1; i < length; i++ {
		prevSAR := sar[i-1]

		// Calculate new SAR
		currSAR := prevSAR + af*(ep-prevSAR)

		if upTrend {
			// Uptrend logic: SAR cannot be above today's or yesterday's low
			limit := math.Min(low[i-1], low[i])
			if currSAR > limit {
				// Flip to downtrend
				upTrend = false
				currSAR = math.Max(high[i-1], high[i]) // new SAR on flip
				sar[i] = currSAR
				af = p.StartAF
				ep = low[i] // new extreme point is current bar's low
			} else {
				// Stay in uptrend
				sar[i] = currSAR
				// Update extreme point if we have a new high
				if high[i] > ep {
					ep = high[i]
					af += p.IncrementAF
					if af > p.MaxAF {
						af = p.MaxAF
					}
				}
			}
		} else {
			// Downtrend logic: SAR cannot be below today's or yesterday's high
			limit := math.Max(high[i-1], high[i])
			if currSAR < limit {
				// Flip to uptrend
				upTrend = true
				currSAR = math.Min(low[i-1], low[i]) // new SAR on flip
				sar[i] = currSAR
				af = p.StartAF
				ep = high[i] // new extreme point is current bar's high
			} else {
				// Stay in downtrend
				sar[i] = currSAR
				// Update extreme point if we have a new low
				if low[i] < ep {
					ep = low[i]
					af += p.IncrementAF
					if af > p.MaxAF {
						af = p.MaxAF
					}
				}
			}
		}
	}

	return sar, nil
}
