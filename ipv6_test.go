package chance

import (
	"testing"
)

func TestIpv6(t *testing.T) {
	c := NewChance()
	actual := c.Ipv6()
	if len(actual) < 15 {
		t.Errorf("Ipv6() was incorrect, expect: [15 or more characters], actual: %s.", actual)
	}
}
