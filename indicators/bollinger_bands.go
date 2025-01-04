package indicators

import (
	"errors"

	"gonum.org/v1/gonum/stat"
)

type BollingerBands struct {
	Window int
	NumStd float64
}

func NewBollingerBands(window int, numStd float64) *BollingerBands {
	return &BollingerBands{Window: window, NumStd: numStd}
}

// Calculate returns mid, upper, lower bands.
func (b *BollingerBands) Calculate(prices []float64) ([]float64, []float64, []float64, error) {
	if len(prices) < b.Window {
		return nil, nil, nil, errors.New("not enough data for BollingerBands")
	}
	mid := make([]float64, len(prices))
	up := make([]float64, len(prices))
	low := make([]float64, len(prices))
	for i := 0; i < b.Window-1; i++ {
		mid[i] = 0
		up[i] = 0
		low[i] = 0
	}
	for i := b.Window - 1; i < len(prices); i++ {
		w := prices[i-b.Window+1 : i+1]
		mean := stat.Mean(w, nil)
		std := stat.StdDev(w, nil)
		mid[i] = mean
		up[i] = mean + b.NumStd*std
		low[i] = mean - b.NumStd*std
	}
	return mid, up, low, nil
}
