package main

import "fmt"

func main() {
	var name string
	var age int = 28

	name = 30
	age = "Herdiyan"

	fmt.Println("Ini adalah namanya ==>", name)
	fmt.Println("Ini adalah umurnya ==>", age)
}

/*
Variable with data type
Nama jika kita tidak memberikan nilai dengan tipe data yang sama pada variable name dan age, maka akan terjadi error seperti pada gambar disebelah kanan yang dimana variable name kita re-assign dengan sebuah nulai dengan tipe int sedangkan variable age kita re-assign dengan sebuah nilai dengan tipe data string.
*/
