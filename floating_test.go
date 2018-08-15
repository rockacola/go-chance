package chance

import (
	"testing"
)

func TestFloating(t *testing.T) {
	c := NewChance()
	_ = c.Floating()
}

func TestFloatingWithParams(t *testing.T) {
	c := NewChance()
	min, max := -100, 100
	actual, err := c.FloatingWithParams(min, max)
	if err != nil {
		t.Errorf("FloatingWithParams() execution error: %s", err.Error())
	} else if actual < float64(min) || actual > float64(max) {
		t.Errorf("FloatingWithParams() was incorrect, expect: [a float value between %d and %d], actual: %6.2f.", min, max, actual)
	}
}
