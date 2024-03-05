package main

import "fmt"

func main() {

	cars := []string{"CR-V", "HR-V", "X-PANDER", "PAJERO", "FORTUNER"}

	newCars := []string{}

	newCars = append(newCars, cars[0:2]...) // CR-V and HR-V

	cars[0] = "Nissan"
	fmt.Println("cars:", cars)
	fmt.Println("newCars:", newCars)

}

/*
cars: [Nissan HR-V X-PANDER PAJERO FORTUNER]
newCars: [CR-V HR-V]



PENJELASAN :
Slice (Creating a new backing array)Ketika kita ingin mendapatkan element-element dari slice yang sudah ada, namun kita juga ingin membuat backing array yang baru, maka kita dapat menggunakan fungsi append untuk melakukannya. Contohnya seperti pada gambar pertama di sebelah kanan. Variable newCars ingin mendapatkan element-element dimulai dari index ke-0 hingga index ke-1 dari variable cars. Namun newCars mendapatkan element-elementnya dengan menggunakan fungsi append walaupun masih menggunakan slicing di dalam fungsi append nya. Ketika index ke-0 dari cars dirubah, maka newCars tidak ikut terubah dikarenakan mereka tidak memiliki backing array yang sama. Jika dijalankan di terminal kita, maka hasilnya akan seperti pada gambar kedua.
*/
