package chance

import (
	"testing"
)

func TestAmPm(t *testing.T) {
	c := NewChance()
	var actual = c.AmPm()
	if actual != "am" && actual != "pm" {
		t.Errorf("AmPm() was incorrect, expect: [string 'am' or 'pm'], actual: %s.", actual)
	}
}
