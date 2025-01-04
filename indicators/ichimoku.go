package indicators

import (
	"errors"
	"math"
)

// Ichimoku holds period configurations for calculating
// Tenkan-sen, Kijun-sen, Senkou Span A/B, and Chikou Span.
type Ichimoku struct {
	TenkanPeriod int
	KijunPeriod  int
	SenkouPeriod int
	Shift        int
}

// NewIchimoku creates an Ichimoku instance with common default values:
// Tenkan=9, Kijun=26, Senkou=52, Shift=26.
func NewIchimoku(tenkan, kijun, senkou, shift int) *Ichimoku {
	return &Ichimoku{
		TenkanPeriod: tenkan,
		KijunPeriod:  kijun,
		SenkouPeriod: senkou,
		Shift:        shift,
	}
}

// Calculate computes five slices:
//  1. Tenkan-sen
//  2. Kijun-sen
//  3. Senkou Span A
//  4. Senkou Span B
//  5. Chikou Span
//
// All slices have the same length as input (high, low, close).
// For days where a value can't be computed, we insert math.NaN().
func (i *Ichimoku) Calculate(high, low, close []float64) (
	[]float64, []float64, []float64, []float64, []float64, error,
) {
	n := len(high)
	if n != len(low) || n != len(close) {
		return nil, nil, nil, nil, nil, errors.New("high, low, close must have the same length")
	}
	if n < i.SenkouPeriod {
		return nil, nil, nil, nil, nil, errors.New("not enough data for Ichimoku: need >= senkouPeriod bars")
	}

	tenkanVals := make([]float64, n)
	kijunVals := make([]float64, n)
	spanA := make([]float64, n)
	spanB := make([]float64, n)
	chikouSpan := make([]float64, n)

	// Initialize all to NaN
	for idx := 0; idx < n; idx++ {
		tenkanVals[idx] = math.NaN()
		kijunVals[idx] = math.NaN()
		spanA[idx] = math.NaN()
		spanB[idx] = math.NaN()
		chikouSpan[idx] = math.NaN()
	}

	// Helper function to get min and max in a window
	getMinMax := func(start, end int) (float64, float64) {
		highMax := -math.MaxFloat64
		lowMin := math.MaxFloat64
		for j := start; j <= end; j++ {
			if high[j] > highMax {
				highMax = high[j]
			}
			if low[j] < lowMin {
				lowMin = low[j]
			}
		}
		return lowMin, highMax
	}

	// Compute Tenkan-sen (Conversion Line)
	// For index >= i.TenkanPeriod-1
	for idx := i.TenkanPeriod - 1; idx < n; idx++ {
		start := idx - (i.TenkanPeriod - 1)
		lowMin, highMax := getMinMax(start, idx)
		tenkanVals[idx] = (highMax + lowMin) / 2.0
	}

	// Compute Kijun-sen (Base Line)
	for idx := i.KijunPeriod - 1; idx < n; idx++ {
		start := idx - (i.KijunPeriod - 1)
		lowMin, highMax := getMinMax(start, idx)
		kijunVals[idx] = (highMax + lowMin) / 2.0
	}

	// Senkou Span A = (Tenkan + Kijun) / 2, shifted forward by i.Shift periods
	// We can only compute A if both tenkan and kijun are valid at idx.
	// Then we place that value at [idx + i.Shift], if it exists.
	for idx := 0; idx < n; idx++ {
		if math.IsNaN(tenkanVals[idx]) || math.IsNaN(kijunVals[idx]) {
			continue
		}
		forwardIndex := idx + i.Shift
		if forwardIndex < n {
			spanA[forwardIndex] = (tenkanVals[idx] + kijunVals[idx]) / 2.0
		}
	}

	// Senkou Span B = (highest high + lowest low) / 2 over i.SenkouPeriod,
	// also shifted forward by i.Shift
	for idx := i.SenkouPeriod - 1; idx < n; idx++ {
		start := idx - (i.SenkouPeriod - 1)
		lowMin, highMax := getMinMax(start, idx)
		forwardIndex := idx + i.Shift
		if forwardIndex < n {
			spanB[forwardIndex] = (highMax + lowMin) / 2.0
		}
	}

	// Chikou Span = current close, shifted backward by i.Shift
	for idx := 0; idx < n; idx++ {
		backIndex := idx - i.Shift
		if backIndex >= 0 {
			chikouSpan[backIndex] = close[idx]
		}
	}

	return tenkanVals, kijunVals, spanA, spanB, chikouSpan, nil
}
