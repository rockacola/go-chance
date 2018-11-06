package chance

import (
	"testing"
)

func TestMillisecond(t *testing.T) {
	c := NewChance()
	actual := c.Millisecond()
	if actual < 0 || actual > 999 {
		t.Errorf("Millisecond() was incorrect, expect: [between 0 and 999], actual: %d.", actual)
	}
}
