package main

import (
	"fmt"
	"os"

	"github.com/flames31/gophercises/link/internal/parse"
)

func main() {

	fileName := os.Args[1]
	file, _ := os.Open(fileName)
	fmt.Println("Links found in this repo are")
	fmt.Println()
	links, _ := parse.ParseHTML(file)
	for _, link := range links {
		fmt.Printf("Text : %v\n", link.Text)
		fmt.Printf("Reference to : %v\n", link.Href)
		fmt.Println()
	}
}
