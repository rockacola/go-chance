package chance

import (
	"bytes"
	"errors"

	"github.com/rockacola/go-chance/data"
)

// Generate a random string with a pool of common characters.
func (c *Chance) String() string {
	output := ""
	length := 5
	for i := 0; i < length; i++ {
		output += c.Character()
	}
	return output
}

func (c *Chance) StringWithParams(length int, lowerCaseAlphabets bool, upperCaseAlphabets bool, numerics bool, symbols bool) (string, error) {
	if length < 0 {
		return "", errors.New("Length cannot be less than zero.")
	}

	if length == 0 {
		return "", nil
	}

	// build pool
	var buffer bytes.Buffer
	if lowerCaseAlphabets {
		buffer.WriteString(data.Characters["lower"])
	}
	if upperCaseAlphabets {
		buffer.WriteString(data.Characters["upper"])
	}
	if numerics {
		buffer.WriteString(data.Characters["numberic"])
	}
	if symbols {
		buffer.WriteString(data.Characters["symbol"])
	}

	pool := buffer.String()
	if len(pool) <= 0 {
		return "", errors.New("Invalid pool size, you must have at least 1 option as [true].")
	}

	output := ""
	for i := 0; i < length; i++ {
		randomIndex := c.Rand.Intn(len(pool) - 1)
		output += string(pool[randomIndex])
	}
	return output, nil
}
