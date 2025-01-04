package indicators

import (
	"errors"
)

// KeltnerChannels holds the configuration for computing the middle line (EMA of typical price)
// and bands offset by a multiple of ATR.
type KeltnerChannels struct {
	EmaPeriod int     // Period for the EMA of typical price (middle line)
	AtrPeriod int     // Period for the ATR calculation
	Mult      float64 // Multiplier for the ATR offset
}

// NewKeltnerChannels creates a KeltnerChannels instance with a given EMA period, ATR period, and ATR multiplier.
func NewKeltnerChannels(emaPeriod, atrPeriod int, multiplier float64) *KeltnerChannels {
	return &KeltnerChannels{
		EmaPeriod: emaPeriod,
		AtrPeriod: atrPeriod,
		Mult:      multiplier,
	}
}

// Calculate returns three slices (middle, upper, lower), each the same length as the inputs.
// - middle line = EMA of typical price
// - upper line  = middle line + (mult * ATR)
// - lower line  = middle line - (mult * ATR)
//
// The function needs high, low, close arrays of equal length. (Volume is not used here.)
func (kc *KeltnerChannels) Calculate(high, low, close []float64) ([]float64, []float64, []float64, error) {
	length := len(high)
	if length != len(low) || length != len(close) {
		return nil, nil, nil, errors.New("high, low, and close must have the same length")
	}
	if length == 0 {
		return nil, nil, nil, errors.New("no price data provided")
	}

	// 1) Compute typical price: (High + Low + Close) / 3
	typicalPrices := make([]float64, length)
	for i := 0; i < length; i++ {
		typicalPrices[i] = (high[i] + low[i] + close[i]) / 3.0
	}

	// 2) Compute EMA of typical price
	emaCalc := NewEMA(kc.EmaPeriod)
	emaValues, err := emaCalc.Calculate(typicalPrices)
	if err != nil {
		return nil, nil, nil, err
	}

	// 3) Compute ATR for the same length
	// Note: We already have an ATR indicator that expects (high, low, close).
	// We'll re-use it here.
	atrCalc := NewATR(kc.AtrPeriod)
	atrValues, err := atrCalc.Calculate(high, low, close)
	if err != nil {
		return nil, nil, nil, err
	}

	// 4) Build the final Keltner Channels
	middle := make([]float64, length)
	upper := make([]float64, length)
	lower := make([]float64, length)

	for i := 0; i < length; i++ {
		m := emaValues[i]
		a := atrValues[i]
		middle[i] = m
		upper[i] = m + kc.Mult*a
		lower[i] = m - kc.Mult*a
	}

	return middle, upper, lower, nil
}
