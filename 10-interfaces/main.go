package main

import (
	"fmt"
)

type Speaker interface {
	Speak() string
}

type Animal struct {
	Name  string
	Sound string
}

func (a Animal) Speak() string {
	return fmt.Sprintf("%s says %s", a.Name, a.Sound)
}

func main() {
	var dog = Animal{"Dog", "Woof"}
	var cat = Animal{"Cat", "Meow"}

	var animals = []Speaker{dog, cat}

	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}
