package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type User struct {
	Name string `json:"name"`
	Job string `json:"job"`
}

func main() {
	fp, _ := os.Open("sample2.json")
	defer fp.Close()
	byteData, _ := ioutil.ReadAll(fp)
	var u []User
	json.Unmarshal([]byte(byteData), &u)
	fmt.Println(u)
}