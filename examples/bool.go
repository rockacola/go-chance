package main

import (
	"fmt"
	"reflect"

	chance "github.com/rockacola/go-chance"
)

func main() {
	fmt.Println("=== Bool Demo ===")

	// Arrange
	c := chance.NewChance()
	likelihood := 20 // A number between 0 and 100

	// Act
	n := 1000
	var report = map[bool]int{}
	for i := 0; i < n; i++ {
		result, _ := c.BoolWithParams(likelihood)
		report[result] += 1
		// fmt.Printf("#%d: %t\n", i, result)
	}

	// Report
	keys := reflect.ValueOf(report).MapKeys()
	fmt.Println("## Report")
	fmt.Println("iterations:", n)
	fmt.Println("likelihood:", likelihood)
	for _, key := range keys {
		keyBool := key.Interface().(bool)
		count := report[keyBool]
		percentage := int((float64(count) / float64(n)) * float64(100))
		fmt.Printf("%t: %d (%d%%)\n", keyBool, count, percentage)
	}
}
