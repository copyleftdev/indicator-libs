package indicators

import (
	"errors"
	"math"
)

// ADX calculates the Average Directional Index, returning ADX, +DI, and -DI slices.
// Uses a default Wilder's smoothing approach. The window is often 14.
type ADX struct {
	Window int
}

func NewADX(window int) *ADX {
	return &ADX{Window: window}
}

// Calculate returns three slices: ADX, +DI, and -DI, each the same length as input.
func (a *ADX) Calculate(highs, lows, closes []float64) ([]float64, []float64, []float64, error) {
	if len(highs) != len(lows) || len(lows) != len(closes) {
		return nil, nil, nil, errors.New("highs, lows, and closes must have the same length")
	}
	if len(highs) < a.Window {
		return nil, nil, nil, errors.New("not enough data for ADX")
	}

	length := len(highs)

	// True Range, +DM, -DM
	tr := make([]float64, length)
	pDM := make([]float64, length)
	mDM := make([]float64, length)

	// Initialize first day
	tr[0] = highs[0] - lows[0]
	pDM[0] = 0
	mDM[0] = 0

	for i := 1; i < length; i++ {
		currentHigh := highs[i]
		currentLow := lows[i]
		prevClose := closes[i-1]

		// True Range
		range1 := currentHigh - currentLow
		range2 := math.Abs(currentHigh - prevClose)
		range3 := math.Abs(currentLow - prevClose)
		tr[i] = max(range1, max(range2, range3)) // uses max from utils.go

		// +DM and -DM
		upMove := currentHigh - highs[i-1]
		downMove := lows[i-1] - currentLow

		if upMove > downMove && upMove > 0 {
			pDM[i] = upMove
		} else {
			pDM[i] = 0
		}
		if downMove > upMove && downMove > 0 {
			mDM[i] = downMove
		} else {
			mDM[i] = 0
		}
	}

	// Smoothed TR, +DM, -DM via Wilder's smoothing
	smTr := make([]float64, length)
	smPDM := make([]float64, length)
	smMDM := make([]float64, length)

	var sumTR, sumPDM, sumMDM float64
	for i := 0; i < a.Window; i++ {
		sumTR += tr[i]
		sumPDM += pDM[i]
		sumMDM += mDM[i]
	}
	smTr[a.Window-1] = sumTR
	smPDM[a.Window-1] = sumPDM
	smMDM[a.Window-1] = sumMDM

	for i := a.Window; i < length; i++ {
		smTr[i] = smTr[i-1] - (smTr[i-1] / float64(a.Window)) + tr[i]
		smPDM[i] = smPDM[i-1] - (smPDM[i-1] / float64(a.Window)) + pDM[i]
		smMDM[i] = smMDM[i-1] - (smMDM[i-1] / float64(a.Window)) + mDM[i]
	}

	// Compute +DI, -DI
	plusDI := make([]float64, length)
	minusDI := make([]float64, length)
	for i := 0; i < a.Window-1; i++ {
		plusDI[i] = 0
		minusDI[i] = 0
	}
	for i := a.Window - 1; i < length; i++ {
		if smTr[i] == 0 {
			plusDI[i] = 0
			minusDI[i] = 0
		} else {
			plusDI[i] = (smPDM[i] / smTr[i]) * 100
			minusDI[i] = (smMDM[i] / smTr[i]) * 100
		}
	}

	// Compute DX and then ADX
	dx := make([]float64, length)
	for i := 0; i < a.Window-1; i++ {
		dx[i] = 0
	}
	for i := a.Window - 1; i < length; i++ {
		sumDI := plusDI[i] + minusDI[i]
		diffDI := math.Abs(plusDI[i] - minusDI[i])
		if sumDI == 0 {
			dx[i] = 0
		} else {
			dx[i] = (diffDI / sumDI) * 100
		}
	}

	adx := make([]float64, length)
	// first ADX at index (window - 1): average of DX over that window
	var sumDX float64
	for i := a.Window - 1 - (a.Window - 1); i <= a.Window-1; i++ {
		sumDX += dx[i]
	}
	adx[a.Window-1] = sumDX / float64(a.Window)

	for i := a.Window; i < length; i++ {
		adx[i] = ((adx[i-1] * float64(a.Window-1)) + dx[i]) / float64(a.Window)
	}

	return adx, plusDI, minusDI, nil
}
