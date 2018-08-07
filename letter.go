package chance

// Generate a single length string within a pool of alphabetical characters.
func (c *Chance) Letter() string {
	output, _ := c.CharacterWithParams(true, true, false, false)
	return output
}

func (c *Chance) LetterWithParams(lowerCaseAlphabets bool, upperCaseAlphabets bool) (string, error) {
	output, _ := c.CharacterWithParams(lowerCaseAlphabets, upperCaseAlphabets, false, false)
	return output, nil
}
