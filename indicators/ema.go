package indicators

import (
	"errors"
)

type EMA struct {
	Window int
}

func NewEMA(window int) *EMA {
	return &EMA{Window: window}
}

func (e *EMA) Calculate(prices []float64) ([]float64, error) {
	if len(prices) < e.Window {
		return nil, errors.New("not enough data for EMA")
	}
	out := make([]float64, len(prices))
	k := 2.0 / (float64(e.Window) + 1.0)
	out[0] = prices[0]
	for i := 1; i < len(prices); i++ {
		out[i] = (prices[i] * k) + (out[i-1] * (1.0 - k))
	}
	return out, nil
}
