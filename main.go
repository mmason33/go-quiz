package main

import (
	"flag"
	"fmt"
	"go-quiz/csvReader"
	"go-quiz/quiz"
)

func main() {
	csvFlag := flag.String("csvFile", "./problems.csv", "Relative path to the csv that will power the CLI")
	durationFlag := flag.Int("duration", 30, "The time limit for the quiz")
	flag.Parse()
	file := csvReader.CsvReader(*csvFlag)
	score, totalQuestions := quiz.Quiz(file, *durationFlag)
	if score != -1 && totalQuestions != -1 {
		fmt.Println("You got ", score, "/", totalQuestions, " questions correct!")
	}
}
