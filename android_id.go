package chance

// Generate an Android GCM Registration identifier.
func (c *Chance) AndroidId() string {
	return "APA91" + c.stringFromPool(178, "0123456789abcefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-_")
}
