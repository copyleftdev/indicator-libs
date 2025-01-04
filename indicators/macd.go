package indicators

import (
	"errors"
)

type MACD struct {
	FastPeriod   int
	SlowPeriod   int
	SignalPeriod int
}

func NewMACD(fast, slow, signal int) *MACD {
	return &MACD{FastPeriod: fast, SlowPeriod: slow, SignalPeriod: signal}
}

// Calculate returns macdLine, signalLine, histogram.
func (m *MACD) Calculate(prices []float64) ([]float64, []float64, []float64, error) {
	if len(prices) < m.SlowPeriod {
		return nil, nil, nil, errors.New("not enough data for MACD")
	}
	fastEMA, err := NewEMA(m.FastPeriod).Calculate(prices)
	if err != nil {
		return nil, nil, nil, err
	}
	slowEMA, err := NewEMA(m.SlowPeriod).Calculate(prices)
	if err != nil {
		return nil, nil, nil, err
	}
	macdLine := make([]float64, len(prices))
	for i := 0; i < len(prices); i++ {
		macdLine[i] = fastEMA[i] - slowEMA[i]
	}
	signalLine, err := NewEMA(m.SignalPeriod).Calculate(macdLine)
	if err != nil {
		return nil, nil, nil, err
	}
	hist := make([]float64, len(prices))
	for i := 0; i < len(prices); i++ {
		hist[i] = macdLine[i] - signalLine[i]
	}
	return macdLine, signalLine, hist, nil
}
