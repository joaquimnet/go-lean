package main

import (
	"fmt"
	"strings"
)

func main() {
	var myStr = "Hello, World!"
	fmt.Println(myStr)

	var my2ndStr = "I used to be an adventurer like you but then I took an arrow in the knee."
	var my2ndStrExtra = "Isn't that tragic?"

	fmt.Println("my2ndStr's length is:", len(my2ndStr))
	fmt.Println("The first character is:", my2ndStr[0])
	fmt.Println("The last character is:", my2ndStr[len(my2ndStr)-1])
	fmt.Println("The first 10 characters are:", my2ndStr[:10])
	fmt.Println("The last 10 characters are:", my2ndStr[len(my2ndStr)-10:])
	fmt.Println("This is a full copy of the string:", my2ndStr[:])
	fmt.Println("Concatenation:", my2ndStr+" "+my2ndStrExtra)
	fmt.Println("Does the string starts with 'I used'?", strings.HasPrefix(my2ndStr, "I used"))
	fmt.Println("Does the string ends with 'knee.'?", strings.HasSuffix(my2ndStr, "knee."))
	fmt.Println("The string, in upper case:", strings.ToUpper(my2ndStr))
	fmt.Println("Does the string contain the word 'adventurer'?", strings.Contains(my2ndStr, "adventurer"))
	fmt.Println("How many spaces does the string have?", strings.Count(my2ndStr, " "))
	var reversedStr = strings.ReplaceAll(strings.ReplaceAll(my2ndStr, "you", "me"), "I", "you")
	reversedStr = strings.ToUpper(reversedStr[0:1]) + reversedStr[1:]
	fmt.Println("Reversed string:", reversedStr)
}
