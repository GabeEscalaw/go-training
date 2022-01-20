package main

import "fmt"

func main() {
	country := map[string]string {"HELLO":"ENGLISH","HOLA":"SPANISH","HALLO":"GERMAN","BONJOUR":"FRENCH","CIAO":"ITALIAN","ZDRAVSTVUJTE":"RUSSIAN"}
	mesg := ""
	count := 0
	_, err := fmt.Scan(&mesg)
	for err == nil {
		count++
		if lang, ok := country[mesg]; ok{
			fmt.Printf("Case %d: %s\n", count, lang)
		} else if mesg == "#" {
			break
		} else {
			fmt.Printf("Case %d: UNKNOWN\n", count)
		}
		_, err = fmt.Scan(&mesg)
	}
}