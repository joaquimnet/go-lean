package main

import (
	"fmt"
)

func main() {
	const printNumbersUpTo = 10

	fmt.Println("Printing numbers from 1 to", printNumbersUpTo)

	// for loop
	for i := 0; i < printNumbersUpTo; i++ {
		fmt.Println(i + 1)
	}

	fmt.Println("In go you can omit the condition for the for loop:")
	var tempVar = 0
	for {
		tempVar++
		if tempVar > printNumbersUpTo {
			break
		}
		fmt.Println(tempVar)
	}

	fmt.Println("Iterating through the items of an array or slice:")
	var mySlice = []string{"a", "b", "c", "d", "e"}
	for i, item := range mySlice {
		fmt.Println("Index:", i, "Item:", item)
	}
}
