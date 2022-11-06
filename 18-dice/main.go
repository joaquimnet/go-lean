package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	for {
		fmt.Println("[r = roll, q = quit]")
		cmd := strings.ToLower(StringPrompt(">"))
		switch cmd {
		case "r":
			roll1 := rand.Intn(6) + 1
			roll2 := rand.Intn(6) + 1
			msg := "You rolled [%d,%d]"
			if roll1+roll2 == 12 {
				msg += " Epic!"
			}
			fmt.Printf(msg+"\n", roll1, roll2)
		case "q":
			fmt.Println("Byeeeee! o/")
			return
		}
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
