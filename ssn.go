package chance

import (
	"github.com/rockacola/go-chance/data"
)

// Generate a random social security number.
func (c *Chance) Ssn() string {
	output, _ := c.SsnWithParams(true, true)
	return output
}

func (c *Chance) SsnWithParams(ssnFour bool, dashes bool) (string, error) {
	separator := ""
	if dashes {
		separator = "-"
	}

	if ssnFour {
		output := c.stringFromPool(4, data.Characters["numeric"])
		return output, nil
	} else {
		output := c.stringFromPool(3, data.Characters["numeric"]) + separator +
			c.stringFromPool(2, data.Characters["numeric"]) + separator +
			c.stringFromPool(4, data.Characters["numeric"])
		return output, nil
	}
}
