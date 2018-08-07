package chance

import (
	"errors"
	"math"
)

// Generate a random int32 number between MinInt32 (-2,147,483,648) and MaxInt32 (2,147,483,647)
func (c *Chance) Integer() int {
	output, _ := c.IntegerWithParams(math.MinInt32, math.MaxInt32)
	return output
}

func (c *Chance) IntegerWithParams(min int, max int) (int, error) {
	if min > max {
		return 0, errors.New("Min must be smaller than Max.")
	}

	if min == max {
		return min, nil
	}

	output := min + c.Rand.Intn(max-min)
	return output, nil
}
