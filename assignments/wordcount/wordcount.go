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

// SafeCounter is done to safely access data for concurrent routines.
type SafeCounter struct {
	mu sync.Mutex
	wordMap map[string]int
}

// Inc safely increments the count for that word in the map.
func (sc *SafeCounter) Inc(key string) {
	sc.mu.Lock()
	sc.wordMap[key]++
	sc.mu.Unlock()
}

// Value safely returns the word in the map.
func (sc *SafeCounter) Value(key string) int {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	return sc.wordMap[key]
}

// sortAndRemoveDuplicates sorts the wordsList alphabetically and removes duplicated words from the slice.
func sortAndRemoveDuplicates (wordsList []string, ch chan []string)  {
	sort.Strings(wordsList)
	
	for i := 0; i < len(wordsList)-1; i++ {
		if len(wordsList) == 1 {
			break
		} else if wordsList[0] == wordsList[1] {
			wordsList = append(wordsList[:i], wordsList[i+1:]...)
			i = 0
		} else if wordsList[i] == wordsList[i+1] {
			wordsList = append(wordsList[:i], wordsList[i+1:]...)
			i = 0
		}
	}

	ch <- wordsList
	time.Sleep(2*time.Second)
	defer close(ch)
}

// wordCleaner removes all non-letter characters and turns them into lowercase.
func wordCleaner(word string) string {
	re, err := regexp.Compile(`[^\w]`)
	if err != nil {
		log.Fatal(err)
	}

	word = re.ReplaceAllString(word, "")
	word = strings.ToLower(word)
	
	return word
}

// wordCounter increments the count of a word in a map.
func(sc *SafeCounter) wordCounter(words []string)  {
	for i := 0; i < len(words); i++ {
		go sc.Inc(words[i])	
	}
}

// wordExtractor opens the file and extracts each word then puts them into the slice "wordsList".
func wordExtractor (fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	wordsList := make([]string, 0)
	if err != nil {
		log.Fatal(err)
	}

	for scanner.Scan() {
		tempArr := strings.Fields(scanner.Text())
		for _, dirtyWords := range tempArr {
			if wordCleaner(dirtyWords) != "" {
				wordsList = append(wordsList, wordCleaner(dirtyWords))
			}
		}
	}

	return wordsList
}

// finalSortAndDuplicateRemover sorts and removes the duplicates for the final array.
func finalSortAndDuplicateRemover (wordsList []string) []string {
	sort.Strings(wordsList)

	for i := 0; i < len(wordsList)-1; i++ {
		if len(wordsList) == 1 {
			break
		} else if wordsList[0] == wordsList[1] {
			wordsList = append(wordsList[:i], wordsList[i+1:]...)
			i = 0
		} else if wordsList[i] == wordsList[i+1] {
			wordsList = append(wordsList[:i], wordsList[i+1:]...)
			i = 0
		}
	}

	return wordsList
}

// main takes in all the words from the input files and counts the unique non special character ones.
func main() {
	fileNames := os.Args[1:] 
	sc := SafeCounter{wordMap: make(map[string]int)}
	ch := make(chan []string)
	wordList := make([]string, 0)
	for _, fileName := range fileNames {
		go sc.wordCounter(wordExtractor(fileName))
		go sortAndRemoveDuplicates(wordExtractor(fileName), ch)
		wordList = append(wordList, <- ch...)
	}			

	time.Sleep(time.Second)
	for _, word := range finalSortAndDuplicateRemover(wordList) {
		fmt.Printf("%v %v\n", word, sc.Value(word))
	}
}