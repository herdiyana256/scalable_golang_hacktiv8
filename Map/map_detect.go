package main

import "fmt"

func main() {
	var person = map[string]string{
		"name":     "Universitas saintek muhammadiyah",
		"location": "Jakarta",
		"address":  "Jalan Kelapa 2 wetan, ciracas",
	}

	value, exist := person["location"]

	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("Value does'nt exist")
	}

	delete(person, "location")

	value, exist = person["location"]

	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("Value has been deleted")
	}

}

/*
Jakarta
Value has been deleted


jika  penampung variable lebih dari 1
Value does'nt exist
Value has been deleted


PENJELASAN :
Map (Detecting a value)Kita juga dapat mendeteksi keberadaan suatu value dengan caramengakses key dari map nya lalu memberikan 2 variable sebagaipenampungnya.Seperti yang terdapat pada line 92 dan line 102 pada gambarpertama disebelah kanan. Perlu diingat disini bahwa memberikan 2variable sebagai penampung merupakan hal opsional. Hal ini kitalakukan hanya untuk mendeteksi keberadaan suatu value.Variable value yang kita berikan akan mengembalikan value asli darimap nya jika memang key nya ada, jika tidak ada maka kita akanmendapat zero value dari value nya. Kemudian variable exist yangkita berikan akan mengembalikan nilai dengan tipe data bool.Variable exist akan mengembalikan true jika memang key nya ada,jika tidak maka akan mengembalikan false.Jika kita jalankan pada terminal kita, maka hasilnya akan seperti padagambar kedua.

*/
