package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	var csv string
	csvDescription := "a csv file in the format of 'question,answer' (default\"problems.csv\")"
	flag.StringVar(&csv, "csv", "problems.csv", csvDescription)
	flag.Parse()

	problems, err := parseFile(csv)
	if err != nil {
		panic(err)
	}
	fmt.Println("Answer the following problems!")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	reader := bufio.NewReader(os.Stdin)
	score := 0
	total := len(problems)
	for problem, answer := range problems {
		fmt.Print(problem + ": ")
		guess, _ := reader.ReadString('\n')
		guess = strings.TrimSuffix(guess, "\n")
		if guess == answer {
			score += 1
		}
	}
	fmt.Println("You scored ", score, "/", total)
}

func parseFile(file string) (problems map[string]string, err error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return nil, err
	}

	problems = make(map[string]string, len(lines))
	for _, problem := range lines {
		problems[problem[0]] = problem[1]
	}
	return problems, nil
}
