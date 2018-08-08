package main

import (
	"fmt"

	chance "github.com/rockacola/go-chance"
)

func main() {
	fmt.Println("=== Animal Demo ===")

	// Arrange
	c := chance.NewChance()
	category := "ocean"

	// Act
	n := 100
	for i := 0; i < n; i++ {
		result, _ := c.AnimalWithParams(category)
		fmt.Printf("#%d: %s\n", i, result)
	}
}
