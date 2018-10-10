package chance

import "errors"

// Generate a random year value.
func (c *Chance) Year() int {
	output, _ := c.YearWithParams(1900, 2100)
	return output
}

func (c *Chance) YearWithParams(min int, max int) (int, error) {
	if min < 0 {
		return 0, errors.New("min cannot be less than zero.")
	} else if max < min {
		return 0, errors.New("max must be smaller or equals to min.")
	}

	output, _ := c.NaturalWithParams(min, max)
	return output, nil
}
