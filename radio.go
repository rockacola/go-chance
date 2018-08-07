package chance

import (
	"errors"

	"github.com/rockacola/go-chance/data"
)

// Generate a random radio call sign.
func (c *Chance) Radio() string {
	return c.stringFromPool(1, "KW") + c.stringFromPool(3, data.Characters["upper"])
}

func (c *Chance) RadioWithParams(side string) (string, error) {
	if side != "west" && side != "east" {
		return "", errors.New("side value needs to be either 'west' or 'east'.")
	}

	prefix := ""
	if side == "west" {
		prefix = "K"
	} else {
		prefix = "W"
	}
	output := prefix + c.stringFromPool(3, data.Characters["upper"])
	return output, nil
}
