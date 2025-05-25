package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Enter the json file name : ")
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	fileName := scanner.Text()

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("ERROR : No such file exists! : %v\n", err)
		os.Exit(-1)
	}

	jsonData, err := JsonConverter(file)
	if err != nil {
		fmt.Printf("ERROR : Decoding JSON! : %v\n", err)
		os.Exit(-1)
	}

	chapter := "intro"

	for {

		if len(jsonData[chapter].Options) == 0 {
			fmt.Println("Adventure is over. Thanks for reading!")
			os.Exit(0)
		}
		fmt.Printf("\nTitle : %v\n\n", jsonData[chapter].Title)
		fmt.Printf("Story : %v\n\n", jsonData[chapter].Story)
		fmt.Println("Options :")
		for idx, option := range jsonData[chapter].Options {
			fmt.Printf("Option %v : %v\n", idx+1, option.Text)
		}

		fmt.Println("What option do you use?")
		optionNumber := -1
		for scanner.Scan() {
			optionNumber, err = strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Printf("ERROR : Converting option! : %v\n", err)
				os.Exit(-1)
			}

			if len(jsonData[chapter].Options) < (optionNumber-1) || optionNumber < 0 {
				fmt.Println("Invalid option! Try again")
				continue
			} else {
				break
			}
		}

		chapter = jsonData[chapter].Options[optionNumber-1].Arc
	}
}
