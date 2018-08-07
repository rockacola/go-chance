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
	likelihood := 0
	var actual, err = c.BoolWithParams(likelihood)
	if err != nil {
		t.Errorf("BoolWithParams() execution error: %s", err.Error())
	} else if actual == true {
		t.Errorf("BoolWithParams() was incorrect, expect: [false], actual: %t.", actual)
	}
}
