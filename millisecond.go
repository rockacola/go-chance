package chance

// Generate a random millisecond value.
func (c *Chance) Millisecond() int {
	output, _ := c.NaturalWithParams(0, 999)
	return output
}
