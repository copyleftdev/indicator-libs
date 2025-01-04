package indicators

import (
	"errors"

	"gonum.org/v1/gonum/stat"
)

// SMA computes simple moving average.
type SMA struct {
	Window int
}

func NewSMA(window int) *SMA {
	return &SMA{Window: window}
}

func (s *SMA) Calculate(prices []float64) ([]float64, error) {
	if len(prices) < s.Window {
		return nil, errors.New("not enough data for SMA")
	}
	out := make([]float64, len(prices))
	for i := 0; i < s.Window-1; i++ {
		out[i] = 0
	}
	for i := s.Window - 1; i < len(prices); i++ {
		w := prices[i-s.Window+1 : i+1]
		out[i] = stat.Mean(w, nil)
	}
	return out, nil
}
