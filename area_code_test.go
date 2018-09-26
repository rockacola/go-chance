package chance

import (
	"testing"
)

func TestAreaCode(t *testing.T) {
	c := NewChance()
	actual := c.AreaCode()
	if len(actual) < 2 {
		t.Errorf("AreaCode() was incorrect, expect: [3 or more characters], actual: %s.", actual)
	}
}

func TestAreaCodeWithParams(t *testing.T) {
	c := NewChance()
	parenthesis := false
	actual, err := c.AreaCodeWithParams(parenthesis)
	if err != nil {
		t.Errorf("AreaCodeWithParams() execution error: %s", err.Error())
	} else if len(actual) < 2 {
		t.Errorf("AreaCodeWithParams() was incorrect, expect: [3 or more characters], actual: %s.", actual)
	}
}
