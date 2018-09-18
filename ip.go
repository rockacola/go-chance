package chance

import "strconv"

// Generate a random IP address.
func (c *Chance) Ip() string {
	s1, _ := c.NaturalWithParams(1, 254)
	s2, _ := c.NaturalWithParams(0, 255)
	s3, _ := c.NaturalWithParams(0, 255)
	s4, _ := c.NaturalWithParams(1, 254)
	return strconv.Itoa(s1) + "." + strconv.Itoa(s2) + "." + strconv.Itoa(s3) + "." + strconv.Itoa(s4)
}
