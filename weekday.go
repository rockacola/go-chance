package chance

import "github.com/rockacola/go-chance/data"

// Generate a random weekday.
func (c *Chance) Weekday() string {
	randomIndex := c.Rand.Intn(len(data.Weekdays))
	return data.Weekdays[randomIndex]
}
