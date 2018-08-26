package chance

import (
	"strings"
)

// Generate a random name.
func (c *Chance) Name() string {
	output, _ := c.NameWithParams(false, false, false, "all", "en")
	return output
}

func (c *Chance) NameWithParams(middleName bool, prefix bool, suffix bool, gender string, nationality string) (string, error) {
	nameFragments := make([]string, 0)

	// Prefix
	if prefix {
		prefixValue, err := c.NamePrefixWithParams(false, gender)
		if err != nil {
			return "", err
		} else {
			nameFragments = append(nameFragments, prefixValue)
		}
	}

	// First name
	firstName, err := c.FirstNameWithParams(gender, nationality)
	if err != nil {
		return "", err
	} else {
		nameFragments = append(nameFragments, firstName)
	}

	// Middle Name
	if middleName {
		middleNameValue, err := c.FirstNameWithParams(gender, nationality)
		if err != nil {
			return "", err
		} else {
			nameFragments = append(nameFragments, middleNameValue)
		}
	}

	// Last name
	lastName, err := c.LastNameWithParams(nationality)
	if err != nil {
		return "", err
	} else {
		nameFragments = append(nameFragments, lastName)
	}

	// Suffix
	if suffix {
		suffixValue, err := c.NameSuffixWithParams(false)
		if err != nil {
			return "", err
		} else {
			nameFragments = append(nameFragments, suffixValue)
		}
	}

	output := strings.Join(nameFragments, " ")
	return output, nil
}
