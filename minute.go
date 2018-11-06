package chance

// Generate a random minute value.
func (c *Chance) Minute() int {
	output, _ := c.NaturalWithParams(0, 59)
	return output
}
