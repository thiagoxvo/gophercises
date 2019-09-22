package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/thiagoxvo/gophercises/cyoa"
)

func main() {
	port := flag.String("port", "3000", "the port to start the CYOA server")
	file := flag.String("file", "gopher.json", "the JSON file with the CYOA story")
	flag.Parse()
	fmt.Printf("Using the story in %s\n", *file)

	f, err := os.Open(*file)
	if err != nil {
		panic(err)
	}

	story, err := cyoa.JSONStory(f)
	if err != nil {
		panic(err)
	}

	h := cyoa.NewHandler(story)
	fmt.Printf("Starting the server on port: %s\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", *port), h))
}
