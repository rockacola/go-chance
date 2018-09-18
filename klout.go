package chance

// Generate a random Klout score. Range 1-99.
func (c *Chance) Klout() int {
	output, _ := c.NaturalWithParams(1, 99)
	return output
}
