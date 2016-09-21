package add

// Add a and b. len(b) must be >= len(a)
func Add(a, b []float64) []float64 {
	c := make([]float64, len(a))
	for i := range a {
		c[i] = a[i] + b[i]
	}
	return c
}
