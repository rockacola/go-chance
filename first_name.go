package chance

import (
	"errors"
	"reflect"

	"github.com/rockacola/go-chance/data"
)

// Generate a random first name.
func (c *Chance) FirstName() string {
	genders := reflect.ValueOf(data.FirstNames).MapKeys()
	gender := genders[0].Interface().(string) // Golang natively randomize its key orders, hence simply picking the first item is sufficient.
	nationalities := reflect.ValueOf(data.FirstNames[gender]).MapKeys()
	nationality := nationalities[0].Interface().(string)
	randomIndex := c.Rand.Intn(len(data.FirstNames[gender][nationality]) - 1)
	return data.FirstNames[gender][nationality][randomIndex]
}

func (c *Chance) FirstNameWithParams(gender string, nationality string) (string, error) {
	if gender == "all" {
		gender = c.Gender()
	}

	namesOfGender := data.FirstNames[gender]
	if len(namesOfGender) <= 0 {
		return "", errors.New("Invalid name gender.")
	}
	namesOfGenderNationality := data.FirstNames[gender][nationality]
	if len(namesOfGenderNationality) <= 0 {
		return "", errors.New("Invalid name nationality.")
	}

	randomIndex := c.Rand.Intn(len(namesOfGenderNationality) - 1)
	output := namesOfGenderNationality[randomIndex]
	return output, nil
}
