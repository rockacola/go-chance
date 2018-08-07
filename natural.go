package chance

import (
	"errors"
	"math"
)

// Generate a random int32 number between 0 and MaxInt32 (2,147,483,647).
func (c *Chance) Natural() int {
	output, _ := c.NaturalWithParams(0, math.MaxInt32)
	return output
}

func (c *Chance) NaturalWithParams(min int, max int) (int, error) {
	if min < 0 {
		return 0, errors.New("Min cannot be less than zero.")
	} else if min > max {
		return 0, errors.New("Min must be smaller than Max.")
	}

	if min == max {
		return min, nil
	}

	output := min + c.Rand.Intn(max-min)
	return output, nil
}
