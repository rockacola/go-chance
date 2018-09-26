package chance

// Generate a random latitude, in meters.
func (c *Chance) Latitude() float64 {
	output, _ := c.FloatingWithParams(-90, 90)
	return output
}
