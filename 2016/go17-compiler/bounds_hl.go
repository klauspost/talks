1	package add
2
3	// Add a and b. len(b) must be >= len(a)
4	func Add(a, b []float64) []float64 {
5		c := make([]float64, len(a))
6		for i := range a {
7			c[i] = a[i] + b[i] // HL
8		}
9		return c
10	}
