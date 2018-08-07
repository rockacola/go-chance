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
	fmt.Println("Floating():", c.Floating())
	fmt.Println("Integer():", c.Integer())
	fmt.Println("Letter():", c.Letter())
	fmt.Println("Natural():", c.Natural())
	fmt.Println("Prime():", c.Prime())
	fmt.Println("String():", c.String())

	// Text
	fmt.Println("Syllable():", c.Syllable())
	fmt.Println("Word():", c.Word())

	// Things
	fmt.Println("Animal():", c.Animal())
}
