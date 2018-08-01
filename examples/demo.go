package main

import (
	"fmt"

	chance "github.com/rockacola/go-chance"
)

func main() {
	fmt.Println("=== DEMO ===")

	c := chance.NewChance()
	fmt.Println("Natural():", c.Natural())

	animal, _ := c.AnimalWithParams("fish")
	fmt.Println("AnimalWithParams():", animal)
}
