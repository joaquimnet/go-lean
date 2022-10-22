package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

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

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of problem,answer")
	flag.Parse()

	file, e := os.ReadFile(*csvFilename)
	if e != nil {
		msg := fmt.Sprintf("Failed to open the csv file: %s\n", *csvFilename)
		exit(msg)
	}

	contents := string(file)

	r := csv.NewReader(strings.NewReader(contents))

	records, e := r.ReadAll()
	if e != nil {
		exit("Failed to parse the csv contents of the file.")
	}
	problems := parseLines(records)
	answers := make([]bool, len(problems))

	for i, problem := range problems {
		question := fmt.Sprintf("Problem #%d: %s =", i+1, problem.q)
		answer := StringPrompt(question)
		correctAnswer := problem.a
		isCorrect := answer == correctAnswer
		answers[i] = isCorrect
	}

	correctAnswers := 0

	for _, v := range answers {
		if v == true {
			correctAnswers += 1
		}
	}

	fmt.Printf("You got %d/%d questions right.", correctAnswers, len(records))
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: strings.TrimSpace(line[0]),
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
