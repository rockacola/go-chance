package chance

import "strconv"

// Generate random coordinates, which are latitude and longitude, comma separated.
func (c *Chance) Coordinates() string {
	return strconv.FormatFloat(c.Latitude(), 'f', 8, 64) + ", " + strconv.FormatFloat(c.Longitude(), 'f', 8, 64)
}
