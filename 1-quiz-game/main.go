package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	filename := flag.String("file", "problems.csv", "question description file")
	flag.Parse()

	csvfile, err := os.Open(*filename)
	defer csvfile.Close()

	if err != nil {
		log.Fatalln("Could not open the csv file", err)
	}
	r := csv.NewReader(csvfile)
	lines, err := r.ReadAll()
	problems := parseLines(lines)

	scanner := bufio.NewScanner(os.Stdin)
	score := 0
	for _, problem := range problems {
		fmt.Printf("Question: %s \n", problem.question)
		fmt.Print("Answer: ")
		scanner.Scan()
		answer := scanner.Text()

		if answer == problem.answer {
			score++
		}
	}
	fmt.Printf("Your scored: %v out of %v", score, len(problems))
}

func parseLines(lines [][]string) []problem {
	problems := make([]problem, len(lines))
	for i, line := range lines {
		problems[i] = problem{
			question: line[0],
			answer:   line[1],
		}
	}
	return problems
}

type problem struct {
	question string
	answer   string
}
