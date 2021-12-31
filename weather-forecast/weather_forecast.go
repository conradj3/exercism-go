// Package weather returns current city location and weather status.
package weather

// CurrentCondition indicates the weather status.
var CurrentCondition string

// CurrentLocation indicates the city of which we wish to display the CurrentCondition.
var CurrentLocation string

// Forecast takes an argument of city and condition and returns a string.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
