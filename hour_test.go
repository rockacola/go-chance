package chance

import (
	"testing"
)

func TestHour(t *testing.T) {
	c := NewChance()
	actual := c.Hour()
	if actual < 0 || actual > 23 {
		t.Errorf("Hour() was incorrect, expect: [between 0 and 23], actual: %d.", actual)
	}
}

func TestHourWithParams(t *testing.T) {
	c := NewChance()
	actual, err := c.HourWithParams(false)
	if err != nil {
		t.Errorf("HourWithParams() execution error: %s", err.Error())
	} else if actual < 1 || actual > 12 {
		t.Errorf("HourWithParams() was incorrect, expect: [between 1 and 12], actual: %d.", actual)
	}
}
