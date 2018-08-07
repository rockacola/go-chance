package chance

import (
	"testing"
)

func TestInteger(t *testing.T) {
	c := NewChance()
	_ = c.Integer()
}

func TestIntegerWithParams(t *testing.T) {
	c := NewChance()
	min, max := -100, 100
	actual, err := c.IntegerWithParams(min, max)
	if err != nil {
		t.Errorf("IntegerWithParams() execution error: %s", err.Error())
	} else if actual < min || actual > max {
		t.Errorf("IntegerWithParams() was incorrect, expect: [a float value between %d and %d], actual: %d.", min, max, actual)
	}
}
