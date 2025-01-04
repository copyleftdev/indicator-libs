package indicators

import (
	"errors"
)

// OBV (On-Balance Volume) measures buying and selling pressure as a cumulative indicator.
// Formula:
//
//	OBV[0] = volume[0] (or 0, depending on your preference)
//	For i > 0:
//	    if close[i] > close[i-1] => OBV[i] = OBV[i-1] + volume[i]
//	    if close[i] < close[i-1] => OBV[i] = OBV[i-1] - volume[i]
//	    else => OBV[i] = OBV[i-1]
type OBV struct{}

// NewOBV returns a new instance of OBV.
func NewOBV() *OBV {
	return &OBV{}
}

// Calculate returns a slice of OBV values.
// It expects closes and volumes of the same length.
func (o *OBV) Calculate(closes, volumes []float64) ([]float64, error) {
	if len(closes) != len(volumes) {
		return nil, errors.New("closes and volumes must have the same length")
	}
	if len(closes) == 0 {
		return nil, errors.New("empty data set")
	}

	obv := make([]float64, len(closes))

	// You can set the first OBV to volumes[0] or 0. We'll use volumes[0] here.
	obv[0] = volumes[0]

	for i := 1; i < len(closes); i++ {
		if closes[i] > closes[i-1] {
			obv[i] = obv[i-1] + volumes[i]
		} else if closes[i] < closes[i-1] {
			obv[i] = obv[i-1] - volumes[i]
		} else {
			// price is unchanged, OBV doesn't move
			obv[i] = obv[i-1]
		}
	}

	return obv, nil
}
