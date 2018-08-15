package chance

import (
	"testing"
)

func TestFirstName(t *testing.T) {
	c := NewChance()
	actual := c.FirstName()
	if len(actual) <= 0 {
		t.Errorf("FirstName() was incorrect, expect: [1 or more characters], actual: %s.", actual)
	}
}

func TestFirstNamelWithParams(t *testing.T) {
	c := NewChance()
	gender := "male"
	nationality := "it"
	actual, err := c.FirstNameWithParams(gender, nationality)
	if err != nil {
		t.Errorf("FirstNameWithParams() execution error: %s", err.Error())
	} else if len(actual) <= 0 {
		t.Errorf("FirstNameWithParams() was incorrect, expect: [1 or more characters], actual: %s.", actual)
	}
}
