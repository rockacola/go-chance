package chance

import (
	"testing"
)

func TestTwitter(t *testing.T) {
	c := NewChance()
	actual := c.Twitter()
	if len(actual) < 2 {
		t.Errorf("Twitter() was incorrect, expect: [2 or more characters], actual: %s.", actual)
	}
}
