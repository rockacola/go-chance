package chance

// Generate a random hour value.
func (c *Chance) Hour() int {
	output, _ := c.HourWithParams(true)
	return output
}

func (c *Chance) HourWithParams(twentyFour bool) (int, error) {
	min, max := 1, 12
	if twentyFour {
		min = 0
		max = 23
	}

	output, _ := c.NaturalWithParams(min, max)
	return output, nil
}
