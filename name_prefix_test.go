package chance

import (
	"testing"
)

func TestNamePrefix(t *testing.T) {
	c := NewChance()
	actual := c.NamePrefix()
	if len(actual) <= 1 {
		t.Errorf("NamePrefix() was incorrect, expect: [1 or more characters], actual: %s.", actual)
	}
}

func TestNamePrefixWithParams(t *testing.T) {
	c := NewChance()
	full := true
	gender := "all"
	actual, err := c.NamePrefixWithParams(full, gender)
	if err != nil {
		t.Errorf("NamePrefixWithParams() execution error: %s", err.Error())
	} else if len(actual) <= 1 {
		t.Errorf("NamePrefixWithParams() was incorrect, expect: [1 or more characters], actual: %s.", actual)
	}
}
