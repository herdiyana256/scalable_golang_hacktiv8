package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
}

/*

0 1 2 3 4
1 2 3 4
2 3 4
3 4
4


PENJELASAN :
Loopings (Nested Looping)
Nested looping atau looping bersarang adalah suatu proses looping yang memiliki suatu proses looping didalam nya.
Contohnya seperti pada gambar pertama dibawah . Jika kita perhatikan, kita menggunkan fungsi fmt.Println dan tidak memberikan argumen apapun kedalam nya.
ini bisa dilakukan ketika kita hanya ingin membuat baris baru. Hasilnya sama saja ketika kita memakai fungsi fmt.Print dengan argumen "\n" atau sama dengan (fmt.Print"\n")).
Jika kita jalankan di terminal, maka hasilnya akan seperti pada gambar kedua dibawah.

*/
