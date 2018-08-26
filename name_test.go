package chance

import (
	"testing"
)

func TestName(t *testing.T) {
	c := NewChance()
	actual := c.Name()
	if len(actual) <= 1 {
		t.Errorf("Name() was incorrect, expect: [1 or more characters], actual: %s.", actual)
	}
}

func TestNameWithParams(t *testing.T) {
	c := NewChance()
	actual, err := c.NameWithParams(false, false, false, "all", "en")
	if err != nil {
		t.Errorf("NameWithParams() execution error: %s", err.Error())
	} else if len(actual) <= 1 {
		t.Errorf("NameWithParams() was incorrect, expect: [1 or more characters], actual: %s.", actual)
	}
}
