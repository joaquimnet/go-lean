package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		res := StringPrompt(">")
		if cmd := strings.ToLower(res); cmd == "q" || cmd == "exit" {
			break
		}
		fmt.Println(strings.ToUpper(res))
	}
}

func StringPrompt(label string) string {
	var s string
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, label+" ")
		s, _ = r.ReadString('\n')
		if s != "" {
			break
		}
	}
	return strings.TrimSpace(s)
}
