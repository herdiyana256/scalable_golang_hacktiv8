package main

import "fmt"

func main() {
	city := "Jakarta"

	for i := 0; i < len(city); i++ {
		fmt.Printf("Index: %d, byte: %d\n", i, city[i])
	}
}

/*
OUTPUT:
Index: 0, byte: 74
Index: 1, byte: 97
Index: 2, byte: 107
Index: 3, byte: 97
Index: 4, byte: 114
Index: 5, byte: 116
Index: 6, byte: 97


PENJELASAN :
Bisa dilihat, huruf J berubah menjadi 74, huruf a menjadi 97, dan seterusnya. Angka-angka tersebut merupakanrepresentasi dari angka desimal ASCII Code. Kalian dapat mengunjungi link dibawah ini untuk mengetahui lebih detailtentang ASCII Code.https://www.asciitable.com/

*/
