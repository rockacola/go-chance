package chance

import (
	"testing"
)

func TestGeohash(t *testing.T) {
	c := NewChance()
	actual := c.Geohash()
	if len(actual) <= 0 {
		t.Errorf("Geohash() was incorrect, expect: [1 or more characters], actual: %s.", actual)
	}
}

func TestGeohashWithParams(t *testing.T) {
	c := NewChance()
	length := 7
	actual, err := c.GeohashWithParams(length)
	if err != nil {
		t.Errorf("GeohashWithParams() execution error: %s", err.Error())
	} else if len(actual) <= 0 {
		t.Errorf("GeohashWithParams() was incorrect, expect: [1 or more characters], actual: %s.", actual)
	}
}
