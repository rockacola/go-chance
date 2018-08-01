package chance

import (
	"math/rand"
	"reflect"

	"./data"
)

func (c *Chance) Animal() string {
	animalTypes := reflect.ValueOf(data.Animal).MapKeys()
	animalType := animalTypes[0].Interface().(string) // Golang natively randomize its key orders, hence simply picking the first item is sufficient.
	randomIndex := rand.Intn(len(data.Animal[animalType]) - 1)
	return data.Animal[animalType][randomIndex]
}
