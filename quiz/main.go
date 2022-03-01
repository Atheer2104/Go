package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// the file extension is csv and the default name of that file is problems.csv
	// the user could als
	// and lastly the string is telling the user what this file does
	cvsFilename := flag.String("csv", "problems.csv", " a csv file in the format of 'question, answear'")
	flag.Parse()

	// csvFilename by default is a pointer
	file, err := os.Open(*cvsFilename)
	if err != nil {
		exit(fmt.Sprintln("failed to open the CSV file: %s\n", *cvsFilename))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided the csv file ")
	}
	problems := parseLines(lines)
	
	correct := 0
	// i - index, p - problems
	for i, p := range problems {
		// the value after # is going to be i+1 and %s is going to be question 
		fmt.Printf("Problem #%d: %s = \n", i+1, p.question)
		var answear string 
		// getting user input and assignig it to answear
		fmt.Scanf("%s\n", &answear)
		if answear == p.answear {
			correct++
		} 
	}
	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}

type problem struct {
	question string
	answear  string
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))

	for i, line := range lines {
		ret[i] = problem{
			question: line[0],
			answear:  strings.TrimSpace(line[1]),
		}
	}

	return ret
}

func exit(msg string) {
	fmt.Println(msg)
	// exit with 1 means something went wrong
	os.Exit(1)
}
