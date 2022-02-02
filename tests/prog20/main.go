package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const baseURL = "http://localhost:8080/store"

func main() {
	cmd := flag.String("cmd", "", "add, list, delete")
	data := flag.String("data", "", "data to be added")
	id := flag.Int("id", -1, "ID of the record to delete")
	flag.Parse()
	switch *cmd {
	case "add":
		addItem(*data)
	case "list":
		listItems()
	case "delete":
		deleteItem(*id)
	}
}

func addItem(inData string) {
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	jsonRec := fmt.Sprintf("{\"name\":\"%s\"}", inData)
	outData := bytes.NewBuffer([]byte(jsonRec))
	resp, err := c.Post(baseURL, "application/json", outData)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(resp.Body)	
	fmt.Println(string(data))
}

func listItems() {
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	resp, err := c.Get(baseURL)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
}

func deleteItem(id int) {
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	url := fmt.Sprintf("%s/%d", baseURL, id)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := c.Do(req)
	if err != nil { 
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(resp.Body)	
	fmt.Println(string(data))
} 