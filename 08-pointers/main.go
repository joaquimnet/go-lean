package main

import (
	"fmt"
)

func main() {
	/*
		Notes:
		A pointer is a variable that stores the memory address of another variable.
		A pointer is declared using an asterisk.
		Using the ampersand you can get its memory address.
	*/

	var year = 2022

	fmt.Println("This code was written in:", year)
	increaseYear(&year)
	fmt.Println("The following year is:", year)
}

func increaseYear(year *int) {
	*year++
}
