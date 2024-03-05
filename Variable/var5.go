package main

import "fmt"

func main() {
	var student1, student2, student3 string = "satu", "dua", "tiga"

	var first, second, third int

	first, second, third = 1, 2, 3

	fmt.Println(student1, student2, student3)

	fmt.Println(first, second, third)

}

/*
satu dua tiga
1 2 3
*/

/*
Multiple variable declarations

Pada bahasa golang, kita dapat membuat lebih dari 1 variable pada satu baris/ line yang sama . Caranya seperti gambar pertama .
Jika kita perhatikan, kita telah membuat 3 variable pada baris pertama dan membuat 3 variable pada baris kedua. pada baris pertama, ketiga variable tersebut memiliki tipe data string dan kita langsung memberikan nilai kepadanya. sedangkan pada baris kedua, ketiga variable tersebut memiliki tipe data int dan kita memberikan nilai kepadanya dengan cara me-reassign nilai-nilai keada ketiga variavle pada baris kedua tersebut.


dan ketika kita halankan pada terminal kita, maka hasilnya akan sseperti pada gambar kedua
*/
