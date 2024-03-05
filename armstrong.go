package main

import (
	"fmt"
	"math"
)

func isArmstrongNumber(num int) bool {
	temp, sum := num, 0
	digits := int(math.Log10(float64(num))) + 1
	for temp != 0 {
		digit := temp % 10
		sum += int(math.Pow(float64(digit), float64(digits)))
		temp /= 10
	}
	return sum == num
}

func main() {
	num := 153
	if isArmstrongNumber(num) {
		fmt.Println(num, "adalah bilangan Armstrong")
	} else {
		fmt.Println(num, "bukan bilangan Armstrong")
	}
}
