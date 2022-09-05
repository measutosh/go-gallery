package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// collect the file using flag package
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of questions and answers")
	// creating a flag for the timer
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	// after collecting it parse the file
	flag.Parse()

	// read the file, use * before the name cuz it's gonna be a pointer to a string
	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("failed to open the CSV file %s\n", *csvFilename))
	}

	// using Newreader function to read the CSV file that takes an io.reader
	r := csv.NewReader(file)
	// read all the lines from the CSV file
	lines, err := r.ReadAll()
	if err != nil {
		exit("failed to parse the provided csv file")
	}
	problems := parseLines(lines)

	
	// setting up the timer with a rhythm so that once the timer is over the numbers should rearrange accordingly
	// if *timeLimit=30 then give a timer of time.Second(which is equals to 30seconds)
	// to avoid a type error, a typecasting into duration is done with timeLimit
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	    

	// to keep the scores
	correct := 0

	
	// to show the problems to the users, take response from them to give them back the results
	for i, p := range problems {
		// go routine will be introduced here
		fmt.Printf("Problem number #%d : %s = ", i + 1, p.q)
		// to get the answer from scanf from below an answer channel is created
		answerCh := make(chan string)

		go func() {
			var answer string
			// %s will pickup the empty space too, scanf trims all spaces
			// a newline is added so that when the user clicks enter after typing something then that would be taken into account
			// a reference to answer is also added to pickup the value of answer
			fmt.Scanf("%s\n", &answer)
			// adding the closure here that came outside of this anonymous function will
			// it sends the answer from scanf to answer channel
			answerCh <- answer

		}() //here this () will all this above func
		// after this go routine will the default wont be required anymore in the below codeblock

		// if it gets a message from the timer channel then it will stop or else it will keep showing questions
		select {
		case <-timer.C:
			fmt.Printf("\nYou scored %d out of %d tests.\n", correct, len(problems))
			return
		case answer := <-answerCh :
			if answer == p.a {
				correct++
			} 
		}
	}

	// printing the results
	fmt.Printf("You scored %d out of %d tests.\n", correct, len(problems))
}

// use the result came from slice to create problems
func parseLines(lines [][]string) []problem {
	//  return the problem same length as the total number of lines
	ret := make([]problem, len(lines))
	for i, line := range lines {
		// creating new problems 
		ret[i] = problem{
			q: line[0],
			// handling an edge case : if the csv file contains invalid values or spaces
			// makes sure that all the questions we are getting are actually answerable.
			a: strings.TrimSpace(line[1]),
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
