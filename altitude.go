package chance

// Generate a random altitude, in meters.
func (c *Chance) Altitude() float64 {
	output, _ := c.FloatingWithParams(0, 8848)
	return output
}
