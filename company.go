package chance

import (
	"github.com/rockacola/go-chance/data"
)

// Generate a random company name.
func (c *Chance) Company() string {
	randomIndex := c.Rand.Intn(len(data.Companies))
	return data.Companies[randomIndex]
}
