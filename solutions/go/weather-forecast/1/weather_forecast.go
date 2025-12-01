// Package weather show the forecase for current location
// and current condition.
package weather

// CurrentCondition is a global variable for current condition.
var CurrentCondition string

// CurrentLocation is a global variable for current location.
var CurrentLocation string

// Forecast returns a string to show the current weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
