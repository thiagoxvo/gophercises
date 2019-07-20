package main

import (
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

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Question: %s\n", record[0])
	}
}
