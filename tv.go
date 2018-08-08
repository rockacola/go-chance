package chance

// Generate a TV station call sign. This is an alias for Radio() since they both follow the same rules.
func (c *Chance) Tv() string {
	return c.Radio()
}

func (c *Chance) TvWithParams(side string) (string, error) {
	return c.RadioWithParams(side)
}
