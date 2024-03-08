// Pointer (Changing value throgh pointer)

package main

import (
	"fmt"
	"strings"
)

func main() {
	var firstPerson string = "Herdiyan"     // Mendeklarasikan variabel firstPerson dengan nilai "Herdiyan"
	var secondPerson *string = &firstPerson // Mendeklarasikan variabel secondPerson sebagai pointer ke string dan mengisinya dengan alamat memori dari firstPerson

	fmt.Println("firstPerson (value) :", firstPerson)            // Mencetak nilai dari variabel firstPerson
	fmt.Println("firstPerson (memory address) :", &firstPerson)  // Mencetak alamat memori dari variabel firstPerson
	fmt.Print("secondNumber (value) :", *secondPerson)           // Mencetak nilai yang ditunjuk oleh pointer secondPerson (yang mengacu pada firstPerson)
	fmt.Println("secondNumber (memory address) :", secondPerson) // Mencetak alamat memori yang disimpan dalam variabel secondPerson

	fmt.Println("\n", strings.Repeat("#", 28), "\n") // Mencetak pemisah garis

	*secondPerson = "Adam" // Mengubah nilai yang ditunjuk oleh pointer secondPerson menjadi "Adam"

	fmt.Println("firstPerson (value) :", firstPerson)            // Mencetak nilai dari variabel firstPerson setelah perubahan nilai melalui pointer
	fmt.Println("firstPerson (memory address) :", &firstPerson)  // Mencetak alamat memori dari variabel firstPerson (yang tetap sama)
	fmt.Println("secondNumber (value) :", *secondPerson)         // Mencetak nilai yang ditunjuk oleh pointer secondPerson setelah perubahan
	fmt.Println("secondNumber (memory address) :", secondPerson) // Mencetak alamat memori yang disimpan dalam variabel secondPerson (yang tetap sama)
}

/*
Pada awalnya, variabel firstPerson memiliki nilai "Herdiyan" dan variabel secondPerson menyimpan alamat memori dari firstPerson. Setelah itu, kita mengubah nilai yang ditunjuk oleh secondPerson menjadi "Adam" menggunakan operator dereferensi *. Karena secondPerson mengacu pada firstPerson, perubahan nilai pada secondPerson juga mempengaruhi nilai firstPerson. Oleh karena itu, setelah perubahan, baik firstPerson maupun nilai yang ditunjuk oleh secondPerson adalah "Adam". Alamat memori yang disimpan oleh secondPerson tetap tidak berubah karena kita tidak mengubah alamat memori yang ditunjuk oleh pointer, hanya nilai yang ditunjuknya yang berubah.

*/
