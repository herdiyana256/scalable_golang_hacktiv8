package main

import (
	"fmt"
	"strings"
)

func main() {
	// 	// Mendefinisikan sebuah slice (array dinamis) dengan nama-nama

	var names = []string{"Herdiyan", "Adam"}

	// Memanggil fungsi greet dengan pesan "Hei," dan slice names sebagai argumen

	var printMsg = greet("Hei,", names)

	// Mencetak pesan yang dikembalikan oleh fungsi greet

	fmt.Println(printMsg)
}

// Definisi fungsi greet dengan dua parameter: msg (string) dan names (slice string)
// Fungsi ini mengembalikan sebuah string yang merupakan pesan salam dengan nama-nama yang disertakan
func greet(msg string, names []string) string {

	// Menggabungkan nama-nama dalam slice menjadi satu string yang dipisahkan oleh koma dan spasi
	var joinStr = strings.Join(names, ", ")

	// Membuat pesan salam dengan menggunakan pesan (msg) yang diberikan dan nama-nama yang sudah digabungkan
	var result string = fmt.Sprintf("%s %s", msg, joinStr)

	// Mengembalikan pesan salam yang telah dibuat
	return result // Hei, Herdiyan, Adam
}
