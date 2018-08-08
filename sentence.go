package chance

import (
	"errors"
	"strings"
)

// Generate a semi-pronounceable nonsense word.
func (c *Chance) Sentence() string {
	output, _ := c.SentenceWithParams(12, 18)
	return output
}

func (c *Chance) SentenceWithParams(minWords int, maxWords int) (string, error) {
	if minWords < 1 {
		return "", errors.New("minWords cannot be less than 1.")
	} else if maxWords < 1 {
		return "", errors.New("maxWords cannot be less than 1.")
	} else if minWords > maxWords {
		return "", errors.New("minWords must be smaller than maxWords.")
	}

	wordCount, _ := c.NaturalWithParams(minWords, maxWords)
	words := make([]string, 0)
	for i := 0; i < wordCount; i++ {
		w := c.Word()
		words = append(words, w)
	}
	output := strings.Join(words, " ")
	return output, nil
}
