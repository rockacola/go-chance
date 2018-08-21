package chance

// Generate a random name suffix.
func (c *Chance) NameSuffix() string {
	full := false // Always show abbreviated prefix
	output, _ := c.NameSuffixWithParams(full)
	return output
}

func (c *Chance) NameSuffixWithParams(full bool) (string, error) {
	// Build pool
	pool := []map[string]string{
		{"name": "Doctor of Osteopathic Medicine", "abbreviation": "D.O."},
		{"name": "Doctor of Philosophy", "abbreviation": "Ph.D."},
		{"name": "Esquire", "abbreviation": "Esq."},
		{"name": "Junior", "abbreviation": "Jr."},
		{"name": "Juris Doctor", "abbreviation": "J.D."},
		{"name": "Master of Arts", "abbreviation": "M.A."},
		{"name": "Master of Business Administration", "abbreviation": "M.B.A."},
		{"name": "Master of Science", "abbreviation": "M.S."},
		{"name": "Medical Doctor", "abbreviation": "M.D."},
		{"name": "Senior", "abbreviation": "Sr."},
		{"name": "The Third", "abbreviation": "III"},
		{"name": "The Fourth", "abbreviation": "IV"},
		{"name": "Bachelor of Engineering", "abbreviation": "B.E"},
		{"name": "Bachelor of Technology", "abbreviation": "B.TECH"},
	}

	randomIndex := c.Rand.Intn(len(pool))
	suffixObj := pool[randomIndex]

	var output string
	if full {
		output = suffixObj["name"]
	} else {
		output = suffixObj["abbreviation"]
	}

	return output, nil
}
