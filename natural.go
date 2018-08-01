package chance

import (
	"errors"
	"math"
)

func (c *Chance) Natural() int {
	output, _ := c.NaturalWithParams(0, math.MaxInt32)
	return output
}

func (c *Chance) NaturalWithParams(min int, max int) (int, error) {
	if min < 0 {
		return 0, errors.New("Min cannot be less than zero.")
	} else if max > math.MaxInt32 {
		return 0, errors.New("Max cannot be greater than MaxInt32.")
	}

	output := min + c.Rand.Intn(max-min)
	return output, nil
}
