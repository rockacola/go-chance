package chance

// Generate a random gender.
func (c *Chance) Gender() string {
	if c.Bool() {
		return "male"
	} else {
		return "female"
	}
}

func (c *Chance) GenderWithParams(extraGenders []string) (string, error) {
	genders := append(extraGenders, "male", "female")
	randomIndex := c.Rand.Intn(len(genders) - 1)
	output := genders[randomIndex]
	return output, nil
}
