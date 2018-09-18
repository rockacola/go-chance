package chance

import (
	"testing"
)

func TestDomain(t *testing.T) {
	c := NewChance()
	actual := c.Domain()
	if len(actual) <= 1 {
		t.Errorf("Domain() was incorrect, expect: [1 or more characters], actual: %s.", actual)
	}
}

func TestDomainWithParams(t *testing.T) {
	c := NewChance()
	tld := "com"
	actual, err := c.DomainWithParams(tld)
	if err != nil {
		t.Errorf("DomainWithParams() execution error: %s", err.Error())
	} else if len(actual) <= 1 {
		t.Errorf("DomainWithParams() was incorrect, expect: [1 or more characters], actual: %s.", actual)
	}
}
