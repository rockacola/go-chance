package chance

import "errors"

// Generate a random email address.
func (c *Chance) Email() string {
	return c.Word() + "@" + c.Word() + "." + c.Tld()
}

func (c *Chance) EmailWithParams(domain string) (string, error) {
	if len(domain) < 1 {
		return "", errors.New("domain cannot be less than 1 character.")
	}

	output := c.Word() + "@" + domain
	return output, nil
}
