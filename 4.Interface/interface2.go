package main

import (
	"math"
)

type share interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.height + r.width)
}

/*
InterfacePerhatikan pada gambar disebelah kanan. Terdapat sebuah interface bernama shape yang sudah kita bahas pada halaman sebelumnya. Lalu terdapat 2 buah struct bernama circle dan rectangle. Struct circle mempunyai satu field bernama radius yang memiliki tipe data float64. Kemudian struct rectangle memiliki 2 field bernama width dan height yang keduanya memiliki tipe data yang sama yaitu float64.Karena kita ingin menggunakan tipe data interface shape dan ingin kita gunakan untuk struct circle dan shape, maka kedua struct tersebut perlu mengimplementasikan method-method yang didefinisikan oleh interface shape. ●Pada line 20 - 22, struct circle mengimplementasikan method area untuk menghitung luas lingkaran.●Pada line 24-26, struct rectangle mengimplementasikanmethod area untuk menghitung luas persegi panjang.●Pada line 28-30, struct circle mengimplementasikanmethod perimeter untuk menghitung keliling lingkaran.●Pada line 32-34, struct rectangle mengimplementasikanmethod area untuk menghitung keliling persegi panjang.


*/
