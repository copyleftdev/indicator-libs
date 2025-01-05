package indicators

import (
	"errors"
)

/*
T3 Moving Average (Tillson's T3)

Overview:
-----------
T3 is an advanced smoothing technique proposed by Tim Tillson.
It applies multiple sequential EMAs and then combines them
in a way that depends on a "volume factor" (v).

Common defaults:
- period = 14 (or 20, etc.)
- v = 0.7 (or 0.5, 0.8, etc.)

One known formula for T3 uses 6 consecutive EMAs, then combines them:

 Let e1 = EMA(prices, period)
     e2 = EMA(e1, period)
     e3 = EMA(e2, period)
     e4 = EMA(e3, period)
     e5 = EMA(e4, period)
     e6 = EMA(e5, period)

 Then T3 = e6 * c1 + e5 * c2 + e4 * c3 + e3 * c4 + e2 * c5 + e1 * c6

 with certain coefficients that are functions of volume factor v.
 A commonly referenced set:
   c1 = -v^3
   c2 = 3v^2 + 3v^3
   c3 = -6v^2 - 3v - 3v^3
   c4 = 1 + 3v + v^3 + 3v^2
   (some variations exist)
But a more typical approach uses a simpler version (the one below).

Implementation Approach:
-----------
We'll do a 6-EMA approach as documented in certain code references:

  T3 = e6*(1 + v^4)
        - e5*(4v^4)
        + e4*(6v^4)
        - e3*(4v^4)
        + e2*(v^4)

(You can find multiple slightly different “T3 formula” references.
The key idea is that T3 is a weighted combination of multiple
EMAs for deeper smoothing.)

We need to:
1) Compute 6 sequential EMAs of 'prices' using the same 'period'.
2) Combine them using the T3 weighting formula.

Note: Because T3 is a multi-layer smoothing, it has a warm-up period
potentially longer than a single 'period'. Early values might be
less reliable until the chain of EMAs stabilizes.
*/

type T3 struct {
	Period       int     // the EMA period
	VolumeFactor float64 // the volume factor (v), often between 0.5 and 0.8
}

// NewT3 returns a T3 instance with a specified period and volume factor.
func NewT3(period int, volumeFactor float64) *T3 {
	return &T3{
		Period:       period,
		VolumeFactor: volumeFactor,
	}
}

// Calculate returns a slice of T3 values, the same length as prices.
// If the input slice is shorter than 'period', it returns an error.
func (t *T3) Calculate(prices []float64) ([]float64, error) {
	n := len(prices)
	if n < t.Period {
		return nil, errors.New("not enough data for T3 calculation")
	}
	if t.VolumeFactor < 0 || t.VolumeFactor > 1 {
		return nil, errors.New("volumeFactor should be between 0 and 1 for typical T3 usage")
	}

	// We'll compute 6 EMAs in sequence: e1, e2, e3, e4, e5, e6
	e1, err := computeEMA(prices, t.Period)
	if err != nil {
		return nil, err
	}
	e2, err := computeEMA(e1, t.Period)
	if err != nil {
		return nil, err
	}
	e3, err := computeEMA(e2, t.Period)
	if err != nil {
		return nil, err
	}
	e4, err := computeEMA(e3, t.Period)
	if err != nil {
		return nil, err
	}
	e5, err := computeEMA(e4, t.Period)
	if err != nil {
		return nil, err
	}
	e6, err := computeEMA(e5, t.Period)
	if err != nil {
		return nil, err
	}

	t3vals := make([]float64, n)
	v := t.VolumeFactor
	// Weighted combination:
	// T3 = e6*(1 + v^4) - e5*(4v^4) + e4*(6v^4) - e3*(4v^4) + e2*(v^4)
	// This is a commonly referenced formula.
	for i := 0; i < n; i++ {
		t3vals[i] = e6[i]*(1+v*v*v*v) -
			e5[i]*(4*v*v*v*v) +
			e4[i]*(6*v*v*v*v) -
			e3[i]*(4*v*v*v*v) +
			e2[i]*(v*v*v*v)
	}

	return t3vals, nil
}

// computeEMA is a small helper for an EMA on top of another series.
// We can reuse the existing EMA code from your library, but here's
// a self-contained version for T3:
func computeEMA(data []float64, period int) ([]float64, error) {
	if period < 1 {
		return nil, errors.New("period must be >= 1")
	}
	length := len(data)
	if length == 0 {
		return nil, errors.New("empty data for EMA")
	}

	ema := make([]float64, length)
	k := 2.0 / (float64(period) + 1.0)

	// Start EMA at the first data point or an average of first 'period'?
	// We'll just set first EMA = data[0].
	ema[0] = data[0]
	for i := 1; i < length; i++ {
		ema[i] = data[i]*k + ema[i-1]*(1.0-k)
	}

	return ema, nil
}
