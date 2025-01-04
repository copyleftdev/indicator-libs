package indicators

import (
	"errors"
)

// MoneyFlowIndex (MFI) uses volume data to weigh price changes,
// indicating buying or selling pressure over a specified window.
type MoneyFlowIndex struct {
	Window int
}

// NewMFI returns a new MFI instance with the specified window (often 14).
func NewMFI(window int) *MoneyFlowIndex {
	return &MoneyFlowIndex{Window: window}
}

// Calculate expects three slices (high, low, close) and a volume slice (all same length).
// Returns a slice of MFI values. The first (window-1) values can be 0 or NaN.
func (m *MoneyFlowIndex) Calculate(high, low, close, volume []float64) ([]float64, error) {
	length := len(high)
	if length != len(low) || length != len(close) || length != len(volume) {
		return nil, errors.New("high, low, close, volume must have the same length")
	}
	if length < m.Window {
		return nil, errors.New("not enough data for MFI")
	}

	mfiVals := make([]float64, length)

	// Typical Price (TP) and Raw Money Flow
	typicalPrice := make([]float64, length)
	rawMoneyFlow := make([]float64, length)

	for i := 0; i < length; i++ {
		typicalPrice[i] = (high[i] + low[i] + close[i]) / 3.0
		rawMoneyFlow[i] = typicalPrice[i] * volume[i]
	}

	// For convenience, we define arrays to mark if the price increased or decreased
	// from the previous period.
	positiveFlow := make([]float64, length)
	negativeFlow := make([]float64, length)

	// i=0 has no previous day; we skip it or treat it as 0
	for i := 1; i < length; i++ {
		if typicalPrice[i] > typicalPrice[i-1] {
			positiveFlow[i] = rawMoneyFlow[i]
			negativeFlow[i] = 0
		} else if typicalPrice[i] < typicalPrice[i-1] {
			positiveFlow[i] = 0
			negativeFlow[i] = rawMoneyFlow[i]
		} else {
			// No change in typical price => no flow added
			positiveFlow[i] = 0
			negativeFlow[i] = 0
		}
	}

	// Calculate MFI using the window-based sums
	// MFI = 100 - (100 / (1 + moneyRatio)), where moneyRatio = positiveFlowSum / negativeFlowSum
	for i := 0; i < m.Window-1; i++ {
		mfiVals[i] = 0 // or math.NaN()
	}

	// From i=(window-1) onwards, compute sums of posFlow/negFlow for last window
	for i := m.Window - 1; i < length; i++ {
		posSum := 0.0
		negSum := 0.0
		start := i - m.Window + 1
		for j := start; j <= i; j++ {
			posSum += positiveFlow[j]
			negSum += negativeFlow[j]
		}
		if negSum == 0 {
			// If there's no negative flow, money ratio would be infinite => MFI ~ 100
			mfiVals[i] = 100.0
		} else {
			moneyRatio := posSum / negSum
			mfiVals[i] = 100.0 - (100.0 / (1.0 + moneyRatio))
		}
	}

	return mfiVals, nil
}
