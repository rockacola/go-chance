package chance

import (
	"testing"

	"../../pkg/chance"
)

func TestNatural(t *testing.T) {
	c := chance.NewChance()
	actual := c.Natural()
	if actual < 0 {
		t.Errorf("Natural() was incorrect, expect: [positive number], actual: %d.", actual)
	}
}
func BenchmarkNatural(b *testing.B) {
	c := chance.NewChance()
	for n := 0; n < b.N; n++ {
		c.Natural()
	}
}
