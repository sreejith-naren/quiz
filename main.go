package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"reflect"
	"strings"
)

func main() {

	filename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'.")
	flag.Parse()

	file, err := os.Open(*filename)
	if err != nil {

		exit(fmt.Sprintf("Failed to open file: %s", *filename))
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()

    fmt.Println(reflect.TypeOf(lines))

	if err != nil {

		exit("Failed to parse the provided csv file.")
	}

	fmt.Println(lines)

	problems := parseLines(lines)

	correct := 0

	for index, problem := range problems {

		fmt.Println("Problem #%d: %s = \n", index+1, problem.question)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == problem.answer {

			correct++
		}

	}

	fmt.Println("You scored %d out of %d", correct, len(lines))
}

func parseLines(lines [][]string) []problem {

	ret := make([]problem, len(lines))
	for i, line := range lines {

		ret[i] = problem {

			question : line[0],
			answer : strings.TrimSpace(line[1]),
		}
	}
	return ret
}

type problem struct {

	question string
	answer string
}

func exit(msg string) {

	fmt.Println(msg)
	os.Exit(1)
}