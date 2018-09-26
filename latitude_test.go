package chance

import (
	"testing"
)

func TestLatitude(t *testing.T) {
	c := NewChance()
	actual := c.Latitude()
	if actual < -90 || actual > 90 {
		t.Errorf("Latitude() was incorrect, expect: [a float value between 90 and 90], actual: %6.2f.", actual)
	}
}
