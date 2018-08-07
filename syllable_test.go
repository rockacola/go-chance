package chance

import (
	"testing"
)

func TestSyllable(t *testing.T) {
	c := NewChance()
	actual := c.Syllable()
	if len(actual) < 2 || len(actual) > 3 {
		t.Errorf("Syllable() was incorrect, expect: [string value between 2 and 3 characters], actual: %s.", actual)
	}
}
