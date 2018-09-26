package chance

import (
	"testing"
)

func TestAltitude(t *testing.T) {
	c := NewChance()
	actual := c.Altitude()
	if actual < 0 {
		t.Errorf("Altitude() was incorrect, expect: [a float value greater or equal to 0], actual: %6.2f.", actual)
	}
}
