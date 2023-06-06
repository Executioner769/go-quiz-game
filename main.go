package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "A CSV file of the format 'question and answer'")
	flag.Parse()
	// _ = csvFilename

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open CSV file: %s\n", *csvFilename))
	}

	r := csv.NewReader(file)

	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to format the CSV file\n")
	}

	problems := parseLines(lines)

	// fmt.Println(problems)

	startQuiz(problems)

}

func startQuiz(problems []problem) {

	correct := 0

	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.q)
		var ans string
		fmt.Scanf("%s\n", &ans)
		if ans == p.a {
			correct++
		}
	}

	fmt.Printf("You scored %d out of %d!\n", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))

	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]), //If there are any spaces in input csv file we can make sure those are removed
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
