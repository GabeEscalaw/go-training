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

const baseURL = "http://localhost:8080/user"
const logoutURL = "http://localhost:8080/logout"
const registerURL = "http://localhost:8080/signup"
const deleteAccURL = "http://localhost:8080/delete"

const listURL = "http://localhost:8080/list"
const addURL = "http://localhost:8080/add"
const removeURL = "http://localhost:8080/remove"
const markURL = "http://localhost:8080/mark"

func main() {
	cmd := flag.String("cmd", "", "login, logout, register, delete")
	user := flag.String("u", "", "username input")
	pass := flag.String("p", "", "password input")
	brand := flag.String("brand", "", "login, logout, register, delete")
	model := flag.String("model", "", "username input")
	dialSize := flag.String("dialSize", "", "login, logout, register, delete")
	price := flag.String("price", "", "password input")
	flag.Parse()
	switch *cmd {
	case "in":
		login(*user, *pass)
	case "out":
		logout()
	case "reg":
		register(*user, *pass)
	case "del":
		deleteAccount(*user, *pass)

	case "list":
		watchList()
	case "addW":
		addWatch(*brand, *model, *dialSize, *price)
	case "delW":
		deleteWatch(*brand, *model)
	case "markW":
		markWatch(*brand, *model)
	}
}

func login(username string, password string) {
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	jsonRec := fmt.Sprintf("{\"username\":\"%s\",\"password\":\"%s\"}", username, password)
	outData := bytes.NewBuffer([]byte(jsonRec))
	resp, err := c.Post(baseURL, "application/json", outData)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
}

func logout() {
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	jsonRec := fmt.Sprintf("{\"username\":\"%s\",\"password\":\"%s\"}", "na", "na")
	outData := bytes.NewBuffer([]byte(jsonRec))
	resp, err := c.Post(logoutURL, "application/json", outData)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
}

func register(username string, password string) {
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	jsonRec := fmt.Sprintf("{\"username\":\"%s\",\"password\":\"%s\"}", username, password)
	outData := bytes.NewBuffer([]byte(jsonRec))
	resp, err := c.Post(registerURL, "application/json", outData)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
}

func deleteAccount(username string, password string) {
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	jsonRec := fmt.Sprintf("{\"username\":\"%s\",\"password\":\"%s\"}", username, password)
	outData := bytes.NewBuffer([]byte(jsonRec))
	resp, err := c.Post(deleteAccURL, "application/json", outData)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
}

func watchList() {
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	resp, err := c.Get(listURL)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
}

func addWatch(brand string, model string, dialSize string, price string) {
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	jsonRec := fmt.Sprintf("{\"brand\":\"%s\",\"model\":\"%s\",\"dialSize\":\"%s\",\"price\":\"%s\"}", brand, model, dialSize, price)
	outData := bytes.NewBuffer([]byte(jsonRec))
	resp, err := c.Post(addURL, "application/json", outData)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
}

func deleteWatch(brand string, model string) {
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	jsonRec := fmt.Sprintf("{\"brand\":\"%s\",\"model\":\"%s\"}", brand, model)
	outData := bytes.NewBuffer([]byte(jsonRec))
	resp, err := c.Post(removeURL, "application/json", outData)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
}

func markWatch(brand string, model string) {
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	jsonRec := fmt.Sprintf("{\"brand\":\"%s\",\"model\":\"%s\"}", brand, model)
	outData := bytes.NewBuffer([]byte(jsonRec))
	resp, err := c.Post(markURL, "application/json", outData)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
}