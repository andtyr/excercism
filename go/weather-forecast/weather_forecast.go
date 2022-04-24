// Package weather contains elements to determine weather conditions by city.
package weather

// CurrentCondition describes the current weather.
var CurrentCondition string

// CurrentLocation describes the city wherein weather is happening.
var CurrentLocation string

// Forecast outputs string of weather prediction for chosen location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
