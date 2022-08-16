package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	//"reflect"
	"strings"
	"time"
)

func main() {

	filename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'.")
	timeLimit := flag.Int("limit", 20, "time limit for ending the quiz(in seconds).")

	flag.Parse()

	file, err := os.Open(*filename)
	if err != nil {

		exit(fmt.Sprintf("Failed to open file: %s", *filename))
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()

    //fmt.Println(reflect.TypeOf(lines))

	if err != nil {

		exit("Failed to parse the provided csv file.")
	}

	//fmt.Println(lines)

	problems := parseLines(lines)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	correct := 0

	for index, problem := range problems {

		fmt.Println("Problem #%d: %s = ", index+1, problem.question)
		answerCh := make(chan string)

		go func(){
			
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		} ()

		select {

			case <-timer.C :  

				fmt.Println("You scored %d out of %d", correct, len(lines))
				return

			case answer := <-answerCh :
				
				if answer == problem.answer {
	
				    correct++
			}
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