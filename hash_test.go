package chance

import (
	"testing"
)

func TestHash(t *testing.T) {
	c := NewChance()
	actual := c.Hash()
	if len(actual) != 40 {
		t.Errorf("Hash() was incorrect, expect: [exactly 40 characters], actual: %s.", actual)
	}
}

func TestHashWithParams(t *testing.T) {
	c := NewChance()
	length := 16
	actual, err := c.HashWithParams(length, false)
	if err != nil {
		t.Errorf("HashWithParams() execution error: %s", err.Error())
	} else if len(actual) != length {
		t.Errorf("HashWithParams() was incorrect, expect: [exactly %d characters], actual: %s.", length, actual)
	}
}
