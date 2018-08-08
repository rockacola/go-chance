package chance

import (
	"testing"
)

func TestTv(t *testing.T) {
	c := NewChance()
	actual := c.Tv()
	if len(actual) != 4 {
		t.Errorf("Tv() was incorrect, expect: [exactly 4 characters], actual: %s.", actual)
	}
}

func TestTvWithParams(t *testing.T) {
	c := NewChance()
	side := "east"
	actual, err := c.TvWithParams(side)
	if err != nil {
		t.Errorf("TvWithParams() execution error: %s", err.Error())
	} else if len(actual) != 4 {
		t.Errorf("TvWithParams() was incorrect, expect: [exactly 4 characters], actual: %s.", actual)
	}
}
