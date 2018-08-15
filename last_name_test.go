package chance

import (
	"testing"
)

func TestLastName(t *testing.T) {
	c := NewChance()
	actual := c.LastName()
	if len(actual) <= 1 {
		t.Errorf("LastName() was incorrect, expect: [1 or more characters], actual: %s.", actual)
	}
}

func TestLastNameWithParams(t *testing.T) {
	c := NewChance()
	nationality := "it"
	actual, err := c.LastNameWithParams(nationality)
	if err != nil {
		t.Errorf("LastNameWithParams() execution error: %s", err.Error())
	} else if len(actual) <= 1 {
		t.Errorf("LastNameWithParams() was incorrect, expect: [1 or more characters], actual: %s.", actual)
	}
}
