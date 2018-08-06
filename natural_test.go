package chance

import (
	"testing"
)

func TestNatural(t *testing.T) {
	c := NewChance()
	actual := c.Natural()
	if actual < 0 {
		t.Errorf("Natural() was incorrect, expect: [positive number], actual: %d.", actual)
	}
}
