package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
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

	scanner := bufio.NewScanner(os.Stdin)
	score := 0
	questions := 0
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Question: %s \n", record[0])
		fmt.Print("Answer: ")
		scanner.Scan()
		answer := scanner.Text()

		if len(answer) == 0 {
			log.Fatalln("Invalid answer!")
		}
		if answer == record[1] {
			score++
		}
		questions++
	}
	fmt.Printf("Your scored: %v out of %v", score, questions)
}
