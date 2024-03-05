package main

import "fmt"

func main() {
	// Mendeklarasikan map dengan tipe string sebagai kunci dan int sebagai nilai

	var employeAge = map[string]int{
		"herdi": 28,
		"helmi": 25,
		"adam":  27,
	}

	// Menambakan elemet baru ke map
	employeAge["sukses"] = 38

	// mengakses nilai dalam map
	fmt.Println("Umur herdi:", employeAge["herdi"])

	delete(employeAge, "helmi") // menghapus elemen dari

	//Iterasi melalui map
	for name, age := range employeAge {
		fmt.Println("%s berumur %d tahun\n", name, age)

	}

}

/*
Umur herdi: 28
%s berumur %d tahun
 sukses 38
%s berumur %d tahun
 herdi 28
%s berumur %d tahun
 adam 27
PS D:\HERDIYAN A

*/
