// Package weather allows you to query information about the weather.
package weather

// CurrentCondition describes what the weather is like.
var CurrentCondition string
// CurrentLocation describes the place that CurrentCondition is referencing.
var CurrentLocation string


// Forecast will update the current weather status and return a helpful description.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
