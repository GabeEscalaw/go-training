// Package weather determines the current weather condition and location depending on the city.
package weather

// CurrentCondition sets up the current condition to be then later added with the location for reporting.
var CurrentCondition string 

// CurrentLocation sets up the current location to be then later added with the condition for reporting.
var CurrentLocation string 

// Forecast takes in the conditions and locations mentioned earlier and displays the location and the weather condition it's currently experiencing is in.
func Forecast(city, condition string) string { 
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
