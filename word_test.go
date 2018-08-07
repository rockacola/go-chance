package chance

import (
	"testing"
)

func TestWord(t *testing.T) {
	c := NewChance()
	actual := c.Word()
	if len(actual) < 4 {
		t.Errorf("Word() was incorrect, expect: [4 or more characters], actual: %s.", actual)
	}
}

func TestWordWithParams(t *testing.T) {
	c := NewChance()
	min := 2
	max := 5
	actual, err := c.WordWithParams(min, max)
	if err != nil {
		t.Errorf("WordWithParams() execution error: %s", err.Error())
	} else if len(actual) < 4 || len(actual) > 15 {
		t.Errorf("WordWithParams() was incorrect, expect: [better 4 and 15 characters], actual: %s.", actual)
	}
}
