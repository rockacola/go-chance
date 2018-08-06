package main

import (
	"fmt"

	chance "github.com/rockacola/go-chance"
)

func main() {
	fmt.Println("=== Basic Demo ===")

	// Arrange
	c := chance.NewChance()

	// Basics
	fmt.Println("Bool():", c.Bool())
	fmt.Println("Character():", c.Character())
	fmt.Println("Natural():", c.Natural())

	// Things
	fmt.Println("Animal():", c.Animal())
}
