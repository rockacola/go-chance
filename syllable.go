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
	randomIndex := c.Rand.Intn(len(data.Characters["consonant"]) - 1)
	return string(data.Characters["consonant"][randomIndex])
}

func (c *Chance) vowel() string {
	randomIndex := c.Rand.Intn(len(data.Characters["vowel"]) - 1)
	return string(data.Characters["vowel"][randomIndex])
}
