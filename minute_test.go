package chance

import (
	"testing"
)

func TestMinute(t *testing.T) {
	c := NewChance()
	actual := c.Minute()
	if actual < 0 || actual > 59 {
		t.Errorf("Minute() was incorrect, expect: [between 0 and 59], actual: %d.", actual)
	}
}
