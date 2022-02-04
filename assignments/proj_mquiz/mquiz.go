package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// askQuestions loops through -n questions for the user and displays their score after.
func askQuestions(numQuestions int, lines [][]string) {
	var score int

	for i := 0; i < numQuestions; i++ {
		var ans string 
		rand.Seed(time.Now().UnixNano())
		index := rand.Intn(numQuestions - 0) + 0
		fmt.Printf("Q: %s = ", lines[index][0])
		fmt.Scan(&ans)
		if ans == lines[index][1] {
			fmt.Println("Correct!")
			score++
		}
	}

	fmt.Println("You answered", score, "out of", numQuestions, "questions correctly.")
}

// main reads the files, handles errors, and proceeds with the askQuestions function if there are no errors.
func main() {
	// flag.Type(flag, default value, short description)
	dbname := flag.String("csv", "questions.csv", "database")
	numQuestions := flag.Int("n", 10, "number of questions")
	flag.Parse()
	fp, err := os.Open(*dbname) // fp = file pointer
	if err != nil {
		log.Fatalf("%v", err) // If there's an error, define it
	}
	if filepath.Ext(strings.TrimSpace(*dbname)) != ".csv" {
		log.Fatalf("Incorrect database format. Database should be in .csv format")
	}
	defer fp.Close()
	r := csv.NewReader(fp)
	lines, _ := r.ReadAll()
	if len(lines) < *numQuestions {
		log.Fatalf("Insufficient questions. Database should contain at least -n amount of questions")
	}

	askQuestions(*numQuestions, lines)
}