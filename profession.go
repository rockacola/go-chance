package chance

import (
	"github.com/rockacola/go-chance/data"
)

// Generate a random profession.
func (c *Chance) Profession() string {
	randomIndex := c.Rand.Intn(len(data.Professions))
	return data.Professions[randomIndex]
}

func (c *Chance) ProfessionWithParams(rank bool) (string, error) {
	professionPrefix := ""
	if rank {
		var prefixes = [...]string{"Apprentice", "Junior", "Senior", "Lead"}
		randomPrefixIndex := c.Rand.Intn(len(prefixes))
		professionPrefix = prefixes[randomPrefixIndex] + " "
	}

	randomIndex := c.Rand.Intn(len(data.Professions))
	output := professionPrefix + data.Professions[randomIndex]
	return output, nil
}
