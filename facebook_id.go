package chance

import (
	"github.com/rockacola/go-chance/data"
)

// Generate a random Facebook ID.
func (c *Chance) FacebookId() string {
	return "1000" + c.stringFromPool(11, data.Characters["numeric"])
}
