package chance

import (
	"testing"
)

func TestPrime(t *testing.T) {
	c := NewChance()
	actual := c.Prime()
	if actual < 0 {
		t.Errorf("Prime() was incorrect, expect: [positive number], actual: %d.", actual)
	}
}

func TestPrimeWithParams(t *testing.T) {
	c := NewChance()
	min, max := 100, 200
	actual, err := c.PrimeWithParams(min, max)
	if err != nil {
		t.Errorf("PrimeWithParams() execution error: %s", err.Error())
	} else if actual < min || actual > max {
		t.Errorf("PrimeWithParams() was incorrect, expect: [between %d and %d], actual: %d.", min, max, actual)
	}
}
