package chance

import (
	"errors"
)

// Generate a random boolean value, either 'true' or 'false'
func (c *Chance) Bool() bool {
	return c.Rand.Intn(2) == 1
}

func (c *Chance) BoolWithParams(likelihood int) (bool, error) {
	if likelihood < 0 || likelihood > 100 {
		return false, errors.New("Likelihood musts be between 0 and 100.")
	}
	n, _ := c.NaturalWithParams(0, 99)
	output := n < likelihood
	return output, nil
}
