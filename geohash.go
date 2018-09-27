package chance

// Generate a random geohash.
func (c *Chance) Geohash() string {
	output, _ := c.GeohashWithParams(7)
	return output
}

func (c *Chance) GeohashWithParams(length int) (string, error) {
	output := c.stringFromPool(length, "0123456789bcdefghjkmnpqrstuvwxyz")
	return output, nil
}
