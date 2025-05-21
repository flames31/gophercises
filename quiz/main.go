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
	score := 0
	fileName := flag.String("file", "problems.csv", "file name(default:problems.csv)")
	quizTime := flag.Int("quizTime", 5, "time for quiz")
	file, err := os.Open(*fileName)
	if err != nil {
		log.Fatal("ERROR : openeing file")
	}

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	questions, err := reader.ReadAll()
	if err != nil {
		log.Fatal("ERROR : openeing file")
	}

	timer := time.NewTimer(time.Second * time.Duration(*quizTime))

	scanner := bufio.NewScanner(os.Stdin)

quizLoop:
	for idx, row := range questions {
		fmt.Printf("Question %v\n%v\n", idx+1, row[0])
		answer := make(chan string)

		go func() {
			if scanner.Scan() {
				answer <- scanner.Text()
			}
		}()

		select {
		case <-timer.C:
			fmt.Println("Time is up!")
			break quizLoop
		case providedAnswer := <-answer:
			if providedAnswer == row[1] {
				score++
			}
		}
	}

	fmt.Printf("You scored %v points out of %v.\n", score, len(questions))
}
