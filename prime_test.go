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
	actual, err := c.PrimeWithParams(100, 200)
	if err != nil {
		t.Errorf("PrimeWithParams() execution error: %s", err.Error())
	} else if actual < 100 || actual > 200 {
		t.Errorf("PrimeWithParams() was incorrect, expect: [between 100 and 200], actual: %d.", actual)
	}
}
