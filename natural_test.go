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

func TestNaturalWithParams(t *testing.T) {
	c := NewChance()
	min, max := 100, 200
	actual, err := c.NaturalWithParams(min, max)
	if err != nil {
		t.Errorf("NaturalWithParams() execution error: %s", err.Error())
	} else if actual < min || actual > max {
		t.Errorf("NaturalWithParams() was incorrect, expect: [between %d and %d], actual: %d.", min, max, actual)
	}
}
