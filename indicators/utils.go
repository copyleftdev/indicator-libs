// indicators/utils.go
package indicators

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
