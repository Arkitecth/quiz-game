package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)


func main() {
	file, timeLimit := parseArgs()
	if timeLimit == "" {
		startQuiz(file, time.Duration(30))
	}
	convertedTimeLimit, err := strconv.Atoi(timeLimit)
	if err != nil {
		log.Fatal("Invalid time Limit")
	}
	startQuiz(file, time.Duration(convertedTimeLimit))

}

func parseArgs() (string, string) {
	arguments := os.Args
	if len(arguments) < 2 {
		return "", ""
	}
	if arguments[1] == "-h" {
		fmt.Println("Usage of ./quiz:")
		fmt.Println((" -csv string"))
		fmt.Println("	a csv file in the format of 'question,answer' (default 'problems.csv')")
		fmt.Println((" -limit int"))
		fmt.Println("	the time limit for the quiz in seconds' (default 30)")
		os.Exit(0)
	}
	if len(arguments) < 3 {
		fmt.Println("Invalid command")
		os.Exit(-1)
	}
	if arguments[1] == "-csv" {
		return arguments[2], ""
	}
	if arguments[1] == "-limit" {
		return  "", arguments[2]
	}
	return "",""
}
func startQuiz(filename string, interval time.Duration){
	records, err := readCSV(filename)
	if err != nil {
		fmt.Println("An error occurred")
	}
	
	score := 0
	scorePointer := &score
	timeStartScanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Press Enter to Start Quiz")
	timeStartScanner.Scan()
	timer1 := time.NewTimer(interval * time.Second)
	go func(score *int, length int) {
		<-timer1.C
		fmt.Printf("\nYou had a score of %v/%v\n", *score, length)
		os.Exit(0)
	}(scorePointer, len(records))

	for i, record := range records {
		fmt.Printf("Problem #%d %s: ", i+1, record[0])
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		answer := scanner.Text() 
		if answer == record[len(record)-1] {
			score += 1
			*scorePointer = score
		}
	}
	
}

func readCSV(filename string) ([][]string, error) {
	if filename == "" {
		filename = "problems.csv"
	}
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Error while reading the file ", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	return records, nil
}