# Go Quiz
This is a CLI exercise from [https://github.com/gophercises/quiz](here) one of many found on [https://gophercises.com/](https://gophercises.com/).

- Package seperation
- CSV parsing
- Standard I/O operations
- CLI flag parsing and defaults
- Goroutines
- Unidirectional Channels
- For Select Channels evaluations
- Int to String type conversion

### Defaults
```
// clone
git clone https://github.com/mmason33/go-quiz.git

cd go-quiz

// run
go run main.go

+-+-+ +-+-+-+-+
|G|o| |Q|u|i|z|
+-+-+ +-+-+-+-+
Are you ready for the quiz?
Keep in mind you will have 30 seconds to complete it! [y/n]: y
-----------------------
What is 5+5?
10
What is 7+3?
10
What is 1+1?
2
What is 8+3?
You ran out of time :(
You got  3 / 13  questions correct!
```

### Flags
```
// with cli flags
go run main.go -duration=5 -csvFile=./problems-flag.csv
```
### View CLI Flags
```
// build
go build main.go

// Run bin
./main -h

// output
Usage of ./main:
  -csvFile string
        Relative path to the csv that will power the CLI (default "./problems.csv")
  -duration int
        The time limit for the quiz (default 30)
```