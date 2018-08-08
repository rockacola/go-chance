package main

import (
	"fmt"

	chance "github.com/rockacola/go-chance"
)

func main() {
	fmt.Println("=== Natural Number Demo ===")

	// Arrange
	c := chance.NewChance()
	min := 0
	max := 100

	// Act
	n := 100
	for i := 0; i < n; i++ {
		result, _ := c.NaturalWithParams(min, max)
		fmt.Printf("#%d: %d\n", i, result)
	}
}
