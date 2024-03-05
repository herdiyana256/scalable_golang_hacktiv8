package main

import (
	"fmt"
	"math"
)

func circleProperties(radius float64) (float64, float64) {
	area := math.Pi * math.Pow(radius, 2)
	circumference := 2 * math.Pi * radius
	return area, circumference
}

func main() {
	radius := 5.0
	area, circumference := circleProperties(radius)
	fmt.Println("Luas Lingkaran:", area)
	fmt.Println("Keliling Lingkaran:", circumference)
}
