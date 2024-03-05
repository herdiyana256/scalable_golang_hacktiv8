package main

import "fmt"

func main() {
	var integer int
	var float float64
	var boolean bool
	var stringVar string

	fmt.Println("Nilai-nilai awal: ")
	fmt.Println("Integer:", integer)
	fmt.Println("Float:", float)
	fmt.Println("Boolean:", boolean)
	fmt.Println("String:", stringVar)

	// Mengubah nilai variabel

	integer = 10
	float = 3.14
	boolean = true
	stringVar = "Hello, World!"

	fmt.Println("\nNilai-nilai setelah diubah")
	fmt.Println("Integer:", integer)
	fmt.Println("Float:", float)
	fmt.Println("Boolean:", boolean)
	fmt.Println("String:", stringVar)
}
/*
Nilai-nilai awal: 
Integer: 0
Float: 0
Boolean: false
String:

Nilai-nilai setelah diubah
Integer: 10
Float: 3.14
Boolean: true
String: Hello, World!

8/


// Penjelasan : Program ini mendeklarasikan beberapa variable dengan tipe data yang berbeda dan kemudian mencetak nilai-nilai awalnya karena variabel tersebut tidak di inisiasi , mereka akan memiliki nilai nol sesuai dengan tipe datanya ( integer:0, float:0.0, boolean: false, string: ""). setelah itu, program mengubah nilai nilai variable tersebut dan mencetaknya lagi .