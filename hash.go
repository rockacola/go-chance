package chance

import (
	"errors"
	"strings"

	"github.com/rockacola/go-chance/data"
)

// Generate a random hex hash.
func (c *Chance) Hash() string {
	output, _ := c.HashWithParams(40, false)
	return output
}

func (c *Chance) HashWithParams(length int, toUpperCase bool) (string, error) {
	if length <= 0 {
		return "", errors.New("length cannot be less than zero.")
	}

	output := c.stringFromPool(length, data.Characters["hex"])
	if toUpperCase {
		output = strings.ToUpper(output)
	}
	return output, nil
}
