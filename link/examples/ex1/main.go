package main

import (
	"fmt"
	"strings"

	"github.com/thiagoxvo/gophercises/link"
)

var exampleHTML = `
<html>
  <body>
    <h1>Hello!</h1>
		<a href="/other-page">A 
		
		link to another 
		page</a>
		<a href="/page-two">A link to second page</a>
  </body>
</html>`

func main() {
	r := strings.NewReader(exampleHTML)
	links, err := link.Parse(r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", links)
}
