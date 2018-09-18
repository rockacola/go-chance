package chance

import (
	"testing"
)

func TestCompany(t *testing.T) {
	c := NewChance()
	actual := c.Company()
	if len(actual) < 2 {
		t.Errorf("Company() was incorrect, expect: [2 or more characters], actual: %s.", actual)
	}
}
