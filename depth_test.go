package chance

import (
	"testing"
)

func TestDepth(t *testing.T) {
	c := NewChance()
	actual := c.Depth()
	if actual < -10994 || actual > 0 {
		t.Errorf("Depth() was incorrect, expect: [a float value between -10994 and 0], actual: %6.2f.", actual)
	}
}
