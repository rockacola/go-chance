package chance

import (
	"bytes"
	"errors"
	"reflect"

	"github.com/rockacola/go-chance/data"
)

// Generate a single length string within a pool of common characters
func (c *Chance) Character() string {
	categories := reflect.ValueOf(data.Characters).MapKeys()
	category := categories[0].Interface().(string) // Golang natively randomize its key orders, hence simply picking the first item is sufficient.
	randomIndex := c.Rand.Intn(len(data.Characters[category]) - 1)
	return string(data.Characters[category][randomIndex])
}

func (c *Chance) CharacterWithParams(lowerCaseAlphabets bool, upperCaseAlphabets bool, numerics bool, symbols bool) (string, error) {
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

	randomIndex := c.Rand.Intn(len(pool) - 1)
	output := string(pool[randomIndex])
	return output, nil
}
