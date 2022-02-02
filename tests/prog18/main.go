package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
	}
	v, ok := r.Form["compute"]
	fmt.Fprintf(w, "%v %v\n", v, ok)
	var x, y int
	if n, _ := fmt.Sscanf(v[0], "%d*%d", &x, &y); n == 2 {
		fmt.Fprintf(w, "%d * %d = %d\n", x, y, x*y)
	}

}