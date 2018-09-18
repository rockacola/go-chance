package chance

import (
	"testing"
)

func TestGoogleAnalytics(t *testing.T) {
	c := NewChance()
	actual := c.GoogleAnalytics()
	if len(actual) != 12 {
		t.Errorf("GoogleAnalytics() was incorrect, expect: [exactly 12 characters], actual: %s.", actual)
	}
}
