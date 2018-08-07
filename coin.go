package chance

// Flip a coin and returns 'head' or tail'.
func (c *Chance) Coin() string {
	if c.Bool() {
		return "head"
	} else {
		return "tail"
	}
}
