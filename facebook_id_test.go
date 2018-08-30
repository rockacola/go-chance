package chance

import (
	"testing"
)

func TestFacebookId(t *testing.T) {
	c := NewChance()
	actual := c.FacebookId()
	if len(actual) != 15 {
		t.Errorf("FacebookId() was incorrect, expect: [exactly 15 characters], actual: %s.", actual)
	}
}
