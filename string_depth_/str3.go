package main

import "fmt"

func main() {
	str := "Herdiyan"

	for i, s := range str {
		fmt.Printf("Index => %d, string => %s\n", i, string(s))
	}
}
