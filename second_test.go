package chance

import (
	"testing"
)

func TestSecond(t *testing.T) {
	c := NewChance()
	actual := c.Second()
	if actual < 0 || actual > 59 {
		t.Errorf("Second() was incorrect, expect: [between 0 and 59], actual: %d.", actual)
	}
}
