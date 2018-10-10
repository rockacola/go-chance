package chance

// Generate a random second value.
func (c *Chance) Second() int {
	output, _ := c.NaturalWithParams(0, 59)
	return output
}
