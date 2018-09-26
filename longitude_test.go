package chance

import (
	"testing"
)

func TestLongitude(t *testing.T) {
	c := NewChance()
	actual := c.Longitude()
	if actual < -180 || actual > 180 {
		t.Errorf("Longitude() was incorrect, expect: [a float value between -180 and 180], actual: %6.2f.", actual)
	}
}
