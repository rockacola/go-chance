package chance

import (
	"errors"
	"strings"
)

// Generate a paragraph of semi-pronounceable nonsense word.
func (c *Chance) Paragraph() string {
	output, _ := c.ParagraphWithParams(3, 7)
	return output
}

func (c *Chance) ParagraphWithParams(minSentences int, maxSentences int) (string, error) {
	if minSentences < 1 {
		return "", errors.New("minSentences cannot be less than 1.")
	} else if maxSentences < 1 {
		return "", errors.New("maxSentences cannot be less than 1.")
	} else if minSentences > maxSentences {
		return "", errors.New("minSentences must be smaller than maxSentences.")
	}

	sentenceCount, _ := c.NaturalWithParams(minSentences, minSentences)
	sentences := make([]string, 0)
	for i := 0; i < sentenceCount; i++ {
		s := c.Sentence()
		sentences = append(sentences, s)
	}
	output := strings.Join(sentences, " ")
	return output, nil
}
