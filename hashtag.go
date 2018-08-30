package chance

// Generate a random hashtag.
func (c *Chance) Hashtag() string {
	return "#" + c.Word()
}
