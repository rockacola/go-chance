package chance

import (
	"testing"
)

func TestKlout(t *testing.T) {
	c := NewChance()
	actual := c.Klout()
	if actual < 0 {
		t.Errorf("Klout() was incorrect, expect: [positive number], actual: %d.", actual)
	}
}
