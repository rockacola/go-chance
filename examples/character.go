package main

import (
	"fmt"

	chance "github.com/rockacola/go-chance"
)

func main() {
	fmt.Println("=== Character Demo ===")

	// Arrange
	c := chance.NewChance()

	// Act
	n := 100
	for i := 0; i < n; i++ {
		result, _ := c.CharacterWithParams(true, true, true, true)
		fmt.Printf("#%d: %s\n", i, result)
	}
}
