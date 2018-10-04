package chance

// Return am or pm.
func (c *Chance) AmPm() string {
	if c.Bool() {
		return "am"
	} else {
		return "pm"
	}
}
