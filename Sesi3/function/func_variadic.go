/*
Function (Variadic function #1)

Bahasa Go telah menyediakan sebuah function yang dimana functiontersebut dapat menerima argumen yang tak terbatas jumlahnya.Function tersebut bernama function variadic. Contohnya sepertipada gambar disebelah kanan.Function print merupakan sebuah function variadic. Caramengetahuinya adalah dengan memperhatikan penulisan padaparameter yang diterimanya. Terdapat tanda titik sebanyak tiga kalisebelum keterangan tipe data parameter yang diterimanya ...string.


denngan menambahkan tanda titik tiga tersebut pada sebuahparameter, maka sebuah dianggap sebagai sebauh function variadic,dan parameter yang diterminya akan di konversikan menjadi sebuahslice. Maka dari itu kita dapat melakukan looping terhadapparameter yang diterima oleh function print pada line 118 - 124.Function print  digunakan untuk menerima banyak list darinama-nama murid dan setelah itu dikonversikan menjadi sebuah sliceyang mengandung tipe data map. Pada line 119, kita menggunakanfungsi fmt.Sprintf untuk membuat penamaan key yang unikberdasarkan index looping nya dan index tersebut kita tambahkandengan angka 1 agar tidak dimulai dari angka 0.
*/

package main

import "fmt"

func main() {
	studentLists := print("Herdiyan", "Adam", "Putra", "Pasti", "Cerdas")

	fmt.Printf("%v", studentLists)

}

func print(names ...string) []map[string]string {
	var result []map[string]string

	for i, v := range names {
		key := fmt.Sprintf("student%d", i+1)
		temp := map[string]string{
			key: v,
		}

		result = append(result, temp)
	}

	return result
}

/*
Penjelasan code :

Pada baris-baris awal yang dimulai dengan komentar yang diapit oleh /* dan //* , menjelaskan tentang konsep variadic function dalam Go.

Baris 15-17: Fungsi main() merupakan fungsi utama program. Di dalamnya, kita memanggil fungsi print dengan beberapa nama siswa sebagai argumen. Hasil yang dikembalikan oleh print akan disimpan dalam variabel studentLists.

Baris 19: Menggunakan Printf untuk mencetak nilai studentLists.

Baris 21-30: Fungsi print(names ...string) []map[string]string adalah variadic function yang menerima argumen bertipe string tanpa batas. Fungsi ini mengembalikan slice ([]) dari peta (map[string]string).

Variadic parameter dinyatakan dengan ...string. Ini berarti bahwa kita dapat memasukkan nol atau lebih string sebagai argumen saat memanggil fungsi print().

Di dalam fungsi print(), kita melakukan iterasi melalui semua argumen yang diberikan menggunakan loop for range. Setiap nama siswa akan diberi kunci yang unik, seperti "student1", "student2", dst. Kemudian, sebuah peta sementara dibuat dengan kunci tersebut dan nama siswa sebagai nilainya.

Peta sementara tersebut kemudian ditambahkan ke slice result menggunakan fungsi append().

Setelah selesai iterasi, slice result yang berisi semua peta dengan nama siswa dan kunci uniknya dikembalikan sebagai hasil dari fungsi print().
*/

// output : [map[student1:Herdiyan] map[student2:Adam] map[student3:Putra] map[student4:Pasti] map[student5:Cerdas]]
