package chance

import (
	"strings"

	"github.com/rockacola/go-chance/data"
)

// Generate a random IPv6 address.
func (c *Chance) Ipv6() string {
	sections := []string{}
	for i := 0; i < 8; i++ {
		sections = append(sections, c.stringFromPool(4, data.Characters["hex"]))
	}
	return strings.Join(sections, ":")
}
