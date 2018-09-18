package chance

import "errors"

// Generate a random domain name with TLD.
func (c *Chance) Domain() string {
	return c.Word() + "." + c.Tld()
}

func (c *Chance) DomainWithParams(tld string) (string, error) {
	if len(tld) < 1 {
		return "", errors.New("tld cannot be less than 1 character.")
	}

	output := c.Word() + "." + tld
	return output, nil
}
