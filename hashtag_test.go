package chance

import (
	"testing"
)

func TestHashtag(t *testing.T) {
	c := NewChance()
	actual := c.Hashtag()
	if len(actual) < 2 {
		t.Errorf("Hashtag() was incorrect, expect: [2 or more characters], actual: %s.", actual)
	}
}
