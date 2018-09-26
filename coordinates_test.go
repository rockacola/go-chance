package chance

import (
	"testing"
)

func TestCoordinates(t *testing.T) {
	c := NewChance()
	actual := c.Coordinates()
	if len(actual) < 10 {
		t.Errorf("Coordinates() was incorrect, expect: [more than 10 characters], actual: %s.", actual)
	}
}
