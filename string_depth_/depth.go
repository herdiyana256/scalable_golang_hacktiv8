// Kali ini kita akan membahas lebih jauh tentang tipe data string pada Go.Tipe data string pada Go terbentuk dari kumpulan tipe data byte yang di letakkan di dalam slice atau bisa kita sebutdengan slice of bytes.Tipe data byte pada Go merupakan tipe data alias dari tipe data uint8.Ketika kita melakukan indexing terhadap suatu string, maka kita akan mendapat nilai representasi dari byte nya.

package main

import "fmt"

func main() {
	// Declaring a string variable
	var str string = "Hello, Everyone!"

	// Printing the string
	fmt.Println(str)

	// Getting the length of the string
	fmt.Println("Length:", len(str))

	// Accessing individual characters in the string
	fmt.Println("First character:", string(str[0]))

	// Concatenating strings
	str2 := "Welcome"
	str3 := " to myWorld!"
	result := str2 + str3
	fmt.Println(result)

	// String formatting
	name := "Herdiyan Adam Putra"
	age := 28
	fmt.Printf("My name is %s and I am %d years old.\n", name, age)

	// String comparison
	if str == "Hello, Soyoung!" {
		fmt.Println("Strings are equal")
	} else {
		fmt.Println("Strings are not equal")
	}

	// Iterating over characters in a string
	for i, char := range str {
		fmt.Printf("Character at index %d: %c\n", i, char)
	}
}

/*

Index: 0, byte: 74
Index: 1, byte: 97
Index: 2, byte: 107
Index: 3, byte: 97
Index: 4, byte: 114
Index: 5, byte: 116
Index: 6, byte: 97


*/
