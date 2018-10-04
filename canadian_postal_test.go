package chance

import (
	"testing"
)

func TestCanadianPostal(t *testing.T) {
	c := NewChance()
	actual := c.CanadianPostal()
	if len(actual) != 7 {
		t.Errorf("CanadianPostal() was incorrect, expect: [exactly 7 characters], actual: %s.", actual)
	}
}
