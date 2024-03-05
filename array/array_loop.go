package main

import (
	"fmt"
	"strings"
)

func main() {
	var fruits = [3]string{"apple", "banana", "mango"}
	// Cara pertama

	for i, v := range fruits {
		fmt.Printf("Index: %d, Value: %s\n", i, v)

	}
	fmt.Println(strings.Repeat("#", 25))

	for i := 0; i < len(fruits); i++ {
		fmt.Printf("Index: %d, Value: %s\n", i, fruits[i])

	}

}

/*

Index: 0, Value: apple
Index: 1, Value: banana
Index: 2, Value: mango
#########################
Index: 0, Value: apple
Index: 1, Value: banana
Index: 2, Value: mango




Ada 2 cara penulisan agar kita dapat melakukan looping untukmengakses element-element yang terdapat pada sebuah array.Perhatikan pada gambar pertama disebelah kanan. Carapertamanya adalah dengan menggunakan range loop.Dengan menggunakan range loop kita bisa dengan mudahmengakses seluruh element pada sebuah array dengan bisa jugasekaligus mendapatkan indexnya.Lalu cara keduanya adalah kita bisa menggunakan looping biasadan kita perlu menggunakan bantuan fungsi len untukmendapatkan panjang dari array nya dengan cara penulisanlen(<nama array>). Jika kita jalankan syntax pada gambarpertama, maka hasilnya akan seperti pada gambar kedua disebelah kanan.

*/
