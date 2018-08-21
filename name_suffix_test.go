package chance

import (
	"testing"
)

func TestNameSuffix(t *testing.T) {
	c := NewChance()
	actual := c.NameSuffix()
	if len(actual) <= 1 {
		t.Errorf("NameSuffix() was incorrect, expect: [1 or more characters], actual: %s.", actual)
	}
}

func TestNameSuffixWithParams(t *testing.T) {
	c := NewChance()
	full := true
	actual, err := c.NameSuffixWithParams(full)
	if err != nil {
		t.Errorf("NameSuffixWithParams() execution error: %s", err.Error())
	} else if len(actual) <= 1 {
		t.Errorf("NameSuffixWithParams() was incorrect, expect: [1 or more characters], actual: %s.", actual)
	}
}
