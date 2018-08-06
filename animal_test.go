package chance

import (
	"testing"
)

func TestAnimal(t *testing.T) {
	c := NewChance()
	actual := c.Animal()
	if len(actual) <= 0 {
		t.Errorf("Animal() was incorrect, expect: [1 or more character], actual: %s.", actual)
	}
}

func TestAnimalWithParams(t *testing.T) {
	c := NewChance()
	actual, err := c.AnimalWithParams("ocean")
	if err != nil {
		t.Errorf("AnimalWithParams() execution error: %s", err.Error())
	} else if len(actual) <= 0 {
		t.Errorf("AnimalWithParams() was incorrect, expect: [1 or more character], actual: %s.", actual)
	}
}
