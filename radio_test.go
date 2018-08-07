package chance

import (
	"testing"
)

func TestRadio(t *testing.T) {
	c := NewChance()
	actual := c.Radio()
	if len(actual) != 4 {
		t.Errorf("Radio() was incorrect, expect: [exactly 4 characters], actual: %s.", actual)
	}
}

func TestRadioWithParams(t *testing.T) {
	c := NewChance()
	side := "east"
	actual, err := c.RadioWithParams(side)
	if err != nil {
		t.Errorf("RadioWithParams() execution error: %s", err.Error())
	} else if len(actual) != 4 {
		t.Errorf("RadioWithParams() was incorrect, expect: [exactly 4 characters], actual: %s.", actual)
	}
}
