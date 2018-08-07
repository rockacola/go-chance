package chance

import "errors"

// Generate a semi-pronounceable nonsense word.
func (c *Chance) Word() string {
	output, _ := c.WordWithParams(2, 5)
	return output
}

func (c *Chance) WordWithParams(minSyllables int, maxSyllables int) (string, error) {
	if minSyllables < 1 {
		return "", errors.New("minSyllables cannot be less than 1.")
	} else if maxSyllables < 1 {
		return "", errors.New("maxSyllables cannot be less than 1.")
	} else if minSyllables > maxSyllables {
		return "", errors.New("minSyllables must be smaller than maxSyllables.")
	}

	syllableCount, _ := c.NaturalWithParams(minSyllables, maxSyllables)
	output := ""
	for i := 0; i < syllableCount; i++ {
		output += c.Syllable()
	}
	return output, nil
}
