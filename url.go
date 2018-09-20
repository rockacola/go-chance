package chance

// Generate a random URL.
func (c *Chance) Url() string {
	protocol := "http"
	domain := c.Domain()
	path := c.Word()
	return protocol + "://" + domain + "/" + path
}
