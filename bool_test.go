package chance

import (
	"testing"
)

func TestBool(t *testing.T) {
	c := NewChance()
	var _ = c.Bool()
}

func BenchmarkBool(b *testing.B) {
	c := NewChance()
	for n := 0; n < b.N; n++ {
		c.Bool()
	}
}
