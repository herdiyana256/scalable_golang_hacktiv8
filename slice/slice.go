package main

import "fmt"

func main() {
	var fruits = []string{"apple", "banana", "mango"}

	_ = fruits

	// SLICE ( MAKE FUNCTION)
	/*
		PENJELASAN :


		Kita juga bisa membuat sebuah slice dengan menggunakanfungsi make. Argumen pertama yang diberikan pada gambarpertama di sebelah kanan adalah tipe dari slice nya, danargumen keduanya adalah panjang dari slice nya.Jika kita jalankan maka hasilnya akan seperti pada gambarkedua di sebelah kanan. Bisa dilihat pada gambar kedua, slicepada variable fruits belum berisi nilai apapun dan variablefruits memiliki tipe data slice of string atau slice dengan tipedata string.Maka dari ikut ketika di print ke terminal, maka hasilnya adalahtiga string kosong.


	*/
	fmt.Printf("%#v", fruits)
	//[]string{"apple", "banana", "mango"}

}

/*



PENJELASAN :
Slice merupakan suatu tipe data yang mirip dengan tipe data array, yang juga memiliki kegunaan untuk menyimpan satu ataulebih data. Namun tipe data slice dan array memiliki sifat yang berbeda. Slice tidak memiliki sifat fixed-length  yang berartipanjang dari slice tidak tetap sehingga kita bisa dengan leluasa menentukan panjang dari slice nya. Slice termasuk dalamkategori reference type yang dimana jika kita melakukan copy terhadap suatu slice, dan kita ubah element dari yang kitacopytersebut, maka slice semulanya juga akan ikut terubah.Cara membuat slice cukup mudah hampir mirip dengan jika kita membuat suatu array. Yang menjadi perbedaan adalah kitatidak perlu menuliskan panjang dari slice nya tidak seperti array. Contohnya seperti pada gambar dibawah ini.





*/
