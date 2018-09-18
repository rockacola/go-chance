package chance

import (
	"testing"
)

func TestEmail(t *testing.T) {
	c := NewChance()
	actual := c.Email()
	if len(actual) <= 1 {
		t.Errorf("Email() was incorrect, expect: [1 or more characters], actual: %s.", actual)
	}
}

func TestEmailWithParams(t *testing.T) {
	c := NewChance()
	domain := "test.com"
	actual, err := c.EmailWithParams(domain)
	if err != nil {
		t.Errorf("EmailWithParams() execution error: %s", err.Error())
	} else if len(actual) <= 1 {
		t.Errorf("EmailWithParams() was incorrect, expect: [1 or more characters], actual: %s.", actual)
	}
}
