package chance

import (
	"testing"
)

func TestBool(t *testing.T) {
	c := NewChance()
	var _ = c.Bool()
}

func TestBoolWithParams(t *testing.T) {
	c := NewChance()
	var actual, err = c.BoolWithParams(0)
	if err != nil {
		t.Errorf("BoolWithParams() execution error: %s", err.Error())
	} else if actual == true {
		t.Errorf("BoolWithParams() was incorrect, expect: [false], actual: %t.", actual)
	}
}
