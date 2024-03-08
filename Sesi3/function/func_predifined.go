/*
Function (Predefined return value)Kita juga dapat menggunakan sebuah variable untuk mendefinisikannilai return dari sebuah function. Contoh pada function calculatepada gambar pertama di sebelah kanan. Function calculate tersebutmasih memiliki fungsi yang sama seperti pada halaman sebelumnya,namun perbedaannya adalah kita menggunakan teknik predefinedreturn value yang dimana kita dapat memberikan sebuah variablesebagai hasil nilai return dari sebuah function.Jika kita perhatikan pada gambar pertama, variable area dancircumference yang berada di dalam function calculate bukanlahsebuah variable baru, melainkan kedua variable tersebut adalahvariable yang digunakan sebagai nilai return.Jika kita menggunakan teknik predefined return value, maka kitaperlu me-reassign variable yang digunakan sebagai nilai returndengan sebuah nilai yang nantinya akan menghasilkan sebuah nilaireturn. Tapi perlu diingat bahwa kita masih tetap membutuhkankeyword return yang diletakkan di akhir baris function.Jika kita jalankan pada terminal kita, maka hasilnya seperti padagambar kedua.
Function - Sesi 3

*/

package main

import (
	"fmt"
	"math"
)

func main() {
	var diameter float64 = 15
	var area, circumference = calculate(diameter)

	fmt.Println("Area:", area)
	fmt.Println("Circumference:", circumference)

}
func calculate(d float64) (area float64, circumference float64) {

	// Menghitung luas
	area = math.Pi * math.Pow(d/2, 2)
	// Menghitung keliling

	circumference = math.Pi * d

	return
}

/*
	PENJELASAN :
	Area: 176.71458676442586
Circumference: 47.12388980384689


PENJELASAN CODE :
Pada baris 1-3, kita mendefinisikan package main dan mengimpor dua package yaitu fmt (untuk input/output) dan math (untuk fungsi matematika).

Fungsi utama dari program (main()) dimulai pada baris 5 dan berakhir pada baris 12. Di dalamnya:

Pada baris 6, sebuah variabel diameter dengan tipe data float64 dideklarasikan dan diinisialisasi dengan nilai 15.
Pada baris 7, kita memanggil fungsi calculate(diameter) untuk menghitung luas dan keliling lingkaran, dan menyimpan hasilnya dalam variabel area dan circumference.
Pada baris 9-10, hasil perhitungan luas dan keliling lingkaran dicetak menggunakan fungsi fmt.Println().
*/
