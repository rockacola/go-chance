package chance

import "math"

// Return a normally-distributed random variate.
func (c *Chance) Normal() float64 {
	output, _ := c.NormalWithParams(100, 15)
	return output
}

func (c *Chance) NormalWithParams(mean int, deviation int) (float64, error) {
	// The Marsaglia Polar method
	var s float64 = 1
	var u float64
	var v float64

	for s >= 1 {
		// U and V are from the uniform distribution on (-1, 1)
		u = c.Rand.Float64()*2 - 1
		v = c.Rand.Float64()*2 - 1
		s = u*u + v*v
	}

	// Compute the standard normal variate
	norm := u * math.Sqrt(-2*math.Log(s)/s)

	// Shape and scale
	output := float64(deviation)*norm + float64(mean)
	return output, nil
}
