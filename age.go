package chance

import "errors"

// Generate a random age.
func (c *Chance) Age() int {
	output, _ := c.AgeWithParams("all")
	return output
}

func (c *Chance) AgeWithParams(category string) (int, error) {
	if category == "child" {
		return c.NaturalWithParams(0, 12)
	} else if category == "teen" {
		return c.NaturalWithParams(13, 19)
	} else if category == "adult" {
		return c.NaturalWithParams(18, 65)
	} else if category == "senior" {
		return c.NaturalWithParams(65, 100)
	} else if category == "all" {
		return c.NaturalWithParams(0, 100)
	} else {
		return 0, errors.New("Invalid specified category.")
	}
}
