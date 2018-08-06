package chance

import (
	"math"
)

// Generate a random int32 number between MinInt32 (-2,147,483,648) and MaxInt32 (2,147,483,647)
func (c *Chance) Integer() int {
	output, _ := c.IntegerWithParams(math.MinInt32, math.MaxInt32)
	return output
}

func (c *Chance) IntegerWithParams(min int, max int) (int, error) {
	output := min + c.Rand.Intn(max-min)
	return output, nil
}
