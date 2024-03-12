/*
Di Go, sebuah interface adalah kumpulan definisi metode. Sebuah objek dapat mengimplementasikan satu atau lebih interface. Interface memungkinkan kita untuk membuat kode yang lebih fleksibel dan dapat digunakan dengan berbagai tipe data tanpa peduli pada tipe konkrit yang sebenarnya. Mari kita lihat contoh sederhana dan kemudian kita bahas:

*/

package main

// Definisi sebuah interface dengan nama Shpae
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Definisi struct rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

// Implementasi metode area untuk Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Implementasikan metode Perimeter untuk Rectangle
func (r Rectangle) Perimeter() float64 {
	return 2*r.Width + 2*r.Height
}

// Definisikan struct Circle
ype Circle struct {
	Radius float64
}

// Implementasikan metode Area untuk Circle
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// Implementasikan metode Perimeter untuk Circle
func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

func printInformation(s Shape) {
	fmt.Printf("Area: %f, Perimeter: %f\n", s.Area(), s.Perimeter())
}

func main() {
	rect := Rectangle{Width: 5, Height: 3}
	circ := Circle{Radius: 4}

	fmt.Println("Rectangle information:")
	printInformation(rect)

	fmt.Println("Circle information:")
	printInformation(circ)
}

/*
Pada contoh di atas, kita mendefinisikan sebuah interface Shape yang memiliki dua metode: Area() dan Perimeter(). Kemudian kita mendefinisikan dua struktur: Rectangle dan Circle. Kedua struktur tersebut mengimplementasikan interface Shape dengan mengimplementasikan kedua metode yang dimiliki oleh interface tersebut.

Kemudian, kita memiliki fungsi printInformation() yang menerima argumen bertipe Shape. Kita bisa memanggil fungsi ini dengan objek bertipe Rectangle atau Circle karena keduanya mengimplementasikan interface Shape.

Ketika program dijalankan, kita mencetak informasi area dan perimeter dari objek-objek tersebut, tanpa perlu peduli apakah itu Rectangle atau Circle. Ini menunjukkan fleksibilitas yang diberikan oleh penggunaan interface di Go.

Dalam bahasa Indonesia, interface bisa dianggap sebagai semacam kontrak yang menentukan perilaku yang harus dipenuhi oleh tipe data tertentu. Dengan menggunakan interface, kita dapat menulis kode yang lebih fleksibel dan reusable, karena kita fokus pada perilaku yang diinginkan daripada pada tipe data konkrit.
*/