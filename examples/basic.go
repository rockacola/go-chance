package main

import (
	"fmt"

	chance "github.com/rockacola/go-chance"
)

func main() {
	fmt.Println("=== Basic Demo ===")

	// Arrange
	c := chance.NewChance()

	// Rand
	fmt.Println("Rand.Float32():", c.Rand.Float32())
	fmt.Println("Rand.Float64():", c.Rand.Float64())

	// Basics
	fmt.Println("Bool():", c.Bool())
	fmt.Println("Character():", c.Character())
	fmt.Println("Floating():", c.Floating())
	fmt.Println("Integer():", c.Integer())
	fmt.Println("Letter():", c.Letter())
	fmt.Println("Natural():", c.Natural())
	fmt.Println("Prime():", c.Prime())
	fmt.Println("String():", c.String())

	// Person
	fmt.Println("Age():", c.Age())
	fmt.Println("CPF():", c.Cpf())
	fmt.Println("FirstName():", c.FirstName())
	fmt.Println("Gender():", c.Gender())
	fmt.Println("LastName():", c.LastName())
	fmt.Println("NamePrefix():", c.NamePrefix())
	fmt.Println("NameSuffix():", c.NameSuffix())
	fmt.Println("Name():", c.Name())

	// // Text
	// fmt.Println("Syllable():", c.Syllable())
	// fmt.Println("Word():", c.Word())
	// fmt.Println("Sentence():", c.Sentence())
	// fmt.Println("Paragraph():", c.Paragraph())

	// // Things
	// fmt.Println("Animal():", c.Animal())

	// // Misc
	// fmt.Println("Coin():", c.Coin())
	// fmt.Println("Guid():", c.Guid())
	// fmt.Println("Radio():", c.Radio())
	// fmt.Println("Tv():", c.Tv())
	// fmt.Println("Hash():", c.Hash())
	// fmt.Println("Normal():", c.Normal())
}
