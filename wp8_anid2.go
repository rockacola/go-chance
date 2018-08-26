package chance

import (
	"github.com/rockacola/go-chance/data"
)

// Generate a random Windows Phone 8 ANID2.
func (c *Chance) Wp8Anid2() string {
	return c.stringFromPool(32, data.Characters["hex"])
}
