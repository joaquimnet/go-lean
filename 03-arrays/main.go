package main

import "fmt"

func main() {
	/*
		Notes:
		An array can only contain values of the same type.
		The array cannot be resized – you have to explicitly define the length of an array in Go.
		That’s part of the type of an array.
		Also, you cannot use a variable to set the length of the array.
	*/
	var myNamesArray = [...]string{"Link", "Mario", "Kirby"}

	fmt.Println("myNamesArray:", myNamesArray)
	fmt.Println("First name:", myNamesArray[0])
	fmt.Println("Length of the array:", len(myNamesArray))

	// arrays are value types
	var my2ndNamesArray = myNamesArray
	my2ndNamesArray[0] = "Zelda"
	fmt.Println("Arrays are value types:")
	fmt.Println("myNamesArray:", myNamesArray)
	fmt.Println("my2ndNamesArray:", my2ndNamesArray)

	// slices
	fmt.Println("Slices are like array, but better.")
	var myNamesSlice = []string{"SlicedLink", "SlicedMario", "SlicedKirby"}
	fmt.Println("myNamesSlice:", myNamesSlice)
	fmt.Println("Length of the slice:", len(myNamesSlice))
	fmt.Println("Capacity of the slice:", cap(myNamesSlice))
	fmt.Println("New slice of 3 empty strings:", make([]string, 3))
	fmt.Println("Appended slice (this is a new slice):", append(myNamesSlice, "SlicedZelda"))
	fmt.Println("Slice from an array:", myNamesArray[:])
	var copiedSlice1 = myNamesSlice[:]
	var copiedSlice2 = myNamesSlice[:]
	fmt.Println("Slices created from the same array share the same memory:", copiedSlice1, copiedSlice2)
	myNamesSlice[0] = "SlicedZelda2"
	copiedSlice2[1] = "SlicedMario2"
	fmt.Println("Slices created from the same array share the same memory:", copiedSlice1, copiedSlice2)
}
