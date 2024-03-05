package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{10, 5, 8, 20, 3}
	sort.Ints(numbers)
	max := numbers[len(numbers)-1]
	fmt.Println("Nilai maksimum:", max)
}

//Nilai maksimum: 20
