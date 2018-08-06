package chance

import (
	"errors"
	"reflect"

	"github.com/rockacola/go-chance/data"
)

func (c *Chance) Animal() string {
	categories := reflect.ValueOf(data.Animal).MapKeys()
	category := categories[0].Interface().(string) // Golang natively randomize its key orders, hence simply picking the first item is sufficient.
	randomIndex := c.Rand.Intn(len(data.Animal[category]) - 1)
	return data.Animal[category][randomIndex]
}

func (c *Chance) AnimalWithParams(category string) (string, error) {
	animalsOfCategory := data.Animal[category]
	if len(animalsOfCategory) <= 0 {
		return "", errors.New("Invalid animal category.")
	}

	randomIndex := c.Rand.Intn(len(animalsOfCategory) - 1)
	output := animalsOfCategory[randomIndex]
	return output, nil
}
