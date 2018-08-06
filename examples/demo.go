package main

import (
	"fmt"

	chance "github.com/rockacola/go-chance"
)

func main() {
	fmt.Println("=== DEMO ===")

	// Constructor
	c := chance.NewChance()

	// Maths
	fmt.Println("Natural():", c.Natural())

	// Animals
	animal, _ := c.AnimalWithParams("desert")
	fmt.Println("AnimalWithParams():", animal)
}
