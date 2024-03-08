// Function ( returning multiple values )
/*

Function pada bahasa Go dapat me-return lebih dari satu nilai.Contohnya seperti function bernama calculate pada gambar pertamadi sebelah kanan.  Function calculate tersebut berfungsi untukmenghitung luas dan keliling lingkaran. Lalu function calculate akanme-return 2 nilai yaitu hasil dari kalkulasi luas dan kelilingberdasarkan diameter yang diberikan pada parameter functioncalculate.Untuk menghitung luas nya, kita memakai bantuan dari packagemath dengan memanfaatkan konstanta Pi. Konstanta math.Pi adalahkonstanta bawaan package math yang merepresentasikan Pi atau22/7.Kemudian untuk menghitung kelilingnya, kita memakai fungsimath.Pow yang digunakan untuk memangkat nilai. math.Pow(2, 3)berarti 2 pangkat 3, hasilnya 8. Fungsi math.Pow juga berasal daripackage math.


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
func calculate(d float64) (float64, float64) {
	// Menghitung luas

	var area float64 = math.Pi * math.Pow(d/2, 2)
	// Menghitung keliling

	var circumference = math.Pi * d

	return area, circumference
}
