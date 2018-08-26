package chance

import (
	"testing"
)

func TestAndroidId(t *testing.T) {
	c := NewChance()
	actual := c.AndroidId()
	if len(actual) != 183 {
		t.Errorf("AndroidId() was incorrect, expect: [183 characters], actual: %s.", actual)
	}
}
