package main

import (
	"fmt"
)

func main() {
	/*
		Notes:
		A struct is a collection of fields.
		Each field has a name and a type.
		Fields can be accessed using a dot.
		Uppercase fields are exported.
		Lowercase fields are private to the package.
	*/
	type Game struct {
		Name  string
		Year  int
		Genre string
	}

	var game1 = Game{
		Name:  "Super Mario Bros.",
		Year:  1985,
		Genre: "Platform",
	}

	var game2 = Game{"The Legend of Zelda", 1986, "Action-Adventure"}

	fmt.Println("game1:", game1)
	fmt.Println("game2:", game2)
}
