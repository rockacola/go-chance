package chance

import (
	"testing"
)

func TestCoin(t *testing.T) {
	c := NewChance()
	var actual = c.Coin()
	if actual != "head" && actual != "tail" {
		t.Errorf("Coin() was incorrect, expect: [string 'head' or 'tail'], actual: %s.", actual)
	}
}
