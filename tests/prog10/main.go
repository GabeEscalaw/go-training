package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

// --- READ ---
// func main() {
// 	fp, _ := os.Open("sample.csv")
// 	defer fp.Close()
// 	r := csv.NewReader(fp)
// 	lines, _ := r.ReadAll()
// 	fmt.Println(lines)
// }

// --- WRITE ---
func main() {
	fp, err := os.Open("sample.csv")
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer fp.Close()
	r := csv.NewReader(fp)
	lines, err := r.ReadAll()
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Println(lines)

	fp2, _ := os.Create("output.csv")
	w := csv.NewWriter(fp2)
	for _, row := range lines {
		_ = w.Write(row)
	}
	w.Flush()
	fp2.Close()
}