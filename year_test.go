package chance

import (
	"testing"
)

func TestYear(t *testing.T) {
	c := NewChance()
	actual := c.Year()
	if actual < 0 {
		t.Errorf("Year() was incorrect, expect: [positive number], actual: %d.", actual)
	}
}

func TestYearWithParams(t *testing.T) {
	c := NewChance()
	min, max := 1000, 2000
	actual, err := c.YearWithParams(min, max)
	if err != nil {
		t.Errorf("YearWithParams() execution error: %s", err.Error())
	} else if actual < min || actual > max {
		t.Errorf("YearWithParams() was incorrect, expect: [between %d and %d], actual: %d.", min, max, actual)
	}
}
