package chance

import (
	"testing"
)

func TestGuid(t *testing.T) {
	c := NewChance()
	actual := c.Guid()
	if len(actual) != 36 {
		t.Errorf("Guid() was incorrect, expect: [exactly 36 characters], actual: %s.", actual)
	}
}

func TestGuidWithParams(t *testing.T) {
	c := NewChance()
	version := 5
	actual, err := c.GuidWithParams(version)
	if err != nil {
		t.Errorf("GuidWithParams() execution error: %s", err.Error())
	} else if len(actual) != 36 {
		t.Errorf("GuidWithParams() was incorrect, expect: [exactly 36 characters], actual: %s.", actual)
	}
}
