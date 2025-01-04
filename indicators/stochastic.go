package indicators

import (
	"errors"
	"math"
)

// StochasticOscillator computes %K and %D.
type StochasticOscillator struct {
	KPeriod int
	DPeriod int
}

func NewStochasticOscillator(k, d int) *StochasticOscillator {
	return &StochasticOscillator{KPeriod: k, DPeriod: d}
}

// Calculate expects highs, lows, closes (same length).
func (s *StochasticOscillator) Calculate(highs, lows, closes []float64) ([]float64, []float64, error) {
	if len(highs) < s.KPeriod || len(lows) < s.KPeriod || len(closes) < s.KPeriod {
		return nil, nil, errors.New("not enough data for Stochastic")
	}
	kVals := make([]float64, len(closes))
	dVals := make([]float64, len(closes))
	for i := 0; i < s.KPeriod-1; i++ {
		kVals[i] = math.NaN()
		dVals[i] = math.NaN()
	}
	for i := s.KPeriod - 1; i < len(closes); i++ {
		lowV := math.MaxFloat64
		highV := -math.MaxFloat64
		for j := i - s.KPeriod + 1; j <= i; j++ {
			if lows[j] < lowV {
				lowV = lows[j]
			}
			if highs[j] > highV {
				highV = highs[j]
			}
		}
		denom := highV - lowV
		if denom == 0 {
			kVals[i] = 100
		} else {
			kVals[i] = (closes[i] - lowV) / denom * 100
		}
	}
	for i := 0; i < s.KPeriod-1+(s.DPeriod-1); i++ {
		dVals[i] = math.NaN()
	}
	for i := s.KPeriod - 1 + (s.DPeriod - 1); i < len(kVals); i++ {
		var sum float64
		for j := i - s.DPeriod + 1; j <= i; j++ {
			sum += kVals[j]
		}
		dVals[i] = sum / float64(s.DPeriod)
	}
	return kVals, dVals, nil
}
