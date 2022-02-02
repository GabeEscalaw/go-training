package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
	"sync"
	"time"
)

type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

func (c *SafeCounter) Increment(word string) {
	c.mu.Lock()
	c.v[word]++
	c.mu.Unlock()
}

func (c *SafeCounter) Value(word string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v[word]
}

func countWords(wordsArray []string) {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < len(wordsArray); i++ {
		go c.Increment(wordsArray[i])
	}
}

func fileOpener(fileName string) []string {
	file, err := os.Open("" + fileName)
	if err != nil {
		log.Fatalf("%v %v", err, fileName)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	wordsArray := make([]string, 0)
	reg, err := regexp.Compile("(?m)[^a-z]")
	if err != nil {
		log.Fatal(err)
	}
	for scanner.Scan() {
		tempArr := strings.Fields(scanner.Text())
		for _, notCleanedWord := range tempArr {
			wordsArray = append(wordsArray, strings.ToLower(reg.ReplaceAllString(notCleanedWord, "")))
		}
	}

	return wordsArray
}

func dedupeAndSort(currentArray []string, addArray []string, c chan []string) {
	currentArray = append(currentArray, addArray...)
	keys := make(map[string]bool)
	cleanArray := []string{}

	for _, entry := range currentArray {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			cleanArray = append(cleanArray, entry)
		}
	}
	sort.Strings(cleanArray)
	c <- cleanArray
}

func main() {
	filesInput := os.Args[2:]
	c := SafeCounter{v: make(map[string]int)}
	ch := make(chan []string)
	uniqueWords := make([]string, 0)
	for _, fileName := range filesInput {
		go countWords(fileOpener(fileName))
		go dedupeAndSort(uniqueWords, fileOpener(fileName), ch)
	}
	uniqueWords = append(uniqueWords, <-ch...)
	time.Sleep(time.Second)
	for _, word := range uniqueWords {
		fmt.Println(word+" ", c.Value(word))
	}
}
