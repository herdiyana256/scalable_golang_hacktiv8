// Nested conditions

// Kita juga dapat melakukan kondisional bersarang seperti pada gambar disebelah kanan.

// Kondisional bersarang seperti pada gambar disebelah kanan.

// Kondisional bersarang merupakan sebuah proses kondisional yang didalam nya terdapat proses kondisional kembali. kita dapat menggunakan if, else if, else switch ataupun menggabungkan seluruhnya.

package main

import "fmt"

func main() {
	score := 10

	if score > 7 {
		switch score {
		case 10:
			fmt.Println("perfect!") // perfect!
		default:
			fmt.Println("nice!")
		}
	} else {
		if score == 5 {
			fmt.Println("not bad")
		} else if score == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it")
			if score == 0 {
				fmt.Println("try harder!")
			}
		}
	}
}
