package chance

import (
	"strconv"

	"github.com/rockacola/go-chance/data"
)

// Generates a Canadian postal code.
func (c *Chance) CanadianPostal() string {
	// Postal District
	pd := c.stringFromPool(1, "XVTSRPNKLMHJGECBA")
	// Forward Sortation Area (FSA)
	code, _ := c.NaturalWithParams(0, 9)
	fsa := pd + strconv.Itoa(code) + c.stringFromPool(1, data.Characters["upper"])
	// Local Delivery Unit (LDU)
	code2, _ := c.NaturalWithParams(0, 9)
	code3, _ := c.NaturalWithParams(0, 9)
	ldu := strconv.Itoa(code2) + c.stringFromPool(1, data.Characters["upper"]) + strconv.Itoa(code3)

	return fsa + " " + ldu
}
