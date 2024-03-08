// CLosure - Closure as a return value

package main

import (
	"fmt"
	"strings"
)

func main() {
	var studentLists = []string{"Herdiyan", "Bagast", "Reza", "Dahlia", "Yudho"}

	var find = findStudent(studentLists)

	fmt.Println(find("herdiyan"))
}

func findStudent(students []string) func(string) string {

	return func(s string) string {
		var student string
		var position int

		for i, v := range students {
			if strings.ToLower(v) == strings.ToLower(s) {
				student = v
				position = i
				break
			}
		}

		if student == "" {
			return fmt.Sprintf("%s does'nt exist!!!.", s)

		}
		return fmt.Sprintf("We found %s at position %d", s, position+1)
	} // Closure (Closure as a return value)Closure juga bisa dijadikan sebagai nilai kembalian dari suatu function. Contohnya seperti pada gambar sebelah kanan. Function findStudent pada line 66 merupakan sebuah function yang menerima satu parameter dengan tipe data slice of stringdan me-return sebuah yang closure. Closure yang direturn menerima satu parameter dengan tipe data string yang digunakan untuk mencari satu data student dari parameter yang diterima oleh function findStudent. Jika closure yang direturn menemukan data student nya, maka kalimat yang ada pada line 82 akan di return, jika tidak maka kalimat pada lone 80 akan di return.Perhatikan juga pada line 73, kita menggunakan function strings.ToLower untuk mengecilkan huruf.Function strings.ToLower  berasal dari package strings. Kemudian pada line 60, variable bernama find akan menampung nilai return dari function findStudent. Karena function findStudentme-return sebuah closure, maka untuk memanggil closure yang di return oleh function findStudent, kita perlu memanggil variable find sekaligus memberikan parameternya.

}

/*
Closure (Closure as a return value)Closure juga bisa dijadikan sebagai nilai kembalian dari suatu function. Contohnya seperti pada gambar sebelah kanan. Function findStudent pada line 66 merupakan sebuah function yang menerima satu parameter dengan tipe data slice of stringdan me-return sebuah yang closure. Closure yang direturn menerima satu parameter dengan tipe data string yang digunakan untuk mencari satu data student dari parameter yang diterima oleh function findStudent. Jika closure yang direturn menemukan data student nya, maka kalimat yang ada pada line 82 akan di return, jika tidak maka kalimat pada lone 80 akan di return.Perhatikan juga pada line 73, kita menggunakan function strings.ToLower untuk mengecilkan huruf.Function strings.ToLower  berasal dari package strings. Kemudian pada line 60, variable bernama find akan menampung nilai return dari function findStudent. Karena function findStudentme-return sebuah closure, maka untuk memanggil closure yang di return oleh function findStudent, kita perlu memanggil variable find sekaligus memberikan parameternya.


PENJELASAN CODE :
Di dalam fungsi main, ada sebuah slice yang berisi daftar nama mahasiswa, yaitu studentLists.

Kemudian, ada variabel find yang akan menampung hasil dari pemanggilan fungsi findStudent(studentLists).

Fungsi findStudent adalah fungsi yang menerima slice students (daftar nama mahasiswa) sebagai parameter dan mengembalikan fungsi closure.

Closure yang dikembalikan oleh findStudent memiliki satu parameter, yaitu s (nama yang akan dicari). Closure ini mencari nama mahasiswa dalam daftar mahasiswa yang diberikan.

Closure akan mengecek setiap nama di dalam daftar menggunakan perulangan for. Jika nama yang diberikan ditemukan dalam daftar (ignoring case), maka closure akan mengembalikan pesan yang menyatakan nama ditemukan beserta posisinya. Jika tidak ditemukan, pesan bahwa nama tidak ada akan dikembalikan.

Fungsi strings.ToLower digunakan untuk mengubah nama yang diberikan menjadi huruf kecil, sehingga pencarian tidak bersifat case sensitive.

Hasil dari pemanggilan closure dengan memberikan argumen "herdiyan" kemudian dicetak menggunakan fmt.Println.

Jadi, kesimpulannya, kode ini menggunakan closure untuk membuat sebuah fungsi pencarian yang fleksibel dan dapat digunakan untuk mencari nama mahasiswa dalam daftar mahasiswa yang diberikan.
*/
