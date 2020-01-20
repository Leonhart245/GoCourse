package HW6.1

// Average is unchanged fom original author
func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

// Sum calculates sum of the slice elements
func Sum(xs []float64) (sum float64) {
	for _, x := range xs {
		sum += x
	}
	return sum
}