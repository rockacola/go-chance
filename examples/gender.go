package main

import (
	"fmt"

	chance "github.com/rockacola/go-chance"
)

func main() {
	fmt.Println("=== Gender Demo ===")

	// Arrange
	c := chance.NewChance()

	// Act
	extraGenders := []string{"agender", "genderqueer", "trans"}
	output, _ := c.GenderWithParams(extraGenders)
	fmt.Println("gender:", output)
}
