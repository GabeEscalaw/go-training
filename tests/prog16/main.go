package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// // Get whats inside cobena website
// func main() {
// 	url := "https://cobenagroup.com"
// 	getData(url)
// }
// func getData(url string) {
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		log.Fatalf("%v", err)
// 	}
// 	b, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatalf("%v", err)
// 	}
// 	fmt.Printf("%s", b)
// }

// Get whats inside cobena website
func main() {
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go getData(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
}
func getData(url string, ch chan string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("%v", err)
		return
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("%v", err)
		return
	}
	duration := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %v chars %v", duration, len(b), url)
}

// test by typing:
// go run main.go http://cobenagroup.com http://golang.org http://google.com http://homerpagka.com