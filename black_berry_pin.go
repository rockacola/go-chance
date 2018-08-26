package chance

import (
	"github.com/rockacola/go-chance/data"
)

// Generate a random BlackBerry Device PIN.
func (c *Chance) BlackBerryPin() string {
	return c.stringFromPool(8, data.Characters["hex"])
}
