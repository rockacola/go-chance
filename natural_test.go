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
	actual, err := c.NaturalWithParams(100, 200)
	if err != nil {
		t.Errorf("NaturalWithParams() execution error: %s", err.Error())
	} else if actual < 100 || actual > 200 {
		t.Errorf("NaturalWithParams() was incorrect, expect: [between 0 and 10], actual: %d.", actual)
	}
}
