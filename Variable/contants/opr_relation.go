package main

import "fmt"

func main() {
	var firstCondition bool = 2 < 3
	var secondCondition bool = "vape" == "Vape"
	var thirdCondition bool = 10 != 2.3
	var fourthCondition bool = 11 <= 11

	fmt.Println("first condition:", firstCondition)
	fmt.Println("second condition:", secondCondition)
	fmt.Println("third condition:", thirdCondition)
	fmt.Println("fourth condition:", fourthCondition)
}

/*
first condition: true
second condition: false
third condition: true
fourth condition: true



Penjelasan :
Jika kita perhatikan pada gambar disebelah kanan kita telah membuat 4 variable yang dimana semuanya akan menghasilkan tipe data bool yang dihasilkan dari pengcekan kesamaan nilai menggunakan operator perbandingan


Jika kita jalankan pada terminal kita, maka kita akan melihat hasilnya seperti pada gambar kedua. vairable bernama firstCondition akan menghasilkan  nilai true karena memang angka 2 lebih kecil dari angka 3




Penjelasan II :
Variale bernama secondCondition akan menghasilkan nilai false karena kata "vape" tidak sama dengan kata "Vape".const

Variable bernama thirdCondition akan menghasilkan nilai true karena angka 10 tidak sama dengan angka 2.3

dan yang terakhir variable bernama fourthCondition akan menghasilkan nilai true karena angka 11 kurang dari atau sama dengan angka 11
*/
