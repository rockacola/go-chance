package chance

import (
	"testing"
)

func TestWp7Anid(t *testing.T) {
	c := NewChance()
	actual := c.Wp7Anid()
	if len(actual) != 44 {
		t.Errorf("Wp7Anid() was incorrect, expect: [44 characters], actual: %s.", actual)
	}
}
