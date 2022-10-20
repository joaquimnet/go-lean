package main

import (
	"fmt"
)

func main() {
	/*
		Notes:
		A map is a collection of key-value
		pairs. The key is used to look up
		the value. The key can be of any
		type that is comparable. The value
		can be of any type.
	*/
	var gameProtagonistsMap = map[string]string{
		"mario": "Mario",
		"zelda": "Link",
		"sonic": "Sonic",
	}

	// initialize an empty map
	// agesMap := make(map[string]int)

	fmt.Println("myMap:", gameProtagonistsMap)
	fmt.Println("Length of the map:", len(gameProtagonistsMap))
	gameProtagonistsMap["kirby"] = "Kirby"
	fmt.Println("Adding new item to map:", gameProtagonistsMap)
	delete(gameProtagonistsMap, "sonic")
	fmt.Println("Deleting item from map:", gameProtagonistsMap)
}
