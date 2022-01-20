package main

import "fmt"

func main() {
	reply := map[string]string {"Hajj":"Hajj-e-Akbar","Umrah":"Hajj-e-Asghar"}
	mesg := ""
	count := 0
	_, err := fmt.Scan(&mesg)
	for err == nil {
		count++
		if word, ok := reply[mesg]; ok{
			fmt.Printf("Case %d: %s\n", count, word)
		} else if mesg == "*" {
			break
		} 
		_, err = fmt.Scan(&mesg)
	}
}