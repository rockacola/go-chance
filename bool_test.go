package chance

import (
	"testing"
)

func TestBool(t *testing.T) {
	c := NewChance()
	var _ = c.Bool()
}
