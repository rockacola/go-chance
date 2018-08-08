package chance

import (
	"testing"
)

func TestCpf(t *testing.T) {
	c := NewChance()
	actual := c.Cpf()
	if len(actual) < 9 {
		t.Errorf("Cpf() was incorrect, expect: [9 or more characters], actual: %s.", actual)
	}
}
