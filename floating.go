package chance

import (
	"errors"
	"math"
)

// Generate a random float32 number between MinInt32 (-2,147,483,648) and MaxInt32 (2,147,483,647)
func (c *Chance) Floating() float32 {
	output, _ := c.FloatingWithParams(math.MinInt32, math.MaxInt32)
	return output
}

func (c *Chance) FloatingWithParams(min int, max int) (float32, error) {
	if min > max {
		return 0, errors.New("Min must be smaller than Max.")
	}

	output := float32(min) + c.Rand.Float32()*float32(max-min)
	return output, nil
}
