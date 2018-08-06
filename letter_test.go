package chance

import (
	"testing"
)

func TestLetter(t *testing.T) {
	c := NewChance()
	actual := c.Letter()
	if len(actual) != 1 {
		t.Errorf("Letter() was incorrect, expect: [exactly 1 character], actual: %s.", actual)
	}
}

func TestLetterrWithParams(t *testing.T) {
	c := NewChance()
	actual, err := c.LetterWithParams(true, true)
	if err != nil {
		t.Errorf("LetterWithParams() execution error: %s", err.Error())
	} else if len(actual) != 1 {
		t.Errorf("LetterWithParams() was incorrect, expect: [exactly 1 character], actual: %s.", actual)
	}
}
