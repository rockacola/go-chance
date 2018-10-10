package chance

import "github.com/rockacola/go-chance/data"

// Generate a random month name.
func (c *Chance) Month() string {
	randomIndex := c.Rand.Intn(len(data.Months))
	return data.Months[randomIndex]
}
