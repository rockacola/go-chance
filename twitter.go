package chance

// Generate a random Twitter handler.
func (c *Chance) Twitter() string {
	return "@" + c.Word()
}
