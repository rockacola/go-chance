package chance

import (
	"errors"
	"math"
)

// Generate a random float64 number between MinInt32 (-2,147,483,648) and MaxInt32 (2,147,483,647).
func (c *Chance) Floating() float64 {
	output, _ := c.FloatingWithParams(math.MinInt32, math.MaxInt32)
	return output
}

func (c *Chance) FloatingWithParams(min int, max int) (float64, error) {
	if min > max {
		return 0, errors.New("Min must be smaller than Max.")
	}

	if min == max {
		return float64(min), nil
	}

	output := float64(min) + c.Rand.Float64()*float64(max-min)
	return output, nil
}
