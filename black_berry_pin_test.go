package chance

import (
	"testing"
)

func TestBlackBerryPin(t *testing.T) {
	c := NewChance()
	actual := c.BlackBerryPin()
	if len(actual) != 8 {
		t.Errorf("BlackBerryPin() was incorrect, expect: [8 characters], actual: %s.", actual)
	}
}
