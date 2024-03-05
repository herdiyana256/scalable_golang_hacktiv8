package main

import "fmt"

func main() {
	// Mendeklarasi tipe data alias bernama second
	// type nama_alias = nama_tipe_data
	type second = uint

	var hour second = 3600
	fmt.Printf("hour type:  %T\n", hour) // output : hour type:  uint
}

//hour type:  uint

/*
Kita juga dapat mendeklarasikan tipe data alias dengan nama yang kita buat sendiri, contohnya seperti padagambar diatas:
●Terdapat sebuah tipe data bernama second yang merupakan tipe data alias dari uint.
●Variable hour memilki tipe data second yang merupakan sebuah tipe data alias.
●Ketika variable hour ingin di ketahui tipe data nya dengan menggunakan flag %T, maka hasil padaterminal akan tetapu menunjukan bahwa tipe data dari variable hour ada uint. Hal ini terjadi karena tipedata asli atau underlying type dari tipe data second adalah tipe data uint.
*/
