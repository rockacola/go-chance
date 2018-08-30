package chance

import (
	"testing"
)

func TestTld(t *testing.T) {
	c := NewChance()
	actual := c.Tld()
	if len(actual) < 2 {
		t.Errorf("Tld() was incorrect, expect: [2 or more characters], actual: %s.", actual)
	}
}
