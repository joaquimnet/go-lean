package main

import (
	"fmt"
	"strings"
)

func main() {
	/*
		Notes:
		A function is a block of code that performs an action.
		Functions can be called using a name and parentheses.
		Functions can have parameters and return values.
		Functions can be declared using the func keyword.
	*/

	fmt.Println("Sum function (1, 2, 3, 4, 5):", sum(1, 2, 3, 4, 5))
	fmt.Println("Uwuify function (Hello World!):", uwuify("Hello World!"))
	fmt.Println("Uwuify function (Famous Skyrim quote):", uwuify("I used to be an adventurer like you but then I took an arrow in the knee."))
}

func sum(numbers ...int) int {
	var total int
	for _, number := range numbers {
		total += number
	}
	return total
}

func uwuify(text string) string {
	var result string
	result = strings.ReplaceAll(text, "r", "w")
	result = strings.ReplaceAll(result, "R", "W")
	result = strings.ReplaceAll(result, "l", "w")
	result = strings.ReplaceAll(result, "L", "W")
	return result
}
