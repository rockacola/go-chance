package chance

import (
	"testing"
)

func TestIp(t *testing.T) {
	c := NewChance()
	actual := c.Ip()
	if len(actual) < 7 {
		t.Errorf("Ip() was incorrect, expect: [7 or more characters], actual: %s.", actual)
	}
}
