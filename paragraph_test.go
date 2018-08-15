package chance

import (
	"testing"
)

func TestParagraph(t *testing.T) {
	c := NewChance()
	actual := c.Paragraph()
	if len(actual) < 4 {
		t.Errorf("Paragraph() was incorrect, expect: [4 or more characters], actual: %s.", actual)
	}
}

func TestParagraphWithParams(t *testing.T) {
	c := NewChance()
	min := 1
	max := 3
	actual, err := c.ParagraphWithParams(min, max)
	if err != nil {
		t.Errorf("ParagraphWithParams() execution error: %s", err.Error())
	} else if len(actual) < 4 {
		t.Errorf("ParagraphWithParams() was incorrect, expect: [4 or more characters], actual: %s.", actual)
	}
}
