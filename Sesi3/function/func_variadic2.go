package main

import (
	"fmt"
	"strings"
)

func main() {
	profile("Herdiyan", "Ayam Geprek", "Ayam Rica-rica", "Ayam katsu")

}

func profile(name string, favFood ...string) {
	mergeFavFoofs := strings.Join(favFood, ", ")

	fmt.Println("Hello there!!! I'm", name)
	fmt.Println("I really love to eat", mergeFavFoofs)
}
