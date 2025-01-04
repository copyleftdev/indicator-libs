package indicators

import (
	"errors"
)

type RSI struct {
	Window int
}

func NewRSI(window int) *RSI {
	return &RSI{Window: window}
}

func (r *RSI) Calculate(prices []float64) ([]float64, error) {
	if len(prices) < r.Window {
		return nil, errors.New("not enough data for RSI")
	}
	gains := make([]float64, len(prices)-1)
	losses := make([]float64, len(prices)-1)
	for i := 1; i < len(prices); i++ {
		diff := prices[i] - prices[i-1]
		if diff > 0 {
			gains[i-1] = diff
		} else {
			losses[i-1] = -diff
		}
	}
	rsiVals := make([]float64, len(prices))
	for i := 0; i < r.Window; i++ {
		rsiVals[i] = 0
	}
	var sumG, sumL float64
	for i := 0; i < r.Window; i++ {
		sumG += gains[i]
		sumL += losses[i]
	}
	avgG := sumG / float64(r.Window)
	avgL := sumL / float64(r.Window)
	rs := float64(0)
	if avgL != 0 {
		rs = avgG / avgL
	}
	rsiVals[r.Window] = 100.0 - (100.0 / (1.0 + rs))
	for i := r.Window + 1; i < len(prices); i++ {
		avgG = ((avgG * float64(r.Window-1)) + gains[i-1]) / float64(r.Window)
		avgL = ((avgL * float64(r.Window-1)) + losses[i-1]) / float64(r.Window)
		if avgL == 0 {
			rsiVals[i] = 100
		} else {
			rs = avgG / avgL
			rsiVals[i] = 100.0 - (100.0 / (1.0 + rs))
		}
	}
	return rsiVals, nil
}
