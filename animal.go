package chance

import (
	"errors"
	"reflect"

	"github.com/rockacola/go-chance/data"
)

// Generate a random animal name within a pool of available animals.
func (c *Chance) Animal() string {
	categories := reflect.ValueOf(data.Animals).MapKeys()
	category := categories[0].Interface().(string) // Golang natively randomize its key orders, hence simply picking the first item is sufficient.
	randomIndex := c.Rand.Intn(len(data.Animals[category]))
	return data.Animals[category][randomIndex]
}

func (c *Chance) AnimalWithParams(category string) (string, error) {
	animalsOfCategory := data.Animals[category]
	if len(animalsOfCategory) <= 0 {
		return "", errors.New("Invalid animal category.")
	}

	randomIndex := c.Rand.Intn(len(animalsOfCategory))
	output := animalsOfCategory[randomIndex]
	return output, nil
}
