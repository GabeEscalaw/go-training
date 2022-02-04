package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// handler provides the response writer and requests for the process function if the URL path is correct.
func handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/solve" {
			process(w, r)
		} else {
			message := "404. That's an error.\nThe requested URL " + r.URL.Path + " was not found on this server. That's all we know."
			http.Error(w, message, http.StatusBadRequest)
		}
	}
}

// process extracts the inputs from the URL into the proper values for the cramersRule function
func process(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	for parameter := range values {
		if parameter == "coef" {
			input := r.URL.Query()["coef"]
	
			var fixedInput string 
			fixedInput = input[0]
			splitInput := strings.Split(fixedInput, ",")
					
			coefficients := make([]float64, 0)
			if len(splitInput) == 12 {
				for _, i := range splitInput {
					temp, err := strconv.ParseFloat(i, 32)
					if err != nil {
						fmt.Fprintf(w, "Error Found: %v\nReplacing that value with 0.\n\n", err)
					} 
					coefficients = append(coefficients, temp)
				}
					
				cramersRule(coefficients, w)
			} else {
				fmt.Fprintf(w, "Incorrect number of coefficients")
			}
		} else {
			fmt.Fprintf(w, "Error Found: Parameter ?%v was not found.", parameter)
		}
	}
}

// cramersRule performs an easy-to-visualize Cramer's Rule for 3x3 systems to solve for x, y, and z.
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

	if d == 0 {
		if dx == 0 && dy == 0 {
			fmt.Fprintf(w, "solution:\ndependent -with multiple solutions")
		}
		if dx == 0 || dy == 0 {
			fmt.Fprintf(w, "solution:\ninconsistent - no solution")
		}
	} else {
		fmt.Fprintf(w, "solution:\nx = %.2f, y = %.2f, z = %.2f\n", dx, dy, dz)
	}
	
}

// main handles the http functions
func main() {
	http.HandleFunc("/", handler())
	http.ListenAndServe(":8080", nil)
}