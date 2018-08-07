package chance

import (
	"testing"
)

func TestString(t *testing.T) {
	c := NewChance()
	actual := c.String()
	if len(actual) != 5 {
		t.Errorf("String() was incorrect, expect: [exactly 5 characters], actual: %s.", actual)
	}
}

func TestStringWithParams(t *testing.T) {
	c := NewChance()
	length := 8
	actual, err := c.StringWithParams(length, true, true, true, true)
	if err != nil {
		t.Errorf("StringWithParams() execution error: %s", err.Error())
	} else if len(actual) != length {
		t.Errorf("StringWithParams() was incorrect, expect: [exactly %d characters], actual: %s.", length, actual)
	}
}
