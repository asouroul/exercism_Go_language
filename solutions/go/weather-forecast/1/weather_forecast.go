// Package weather provides tools to display the weather condition for a specific city.
package weather

var (
    // CurrentCondition represents a certain weather condition.
	CurrentCondition string
    // CurrentLocation represents a certain location.
	CurrentLocation  string
)
// Forecast returns a text displaying the weather condition for a specific city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
