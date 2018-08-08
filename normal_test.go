package chance

import (
	"testing"
)

func TestNormal(t *testing.T) {
	c := NewChance()
	_ = c.Normal()
}

func TestNormalWithParams(t *testing.T) {
	c := NewChance()
	mean, dev := 100, 15
	_, err := c.NormalWithParams(mean, dev)
	if err != nil {
		t.Errorf("PrimeWithParams() execution error: %s", err.Error())
	}
}
