package chance

import (
	"github.com/rockacola/go-chance/data"
)

// Generate a random TLD.
func (c *Chance) Tld() string {
	randomIndex := c.Rand.Intn(len(data.Tlds))
	return data.Tlds[randomIndex]
}
