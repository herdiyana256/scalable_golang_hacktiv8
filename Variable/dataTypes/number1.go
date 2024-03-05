package main

import "fmt"

func main() {
	var first uint8 = 96
	var second int8 = -12

	fmt.Printf("tipe data first: %T\n", first)
	fmt.Printf("tipe data second: %T\n", second)

}

/*
tipe data first: uint8
tipe data second: int8
*/

/*
PENJELASAN :

- float
Float merupakan tipe data numerik desimal pada bahasa Go.
Secara umum float dibagi menjadi 2 jenis yaitu float32, dan float34


Perbedaan kedua jenis float tersebut berada di lebar cakupan nilai desimail yang bisa di tampung .
Untuk lebih jelasnya bisa merujuk ke spesifikasi :iIEEE-754 32-bitfloating-point numbershttps://www.geeksforgeeks.org/data-types-in-go/


*/
