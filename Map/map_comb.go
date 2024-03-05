package main

import "fmt"

func main() {

	var people = []map[string]string{

		{"name": "Herdiyan", "age": "28"},
		{"name": "Helmi", "age": "20"},
		{"name": "Mimi", "age": "42"},
	}

	for i, person := range people {
		fmt.Printf("Index: %d, name: %s, age: %s\n", i, person["name"], person["age"])
	}

}

/*

Index: 0, name: Herdiyan, age: 28
Index: 1, name: Helmi, age: 20
Index: 2, name: Mimi, age: 42



PENJELASAN MATERI
Map (Combining slice with map)Kita juga dapat mengkombinasikan slice dengan map.Contohnya seperti pada gambar pertama di sebelah kanan.Kita juga dapat me-looping nya dengan menggunakan rangeloop.Jika kita jalankan pada terminal kita, maka hasilnya akanseperti pada gambar kedua.


PENJELASAN KODE :

Pada baris 6-10, kita mendefinisikan variabel people yang merupakan sebuah slice dari map. Setiap elemen slice people adalah sebuah map yang memiliki kunci (key) "name" dan "age" dengan nilai (value) yang sesuai. Ini berarti kita memiliki beberapa map, masing-masing mewakili seseorang dengan nama dan usia mereka.

Kemudian, pada baris 12-15, dilakukan iterasi terhadap slice people menggunakan range. Dalam setiap iterasi, kita mendapatkan indeks (i) dan masing-masing elemen map (person) dari slice people. Di dalam loop, kita mencetak indeks, nama, dan usia dari setiap orang dengan menggunakan fungsi fmt.Printf().

Hasilnya adalah mencetak setiap orang beserta nama dan usia mereka dalam format yang telah ditentukan.

Jadi, program ini adalah contoh penggunaan slice dan map di Go untuk menyimpan dan mengelola data yang terstruktur.

*/
