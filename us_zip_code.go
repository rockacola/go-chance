package chance

import (
	"github.com/rockacola/go-chance/data"
)

// Generate a random US zip code.
func (c *Chance) UsZipCode() string {
	output, _ := c.UsZipCodeWithParams(false)
	return output
}

func (c *Chance) UsZipCodeWithParams(plusFour bool) (string, error) {
	zipCode := c.stringFromPool(5, data.Characters["numeric"])
	if plusFour {
		zipCode = zipCode + "-" + c.stringFromPool(4, data.Characters["numeric"])
	}
	return zipCode, nil
}
