package chance

import (
	"testing"
)

func TestSentence(t *testing.T) {
	c := NewChance()
	actual := c.Sentence()
	if len(actual) < 4 {
		t.Errorf("Sentenced() was incorrect, expect: [4 or more characters], actual: %s.", actual)
	}
}

func TestSentenceWithParams(t *testing.T) {
	c := NewChance()
	min := 2
	max := 5
	actual, err := c.SentenceWithParams(min, max)
	if err != nil {
		t.Errorf("SentenceWithParams() execution error: %s", err.Error())
	} else if len(actual) < 4 {
		t.Errorf("SentenceWithParams() was incorrect, expect: [4 or more characters], actual: %s.", actual)
	}
}
