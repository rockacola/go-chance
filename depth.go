package chance

// Generate a random depth, in meters. Depths are always negative.
func (c *Chance) Depth() float64 {
	output, _ := c.FloatingWithParams(-10994, 0)
	return output
}
