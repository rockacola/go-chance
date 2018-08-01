package chance

import (
	"testing"

	"../../pkg/chance"
)

func TestBool(t *testing.T) {
	c := chance.NewChance()
	var _ = c.Bool()
}
func BenchmarkBool(b *testing.B) {
	c := chance.NewChance()
	for n := 0; n < b.N; n++ {
		c.Bool()
	}
}
