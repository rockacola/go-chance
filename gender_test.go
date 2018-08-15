package chance

import (
	"testing"
)

func TestGender(t *testing.T) {
	c := NewChance()
	actual := c.Gender()
	if len(actual) <= 1 {
		t.Errorf("Gender() was incorrect, expect: [1 or more characters], actual: %s.", actual)
	}
}

func TestGenderWithParams(t *testing.T) {
	c := NewChance()
	extraGenders := []string{"agender", "genderqueer", "trans"}
	actual, err := c.GenderWithParams(extraGenders)
	if err != nil {
		t.Errorf("GenderWithParams() execution error: %s", err.Error())
	} else if len(actual) <= 1 {
		t.Errorf("GenderWithParams() was incorrect, expect: [1 or more characters], actual: %s.", actual)
	}
}
