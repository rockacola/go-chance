package main

import (
	"fmt"

	chance "github.com/rockacola/go-chance"
)

func main() {
	fmt.Println("=== Prime Demo ===")

	// Arrange
	c := chance.NewChance()
	min := 1
	max := 10000

	// Act
	n := 10
	for i := 0; i < n; i++ {
		min = c.Rand.Intn(5000)
		max = c.Rand.Intn(10000)
		result, err := c.PrimeWithParams(min, max)
		if err != nil {
			fmt.Printf("#%d (min: %d max: %d): Error: %s\n", i, min, max, err.Error())
		} else {
			fmt.Printf("#%d (min: %d max: %d): %d\n", i, min, max, result)
		}
	}
}
