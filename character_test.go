package chance

import (
	"testing"
)

func TestCharacter(t *testing.T) {
	c := NewChance()
	actual := c.Character()
	if len(actual) != 1 {
		t.Errorf("Character() was incorrect, expect: [exactly 1 character], actual: %s.", actual)
	}
}

func TestCharacterWithParams(t *testing.T) {
	c := NewChance()
	actual, err := c.CharacterWithParams(true, true, true, true)
	if err != nil {
		t.Errorf("CharacterWithParams() execution error: %s", err.Error())
	} else if len(actual) != 1 {
		t.Errorf("CharacterWithParams() was incorrect, expect: [exactly 1 character], actual: %s.", actual)
	}
}
