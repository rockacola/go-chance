package chance

import (
	"testing"
)

func TestWp8Anid2(t *testing.T) {
	c := NewChance()
	actual := c.Wp8Anid2()
	if len(actual) != 32 {
		t.Errorf("Wp8Anid2() was incorrect, expect: [32 characters], actual: %s.", actual)
	}
}
