package main

import (
	"fmt"
	"strings"
)

func main() {
	const Name = "Kaffe"

	switch name := Name; {
	case strings.ToLower(name) == "kaffe":
		fmt.Println("Yes, this works.")
	default:
		fmt.Println("hehe")
	}

	if role := getRole(); role == "Customer" {
		fmt.Println("Access denied")
	} else if role == "Manager" {
		fmt.Println("I should give Kaffe a raise.")
	} else {
		fmt.Println("Who am I?")
	}
}

func getRole() string {
	return "Manager"
}
