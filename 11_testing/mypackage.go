package mypackage

func Average(values []float64) float64 {
	total := 0.0
	for _, v := range values {
		total += float64(v)
	}
	return total / float64(len(values))
}
