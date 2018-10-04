package chance

import (
	"testing"
)

func TestUsZipCode(t *testing.T) {
	c := NewChance()
	actual := c.UsZipCode()
	if len(actual) != 5 {
		t.Errorf("UsZipCode() was incorrect, expect: [exactly 5 characters], actual: %s.", actual)
	}
}

func TestUsZipCodeWithParams(t *testing.T) {
	c := NewChance()
	plusFour := true
	actual, err := c.UsZipCodeWithParams(plusFour)
	if err != nil {
		t.Errorf("UsZipCodeWithParams() execution error: %s", err.Error())
	} else if len(actual) != 10 {
		t.Errorf("UsZipCodeWithParams() was incorrect, expect: [exactly 10 characters], actual: %s.", actual)
	}
}
