package chance

import (
	"testing"
)

func TestCpf(t *testing.T) {
	c := NewChance()
	actual := c.Cpf()
	if len(actual) <= 7 {
		t.Errorf("Cpf() was incorrect, expect: [7 or more characters], actual: %s.", actual)
	}
}
