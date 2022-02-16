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

// Each command corresponds to one of these URLs
const loginURL = "http://localhost:8080/login"
const logoutURL = "http://localhost:8080/logout"
const registerURL = "http://localhost:8080/signup"
const deleteAccURL = "http://localhost:8080/delete"

const listURL = "http://localhost:8080/list"
const addURL = "http://localhost:8080/add"
const removeURL = "http://localhost:8080/remove"
const markURL = "http://localhost:8080/mark"

const errURL = "http://localhost:8080/"


// main transfers all of the flagged inputs of the user into their respective functions based on the command specified.
func main() {
	cmd := flag.String("cmd", "", "=in \t-u=username -p=password\n\t\t> log in with the specified credentials\n=out\n\t\t> log out of the current account\n=reg \t-u=username -p=password\n\t\t> register a user to the database with the specified credentials\n=del \t-u=username -p=password\n\t\t> delete a user from the database with the specified credentials\n=list \n\t\t> display all of the logged in user's watches and their respective details\n=addW \t-brand=brand -model=model -width=width -price=price\n\t\t> add a watch into the currently logged in user with the specified details (all fields are required)\n=delW \t-brand=brand -model=model\n\t\t> delete a watch from the user's watch collection with the specified brand and model\n=mark \t-brand=brand -model=model\n\t\t> toggle a watch's Collected boolean with the specified brand and model")
	user := flag.String("u", "", "username input")
	pass := flag.String("p", "", "password input")
	brand := flag.String("brand", "", "Brand input for watch's details")
	model := flag.String("model", "", "Model input for watch's details")
	width := flag.String("width", "", "Dial Size input for watch details")
	price := flag.String("price", "", "Price input for watch's details")
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
		addWatch(*brand, *model, *width, *price)
	case "delW":
		deleteWatch(*brand, *model)
	case "mark":
		markWatch(*brand, *model)
	default :
		err(*cmd)
	}
}

// login throws the username and password json information specified by the user as a login Post request to be received and processed by server.go.
func login(username string, password string) {
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	jsonRec := fmt.Sprintf("{\"username\":\"%s\",\"password\":\"%s\"}", username, password)
	outData := bytes.NewBuffer([]byte(jsonRec))
	resp, err := c.Post(loginURL, "application/json", outData)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
}

// logout throws a logout Post request to be received and processed by server.go.
func logout() {
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	jsonRec := fmt.Sprintf("{\"username\":\"%s\",\"password\":\"%s\"}", "na", "na")
	outData := bytes.NewBuffer([]byte(jsonRec))
	resp, err := c.Post(logoutURL, "application/json", outData)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
}

// register throws the username and password json information specified by the user as a signup Post request to be received and processed by server.go.
func register(username string, password string) {
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	jsonRec := fmt.Sprintf("{\"username\":\"%s\",\"password\":\"%s\"}", username, password)
	outData := bytes.NewBuffer([]byte(jsonRec))
	resp, err := c.Post(registerURL, "application/json", outData)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
}

// deleteAccount throws the username and password json information specified by the user as a delete Post request to be received and processed by server.go.
func deleteAccount(username string, password string) {
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	jsonRec := fmt.Sprintf("{\"username\":\"%s\",\"password\":\"%s\"}", username, password)
	outData := bytes.NewBuffer([]byte(jsonRec))
	resp, err := c.Post(deleteAccURL, "application/json", outData)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
}

// watchList throws a Get request to be received and processed by server.go.
func watchList() {
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	resp, err := c.Get(listURL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
}

// addWatch throws the brand, model, width, and price json information specified by the user as an add Post request to be received and processed by server.go.
func addWatch(brand string, model string, width string, price string) {
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	jsonRec := fmt.Sprintf("{\"brand\":\"%s\",\"model\":\"%s\",\"width\":\"%s\",\"price\":\"%s\"}", brand, model, width, price)
	outData := bytes.NewBuffer([]byte(jsonRec))
	resp, err := c.Post(addURL, "application/json", outData)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
}

// deleteWatch throws the brand and model json information specified by the user as a removeWatch Post request to be received and processed by server.go.
func deleteWatch(brand string, model string) {
	c := http.Client{Timeout: time.Duration(1) * time.Second * 3}
	jsonRec := fmt.Sprintf("{\"brand\":\"%s\",\"model\":\"%s\"}", brand, model)
	outData := bytes.NewBuffer([]byte(jsonRec))
	resp, err := c.Post(removeURL, "application/json", outData)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
}

// markWatch throws the brand and model json information specified by the user as a markWatch Post request to be received and processed by server.go.
func markWatch(brand string, model string) {
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	jsonRec := fmt.Sprintf("{\"brand\":\"%s\",\"model\":\"%s\"}", brand, model)
	outData := bytes.NewBuffer([]byte(jsonRec))
	resp, err := c.Post(markURL, "application/json", outData)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
}

// err throws a URL specified by the user as a Get request to be received and displayed by server.go as a bad request.
func err(cmd string) {
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	newURL := fmt.Sprintf(errURL+cmd)
	resp, err := c.Get(newURL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
}