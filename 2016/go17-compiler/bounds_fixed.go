1	package add
2
3	// Add a and b. len(b) must be >= len(a)
4	func Add(a, b []float64) []float64 {
5		b = b[:len(a)] // HL
6		c := make([]float64, len(a))
7		c = c[:len(a)] // HL
8		for i := range a {
9			c[i] = a[i] + b[i]
10		}
11		return c
12	}
