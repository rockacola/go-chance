package chance

// Generate a random longitude, in meters.
func (c *Chance) Longitude() float64 {
	output, _ := c.FloatingWithParams(-180, 180)
	return output
}
