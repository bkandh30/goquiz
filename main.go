package main

import (
	"bufio"
	"context"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "Question Answer CSV file")
	timeLimit := flag.Int("limit", 30, "Time limit for the quiz in seconds")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
	}
	defer file.Close()

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the CSV file")
	}
	problems := parseLines(lines)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(*timeLimit)*time.Second)
	defer cancel()

	fmt.Printf("Starting the quiz! You have %d seconds.\n", *timeLimit)

	correct := 0
	scanner := bufio.NewScanner(os.Stdin)

problemloop:
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)

		answerCh := make(chan string, 1)

		go func() {
			if scanner.Scan() {
				answerCh <- strings.TrimSpace(scanner.Text())
			}
		}()

		select {
		case <-ctx.Done():
			fmt.Println("\nTime's up!")
			break problemloop
		case answer := <-answerCh:
			if strings.EqualFold(answer, p.a) {
				correct++
			}
		}
	}

	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))

}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))

	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: line[1],
		}
	}
	return ret
}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
