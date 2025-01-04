package indicators

import (
	"errors"
	"math"
)

/*
Kaufman's Adaptive Moving Average (KAMA)

Formula Outline:

 1. Efficiency Ratio (ER):
    ER = abs(Price[t] - Price[t - period + 1]) / sum( abs(Price[i] - Price[i-1]) ) over the last 'period' bars
    - period is often around 10 or more.
    - The numerator is the net price change over the period (absolute value).
    - The denominator is the sum of the absolute day-to-day changes (volatility or “noise”).

2) Smoothing Constant (SC):

  - SC = [ER * (fastSC - slowSC) + slowSC]^2
    where:
    fastSC = 2 / (fastPeriod + 1)
    slowSC = 2 / (slowPeriod + 1)

  - Typical defaults: fastPeriod=2, slowPeriod=30

  - This sets SC between (slowSC^2) and (fastSC^2), depending on ER.

    3. KAMA:
    KAMA[t] = KAMA[t-1] + SC * (Price[t] - KAMA[t-1])

Assumptions & Implementation Details:
- Common 'period' can be 10 for ER calculation (Kaufman suggested 10 in some references).
- Common fastPeriod=2, slowPeriod=30 (these can be tweaked).
- KAMA output is usually “initialized” at the close price or an average at index=0 or somewhere in the first window.

Potential edge cases:
- For the very first value (or first 'period'), we might directly set KAMA to the price or SMA of the first window.
*/
type KAMA struct {
	// Period used for the Efficiency Ratio (ER) calculation.
	ERPeriod int

	// Fast & Slow parameters for Smoothing Constant (SC).
	FastPeriod int
	SlowPeriod int
}

// NewKAMA constructs an instance of KAMA with the specified parameters.
// Common defaults might be: ERPeriod=10, FastPeriod=2, SlowPeriod=30.
func NewKAMA(erPeriod, fastPeriod, slowPeriod int) *KAMA {
	return &KAMA{
		ERPeriod:   erPeriod,
		FastPeriod: fastPeriod,
		SlowPeriod: slowPeriod,
	}
}

// Calculate returns a slice of KAMA values. The length equals the length of 'prices'.
// The first values (before 'ERPeriod') might be less reliable or partially warmed up.
func (k *KAMA) Calculate(prices []float64) ([]float64, error) {
	n := len(prices)
	if n == 0 {
		return nil, errors.New("no price data provided for KAMA")
	}
	if k.ERPeriod < 2 {
		return nil, errors.New("ERPeriod must be >= 2 for KAMA")
	}
	if k.FastPeriod < 1 || k.SlowPeriod < 1 {
		return nil, errors.New("fastPeriod and slowPeriod must be >= 1 for KAMA")
	}
	if n < k.ERPeriod {
		return nil, errors.New("not enough data to compute KAMA for the specified ERPeriod")
	}

	// Pre-calculate the fastSC & slowSC
	fastSC := 2.0 / (float64(k.FastPeriod) + 1.0)
	slowSC := 2.0 / (float64(k.SlowPeriod) + 1.0)

	kama := make([]float64, n)

	// Initialization:
	// We'll just set the first KAMA value as the first price (or the average of the first ERPeriod).
	// It's somewhat arbitrary, but consistent with common practice.
	// For more “accurate” warm-up, you might use an SMA of the first ERPeriod.
	initialVal := prices[0]
	kama[0] = initialVal

	for i := 1; i < n; i++ {
		if i < k.ERPeriod {
			// We haven't reached the full window for ER yet
			// We'll just do a naive approach: continue from the prior KAMA
			// or you might do partial calculations.
			kama[i] = kama[i-1] + fastSC*(prices[i]-kama[i-1])
			continue
		}

		// 1) Efficiency Ratio (ER) over the last ERPeriod
		startIndex := i - k.ERPeriod + 1
		change := math.Abs(prices[i] - prices[startIndex]) // net change

		// sum of absolute intermediate changes
		var volatility float64
		for j := startIndex; j < i; j++ {
			volatility += math.Abs(prices[j+1] - prices[j])
		}

		var er float64
		if volatility == 0 {
			er = 0 // if there's no price movement, ER=0
		} else {
			er = change / volatility
		}

		// 2) Smoothing Constant
		sc := er*(fastSC-slowSC) + slowSC
		sc2 := sc * sc // final smoothing factor

		// 3) KAMA
		prevKama := kama[i-1]
		kama[i] = prevKama + sc2*(prices[i]-prevKama)
	}

	return kama, nil
}
