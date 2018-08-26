package chance

import (
	"github.com/rockacola/go-chance/data"
)

// Generate an Apple Push Token.
func (c *Chance) AppleToken() string {
	return c.stringFromPool(64, data.Characters["hex"])
}
