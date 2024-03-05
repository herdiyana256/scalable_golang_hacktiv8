package main

import "fmt"

func main() {
	var right = true
	var wrong = false

	var wrongAndRight = wrong && right
	fmt.Printf("wrong && right \t(%t)\n", wrongAndRight)

	wrongAndRight = wrong || right
	fmt.Printf("wrong || right \t(%t)\n", wrongAndRight)

	var wrongReverse = !wrong
	fmt.Printf("!wrong \t\t(%t)\n", wrongReverse)
}

/*

Stdout
wrong && right 	(false)
wrong || right 	(true)
!wrong 		(true)



OPERATORS (Logical operators)

Seperti yang kita lihat pada gambar disebelah kanan,terdapat sebuah variable bernama right dengan nilaitrue lalu terdapat sebuah variable bernama wrongdengan nilai false. Lalu terdapat variable dengan nama wrongAndRightyang dimana kita melakukan pengecekan kombinasitipe data bool antart variable wrong dengan right.Variable wrongAndRight akan menghasilkan nilai falsekarena jika kita memakai operator &&, maka keduanilainya harus sama dengan true.


Kemudian terdapat sebuah variable dengan namawrongOrRight yang menghasilkan nilai true karena jikakita memakai operator || dan ingin menghasilkan nilaitrue, maka salah satu dari nilainya boleh false.Dan yang terakhir terdapat sebuah variable bernamawrongReverse yang menghasilkan nilai true, karenamemang jika menggunakan operator dengan simbol !,maka operator tersebut akan membalikan nilai daribool nya, misalkan nilainya true maka akan menjadifalse dan begitu juga sebaliknya.
*/
