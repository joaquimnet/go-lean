package main

import (
	"fmt"
)

type GameCharacter struct {
	Name  string
	Quote string
}

/*
	Notes:

A method can receive a value or a pointer.
To receive a pointer you must use an asterisk.
func SayQuote(c *GameCharacter) {...
*/
func SayQuote(c GameCharacter) {
	fmt.Printf("%s: %s\n", c.Name, c.Quote)
}

func main() {
	/*
		Notes:
		A method is a function that is declared with a receiver.
		Methods can be declared using the func keyword.
	*/

	var mario = GameCharacter{
		Name:  "Mario",
		Quote: "It's-a me, Mario!",
	}

	var link = GameCharacter{
		Name:  "Link",
		Quote: "It's dangerous to go alone! Take this.",
	}

	var dragonborn = GameCharacter{
		Name:  "Dragonborn",
		Quote: "Fus Ro Dah!",
	}

	SayQuote(mario)
	SayQuote(link)
	SayQuote(dragonborn)
}
