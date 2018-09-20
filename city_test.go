package chance

import (
	"testing"
)

func TestCity(t *testing.T) {
	c := NewChance()
	actual := c.City()
	if len(actual) < 2 {
		t.Errorf("City() was incorrect, expect: [2 or more characters], actual: %s.", actual)
	}
}
