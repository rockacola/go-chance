package chance

import (
	"testing"
)

func TestAppleToken(t *testing.T) {
	c := NewChance()
	actual := c.AppleToken()
	if len(actual) != 64 {
		t.Errorf("AppleToken() was incorrect, expect: [64 characters], actual: %s.", actual)
	}
}
