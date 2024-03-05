package main

import "fmt"

func celsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}

func main() {
	celsius := 25.0
	fahrenheit := celsiusToFahrenheit(celsius)
	fmt.Printf("%.2f°C = %.2f°F\n", celsius, fahrenheit)
}

//25.00°C = 77.00°F
