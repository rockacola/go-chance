package chance

import (
	"errors"
	"strconv"

	"github.com/rockacola/go-chance/data"
)

// Generate a random Globally Unique Identifier
func (c *Chance) Guid() string {
	result, _ := c.GuidWithParams(5)
	return result
}

func (c *Chance) GuidWithParams(version int) (string, error) {
	if version < 1 || version > 5 {
		return "", errors.New("version must be between 1 and 5.")
	}

	guidPool := data.Characters["hex"]
	variantPool := data.Characters["guidVariant"]
	output := c.stringFromPool(8, guidPool) + "-" +
		c.stringFromPool(4, guidPool) + "-" +
		strconv.Itoa(version) + c.stringFromPool(3, guidPool) + "-" +
		c.stringFromPool(1, variantPool) + c.stringFromPool(3, guidPool) + "-" +
		c.stringFromPool(12, guidPool)
	return output, nil
}
