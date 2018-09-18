package chance

import (
	"testing"
)

func TestProfession(t *testing.T) {
	c := NewChance()
	actual := c.Profession()
	if len(actual) < 0 {
		t.Errorf("Profession() was incorrect, expect: [1 or more characters], actual: %s.", actual)
	}
}

func TestProfessionWithParams(t *testing.T) {
	c := NewChance()
	rank := true
	actual, err := c.ProfessionWithParams(rank)
	if err != nil {
		t.Errorf("ProfessionWithParams() execution error: %s", err.Error())
	} else if len(actual) < 0 {
		t.Errorf("ProfessionWithParams() was incorrect, expect: [1 or more characters], actual: %s.", actual)
	}
}
