package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	filename := flag.String("file", "problems.csv", "question description file")
	timeLimit := flag.Int("limit", 30, "the limit for the quiz  in seconds")
	flag.Parse()

	csvfile, err := os.Open(*filename)
	defer csvfile.Close()

	if err != nil {
		log.Fatalln("Could not open the csv file", err)
	}
	r := csv.NewReader(csvfile)
	lines, _ := r.ReadAll()
	problems := parseLines(lines)

	scanner := bufio.NewScanner(os.Stdin)
	score := 0
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	for _, problem := range problems {

		fmt.Printf("Question: %s \n", problem.question)
		answerCh := make(chan string)
		go func() {
			fmt.Print("Answer: ")
			scanner.Scan()
			answer := scanner.Text()
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println("Time is over!")
			fmt.Printf("\nYour scored: %v out of %v", score, len(problems))
			return
		case answer := <-answerCh:
			if answer == problem.answer {
				score++
			}
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
