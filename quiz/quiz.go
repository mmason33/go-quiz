package quiz

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var questionPrefix string = "What is "
var questionSuffix string = "?"

type question struct {
	q string
	a string
}

// Parse CSV
// [][]string => []Question
func createQuestionsFromCsv(fileContents [][]string) []question {
	questions := make([]question, len(fileContents))
	for i, sl := range fileContents {
		questions[i] = question{}
		for in, subSl := range sl {
			switch in {
			case 0:
				questions[i].q = subSl
			case 1:
				questions[i].a = subSl
			}
		}
	}

	return questions
}

// User prompt to explain quiz rules and proceed after user input
func userPrompt(s string) bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s [y/n]: ", s)

		response, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		response = strings.ToLower(strings.TrimSpace(response))

		if response == "y" || response == "yes" {
			return true
		} else if response == "n" || response == "no" {
			return false
		}
	}
}

// Simple stdin io reader
func getUserInput(input chan<- string) {
	reader := bufio.NewReader(os.Stdin)
	response, nErr := reader.ReadString('\n')
	if nErr != nil {
		panic(nErr)
	}

	response = strings.Replace(response, "\n", "", -1)
	input <- response
}

// Evaluate the given answer
// Receive value from channel
// Keep checking timer channel for value
// When timer receives value returning -1 will break execution up the chain
func ask(q question, timer <-chan time.Time, answers <-chan string) int {
	fmt.Println(questionPrefix + q.q + questionSuffix)

	for {
		select {
		case <-timer:
			return -1
		case a := <-answers:
			if a == q.a {
				return 1
			}

			return 0
		}
	}
}

// Quiz runs all processes for the package
// Runs prompt initially to inform user of the quiz rules - asks for confirmation
// Block is dependent on yes/no confirmation from user
// Consumes the contents of a csv and creates a slice of type Question
// Duration is defined by cli flag or default - passed from main
// Goroutine to get user input concurrently without blocking the evaluation of the answers
// Timer is implemented using the returned channel from time.NewTimer
func Quiz(fileContents [][]string, duration int) (int, int) {
	directions := userPrompt("+-+-+ +-+-+-+-+\n|G|o| |Q|u|i|z|\n+-+-+ +-+-+-+-+\nAre you ready for the quiz?\nKeep in mind you will have " + strconv.Itoa(duration) + " seconds to complete it!")

	if !directions {
		return -1, -1
	}

	fmt.Println("-----------------------")
	totalScore := 0
	answers := make(chan string)
	timer := time.NewTimer(time.Second * time.Duration(duration))
	data := createQuestionsFromCsv(fileContents)
	totalQuestions := len(data)
	for _, question := range data {
		go getUserInput(answers)
		score := ask(question, timer.C, answers)
		if score == -1 {
			fmt.Println("You ran out of time :(")
			break
		}

		totalScore += score
	}

	return totalScore, totalQuestions
}
