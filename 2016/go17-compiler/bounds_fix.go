1	package add
2
3	// Add a and b. len(b) must be >= len(a)
4	func Add(a, b []byte) []float64 {
5		b = b[:len(a)] // HL
6		c := make([]float64, len(a))
7		for i := range a {
8			c[i] = a[i] + b[i]
9		}
10		return c
11	}
