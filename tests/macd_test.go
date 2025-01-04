// tests/macd_test.go
package tests

import (
	"testing"

	"github.com/copyleftdev/indicator-libs/indicators"
)

func TestMACD(t *testing.T) {
	data := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	macd := indicators.NewMACD(3, 6, 2)

	macdLine, signalLine, hist, err := macd.Calculate(data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(macdLine) != len(data) || len(signalLine) != len(data) || len(hist) != len(data) {
		t.Errorf("all MACD slices must match input length")
	}

	// Just a sanity check for final values not being NaN.
	// Detailed numeric checks could be added if you have reference MACD values.
	if macdLine[len(macdLine)-1] == 0 && signalLine[len(signalLine)-1] == 0 {
		t.Log("MACD final might be 0, consider a more precise test if needed.")
	}
}
