package chance

import (
	"strconv"
	"strings"

	"github.com/rockacola/go-chance/data"
)

// Generate a random Windows Phone 7 ANID.
func (c *Chance) Wp7Anid() string {
	guid := strings.ToUpper(strings.Replace(c.Guid(), "-", "", -1))
	integer, _ := c.IntegerWithParams(0, 9)
	return "A=" + guid + "&E=" + c.stringFromPool(3, data.Characters["hex"]) + "&w=" + strconv.Itoa(integer)
}
