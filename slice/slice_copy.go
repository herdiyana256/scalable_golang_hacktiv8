package main

import "fmt"

func main() {

	var fruits1 = []string{"apple", "banana", "mango"}

	var fruits2 = []string{"durian", "pipeapple", "starfruit"}

	nn := copy(fruits1, fruits2)
	// The length of the slice is not changed. It only copies the elements from one slice to another.

	fmt.Println("Fruits1 =>", fruits1)
	fmt.Println("Fruits2 =>", fruits2)
	fmt.Println("Copied elements =>", nn)
}

/*
Fruits1 => [durian pipeapple starfruit]
Fruits2 => [durian pipeapple starfruit]
Copied elements => 3



PENJELASAN :
Kita juga bisa menggunakan fungsi copy untuk meng-copyseluruh element pada sebuah slice ke dalam slice lainnya.Perlu diingat disini bahwa ketika kita mencoba untukmeng-copy sebuah slice kedalam slice lainnya, maka seluruhelement pada slice lainnya tersebut akan ter-replace olehelement-element yang di copy kannya.Contohnya seperti pada gambar pertama di sebelah kanan.


*/
