/*
ngan menambahkan tanda titik tiga tersebut pada sebuahparameter, maka sebuah dianggap sebagai sebauh function variadic,dan parameter yang diterminya akan di konversikan menjadi sebuahslice. Maka dari itu kita dapat melakukan looping terhadapparameter yang diterima oleh function print pada line 118 - 124.Function print  digunakan untuk menerima banyak list darinama-nama murid dan setelah itu dikonversikan menjadi sebuah sliceyang mengandung tipe data map. Pada line 119, kita menggunakanfungsi fmt.Sprintf untuk membuat penamaan key yang unikberdasarkan index looping nya dan index tersebut kita tambahkandengan angka 1 agar tidak dimulai dari angka 0.

*/

package main  

import "fmt"


func main () {

	numberLists := []int{1, 2, 3, 4, 5, 6, 7, 8}

	result := sum(numberLists...)

	fmt.Println("Result: ", result)
}

	func sum(numbers ...int) int {
			total := 0 

			for _, := range numbers  {
				total += v 
			}

			return total 
	}

	/*
PENJELASAN CODE : 
Pada baris-baris awal yang diapit oleh /* dan /*,  terdapat komentar yang menjelaskan tentang penggunaan variadic function dalam Go.

Baris 11-17: Fungsi main() merupakan fungsi utama dari program. Di dalamnya, kita mendefinisikan sebuah slice numberLists yang berisi beberapa angka bulat. Kemudian, kita memanggil fungsi sum() dengan menggunakan slice numberLists sebagai argumen. Hasil yang dikembalikan oleh sum() akan disimpan dalam variabel result dan dicetak menggunakan fmt.Println().

Baris 19-26: Fungsi sum(numbers ...int) int adalah variadic function yang menerima argumen bertipe int tanpa batas. Fungsi ini mengembalikan nilai bertipe int yang merupakan hasil penjumlahan dari semua argumen yang diberikan.

Variadic parameter dinyatakan dengan ...int. Ini berarti kita dapat memasukkan nol atau lebih nilai int sebagai argumen saat memanggil fungsi sum().

Di dalam fungsi sum(), kita melakukan iterasi melalui semua argumen yang diberikan menggunakan loop for range. Setiap nilai int akan ditambahkan ke variabel total.

Setelah selesai iterasi, nilai total dari semua argumen dijumlahkan akan dikembalikan sebagai hasil dari fungsi sum().

	*/