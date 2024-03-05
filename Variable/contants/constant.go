package main

import "fmt"

func main() {
	const full_name string = "Herdiyan Adam Putra"

	fmt.Printf("Hello %s", full_name)

}

//Hello Herdiyan Adam Putra


/*
Constant (const) atau konstanta merupakan jenis variable pada bahasa Go yang nilainya tidak dapat diubah .
Contohnya jika kita memiliki nilai-nilai tetap seperti PI, kecepatan cahaya, luas lingkaran dan lain lain yang erupakan nilai nilai tetap.

Maka kita bisa menyimpan nilai nilai tersebut dengan variable const. seperto yang kita lihat pada gambar pertama dibawah kita mempunyai suatu variable const yang memiliki tipe data string.
Perlu di ingat disini bahwa ketika kita membuat variable dengan const, maka kita harus langsung memberikan nilai kepadanya. karena jika tidak maka akan timbul error pada saat compile time seperti pada gambar kedua dibawah.


*/

package main 


import "fmt"

func main () {
		const full_name string

		fmt.Println(full_name)
}