// Package weather provides tools to calculate weather condition of a location.
package weather

// CurrentCondition to save the weather condition at the moment.
var CurrentCondition string

// CurrentLocation to save the location for which the weather conditions are evaluated.
var CurrentLocation string

// Forecast function returns the current weather condition of a location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
