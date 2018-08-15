package chance

import (
	"testing"
)

func TestAge(t *testing.T) {
	c := NewChance()
	actual := c.Age()
	if actual < 0 || actual > 100 {
		t.Errorf("Age() was incorrect, expect: [a number between 0 and 100], actual: %d.", actual)
	}
}

func TestAgeWithParams(t *testing.T) {
	c := NewChance()
	category := "all"
	actual, err := c.AgeWithParams(category)
	if err != nil {
		t.Errorf("AgeWithParams() execution error: %s", err.Error())
	} else if actual < 0 || actual > 100 {
		t.Errorf("AgeWithParams() was incorrect, expect: [a number between 0 and 100], actual: %d.", actual)
	}
}
