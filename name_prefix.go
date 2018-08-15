package chance

import (
	"errors"
)

// Generate a random name prefix.
func (c *Chance) NamePrefix() string {
	full := false // Always show abbreviated prefix
	gender, _ := c.GenderWithParams([]string{"all"})
	output, _ := c.NamePrefixWithParams(full, gender)
	return output
}

func (c *Chance) NamePrefixWithParams(full bool, gender string) (string, error) {
	// Build pool
	pool := []map[string]string{
		{"name": "Doctor", "abbreviation": "Dr."},
	}
	if gender == "male" {
		malePrefixes := []map[string]string{
			{"name": "Mister", "abbreviation": "Mr."},
		}
		pool = append(pool, malePrefixes...)
	} else if gender == "female" {
		femalePrefixes := []map[string]string{
			{"name": "Miss", "abbreviation": "Miss"},
			{"name": "Misses", "abbreviation": "Mrs."},
		}
		pool = append(pool, femalePrefixes...)
	} else if gender == "all" {
		// Do nothing
	} else {
		return "", errors.New("Invalid gender.")
	}

	randomIndex := c.Rand.Intn(len(pool))
	prefixObj := pool[randomIndex]

	var output string
	if full {
		output = prefixObj["name"]
	} else {
		output = prefixObj["abbreviation"]
	}

	return output, nil
}
