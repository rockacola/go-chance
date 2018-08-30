package chance

import (
	"github.com/rockacola/go-chance/data"
)

// Generate a random Google Analytics tracking code.
func (c *Chance) GoogleAnalytics() string {
	return "UA-" + c.stringFromPool(6, data.Characters["numeric"]) + "-" + c.stringFromPool(2, data.Characters["numeric"])
}
