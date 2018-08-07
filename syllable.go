package chance

import (
	"github.com/rockacola/go-chance/data"
)

// Generate a semi-speakable syllable with 2 or 3 letters.
func (c *Chance) Syllable() string {
	output := c.consonant() + c.vowel()
	if c.Bool() {
		output += c.consonant()
	}
	return output
}

func (c *Chance) consonant() string {
	return c.stringFromPool(1, data.Characters["consonant"])
}

func (c *Chance) vowel() string {
	return c.stringFromPool(1, data.Characters["vowel"])
}
