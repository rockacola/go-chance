package main

import (
	"fmt"

	chance "github.com/rockacola/go-chance"
)

func main() {
	fmt.Println("=== String Demo ===")

	// Arrange
	c := chance.NewChance()

	// Act
	n := 100
	length := 8
	for i := 0; i < n; i++ {
		result, _ := c.StringWithParams(length, true, true, true, true)
		fmt.Printf("#%d: %s\n", i, result)
	}
}
