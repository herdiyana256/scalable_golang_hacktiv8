/*
Slice (Backing array)Setiap kita membuat suatu slice pada bahasa Go , secara otomatis Go akan membuat suatu array tersembunyi yangdisebut dengan Backing array. Backing array akan bertugas untuk menyimpan element pada slice, bukan slice nya sendiri.Bahasa Go mengimplementasikan slice sebagai sebuah struktur data yang disebut dengan slice header. Slice headerterdiri dari:-Alamat memori/address dari backing array.-Panjang dari slice yang bisa didapatkan dari fungsi len.-Kapasitas dari slice yang bisa didapatkan dari fungsi cap.

*/

package main

import "fmt"

func main() {

	var fruits1 = []string{"Herdiyan", "Adam", "Putra", "Pasti", "Kaya"}

	var fruits2 = fruits1[2:4]

	fruits2[0] = "Sukses"

	fmt.Println("fruits1 => ", fruits1)
	fmt.Println("fruits2 =>", fruits2)

}

/*
fruits1 =>  [Herdiyan Adam Sukses Pasti Kaya]
fruits2 => [Sukses Pasti]


PENJELASAN :
Slice (Backing array)Ketika kita mencoba untuk mendapatkan  beberapa elementdari sebuah slice yang sudah ada dengan cara melakukanslicing, maka Go tidak akan membuat suatu backing arraybaru melainkan slice tersebut akan berbagi backing arrayyang sama dengan slice yang sudah ada.Contohnya seperti pada gambar pertama di sebelah kanan.Variable fruits2 melakukan slicing terhadap fruits1 untukmendapatkan element dari index ke-2 sampai ke-3 dari fruits1atau yang berarti buah durian dan banana.

Setelah itu fruits2 mencoba untuk mengganti buah yangberada pada index ke-0 (buah durian) menjadi buahrambutan.Kemudian ketika kita jalankan pada terminal maka kita dapatmenyadari bahwa element fruits1 pada index ke-2 yang ikutberganti menjadi buah rambutan.Ini terjadi karena variable fruits1 dan fruits2 masih dalam satubacking array yang sama.Hal ini yang menyebabkan penggunaan slice lebih hematmemori jika dibandingkan dengan tipe data array.


*/
