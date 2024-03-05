package main

import "fmt"

func main() {
	var fruits = make([]string, 3)

	fruits = append(fruits, "apple", "banana", "mango")

	fmt.Printf("%#v", fruits)
	// []string{"", "", "", "apple", "banana", "mango"}
}

/*
Slice - Sesi 2 Slice (append function)Jika kita ingin menambahkan element pada slice dari variablefruits pada halaman sebelumnya, kita bisa melakukannyadengan cara mengakses indexnya seperti pada gambarpertama di sebelah kanan.Jika ingin lebih mudah, maka kita dapat memanfaatkan fungsiappend. Fungsi append akan mengembalikan nilai dari sliceyang ditambahkannya, maka dari itu kita harus menyimpanfungsi append ke dalam suatu variable.Karena kita ingin agar variable fruits yang bertambahelementnya, maka dari itu kita me-reassign variable fruitsdengan fungsi append. Parameter pertama yang diberikanpada fungsi append adalah slice yang ingin ditambahkan, laluparameter setelahnya adalah element-element yang inginditambahkan dan jangan lupa dipisahkan dengan kom

*/
