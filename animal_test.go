package chance

import (
	"testing"
)

func TestAnimal(t *testing.T) {
	c := NewChance()
	actual := c.Animal()
	if len(actual) <= 0 {
		t.Errorf("Animal() was incorrect, expect: [1 or more character], actual: %s.", actual)
	}
}

func BenchmarkAnimal(b *testing.B) {
	c := NewChance()
	for n := 0; n < b.N; n++ {
		c.Animal()
	}
}
