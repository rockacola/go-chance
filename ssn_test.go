package chance

import (
	"testing"
)

func TestSsn(t *testing.T) {
	c := NewChance()
	actual := c.Ssn()
	if len(actual) < 4 {
		t.Errorf("Ssn() was incorrect, expect: [4 or more characters], actual: %s.", actual)
	}
}

func TestSsnWithParams(t *testing.T) {
	c := NewChance()
	actual, err := c.SsnWithParams(true, true)
	if err != nil {
		t.Errorf("SsnWithParams() execution error: %s", err.Error())
	} else if len(actual) < 4 {
		t.Errorf("SsnWithParams() was incorrect, expect: [4 or more characters], actual: %s.", actual)
	}
}
