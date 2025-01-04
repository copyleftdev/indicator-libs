package indicators

type Indicator interface {
	Calculate([]float64) ([]float64, error)
}
