package chance

import (
	"strings"
)

// Generate a random city name.
func (c *Chance) City() string {
	return strings.Title(c.Word())
}
