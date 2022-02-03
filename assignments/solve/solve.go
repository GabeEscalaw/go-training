package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		input := r.URL.Query()["coef"]

		var fixedInput string 
		fixedInput = input[0]
		splitInput := strings.Split(fixedInput, ",")
		
		coefficients := make([]float64, 0)
		for _, i := range splitInput {
			temp, err := strconv.ParseFloat(i, 32)
			if err != nil {
				panic(err)
			}
			coefficients = append(coefficients, temp)
		}
	
		cramersRule(coefficients, w)
	}
}

// cramersRule performs an easy-to-visualize Cramer's Rule for 3x3 systems to solve for x, y, and z
func cramersRule (coefficients []float64, w http.ResponseWriter) {
	a1 := coefficients[0]
	b1 := coefficients[1]
	c1 := coefficients[2]
	d1 := coefficients[3]
	a2 := coefficients[4]
	b2 := coefficients[5]
	c2 := coefficients[6]
	d2 := coefficients[7]
	a3 := coefficients[8]
	b3 := coefficients[9]
	c3 := coefficients[10]
	d3 := coefficients[11]

	d := a1*b2*c3 + b1*c2*a3 + c1*a2*b3 - a3*b2*c1 - b3*c2*a1 - c3*a2*b1
	dx := (d1*(b2*c3-b3*c2) - b1*(d2*c3-d3*c2) + c1*(d2*b3-d3*b2)) / d
	dy := (a1*(d2*c3-d3*c2) - d1*(a2*c3-a3*c2) + c1*(a2*d3-a3*d2)) / d
	dz := (a1*(b2*d3-b3*d2) - b1*(a2*d3-a3*d2) + d1*(a2*b3-a3*b2)) / d

	fmt.Fprintf(w, "system:\n%vx + %vy + %vz = %v\n", a1, b1, c1, d1)
	fmt.Fprintf(w, "%vx + %vy + %vz = %v\n",  a2, b2, c2, d2)
	fmt.Fprintf(w, "%vx + %vy + %vz = %v\n\n",  a3, b3, c3, d3)

	fmt.Fprintf(w, "solution:\nx = %.2f, y = %.2f, z = %.2f\n", dx, dy, dz)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

// 4,5,6,7,2,3,1,2,1,2,3,2