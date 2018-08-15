package chance

import (
	"errors"
	"reflect"

	"github.com/rockacola/go-chance/data"
)

// Generate a random last name.
func (c *Chance) LastName() string {
	nationalities := reflect.ValueOf(data.LastNames).MapKeys()
	nationality := nationalities[0].Interface().(string)
	randomIndex := c.Rand.Intn(len(data.LastNames[nationality]) - 1)
	return data.LastNames[nationality][randomIndex]
}

func (c *Chance) LastNameWithParams(nationality string) (string, error) {
	namesOfNationality := data.LastNames[nationality]
	if len(namesOfNationality) <= 0 {
		return "", errors.New("Invalid name nationality.")
	}

	randomIndex := c.Rand.Intn(len(namesOfNationality) - 1)
	output := namesOfNationality[randomIndex]
	return output, nil
}
