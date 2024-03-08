// Pointer (Pointer as a parameter)Pointer juga bisa dijadikan sebagai sebuah parameter pada sebuah function. Contohnya seperti pada gambar pertama di sebelah kanan.Awalnya variable a memiliki nilai 10. Setelah itu kita memberikan variable a kepada function changeValue yang menerima satu parameter dengan tipe data pointer int. Kemudian nilai dari variable a diganti di dalam function changeValue melalui parameter yang diterimanya. Ketika dijalankan maka hasilnya akan seperti pada gambar kedua.

package main

import "fmt"

func main() {
	var a int = 10

	fmt.Println("Before:", a) // Mencetak nilai variabel a sebelum perubahan

	changeValue(&a) // Memanggil fungsi changeValue dengan memberikan alamat memori dari variabel a

	fmt.Println("After:", a) // Mencetak nilai variabel a setelah perubahan
}

func changeValue(number *int) {
	*number = 20 // Mengubah nilai yang ditunjuk oleh pointer number menjadi 20
}

/*
}
Ketika program dijalankan:

Variabel a diinisialisasi dengan nilai 10.
Nilai variabel a dicetak sebelum perubahan menggunakan fmt.Println("Before:", a).
Kemudian, fungsi changeValue() dipanggil dengan memberikan alamat memori dari variabel a sebagai argumen (&a).
Di dalam fungsi changeValue(), nilai yang ditunjuk oleh pointer number (yang merupakan alamat memori dari a) diubah menjadi 20 dengan menggunakan operator dereference (*number = 20).
Setelah panggilan fungsi selesai, nilai variabel a berubah menjadi 20.
Nilai variabel a dicetak kembali setelah perubahan menggunakan fmt.Println("After:", a).
Hasilnya akan mencetak:

Before: 10
After: 20

Ini menunjukkan bahwa nilai variabel a telah berubah setelah dipanggil oleh fungsi changeValue() melalui penggunaan pointer. Dengan menggunakan pointer, kita dapat mengakses dan mengubah nilai variabel yang ada di luar fungsi tersebut.






*/
