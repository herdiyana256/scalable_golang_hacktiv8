/*
InterfaceInterface adalah sebuah tipe data pada bahasa Go yang merupakan kumpulan dari definisi satu atau lebih method. Kita dapat menggunakan tipe data interface jika kita telah mengimplementasikan method-method yang telah didefinisikan oleh interface tersebut melalui tipe data lainnya. Cara membuat sebuah interface cukup mudah, contohnya seperti pada gambar dibawah ini.Gambar diatas merupakan sebuah interface bernama shape  yang mempunyai 2 method bernama area, dan perimeter. Method area merupakan sebuah method yang me-return nilai dengan tipe data float64 dan method perimeter merupakan sebuah method yang me-return nilai dengan tipe data float64. Kita dapat menggunakan interface shape jika kita telah mengimplementasikan method-method yang telah didefinisikan oleh interface shape melalui tipe data lainnya. Untuk sesi kali ini, kita akan mengimplementasikan method-method dari interface shape melalui tipe data struct.

*/

type shape interface {
	area() float64
	perimeter() float64
}

/*
Cara membuat sebuah interface cukup mudah, contohnya seperti pada gambar dibawah ini.Gambar diatas merupakan sebuah interface bernama shape  yang mempunyai 2 method bernama area, dan perimeter. Method area merupakan sebuah method yang me-return nilai dengan tipe data float64 dan method perimeter merupakan sebuah method yang me-return nilai dengan tipe data float64. Kita dapat menggunakan interface shape jika kita telah mengimplementasikan method-method yang telah didefinisikan oleh interface shape melalui tipe data lainnya. Untuk sesi kali ini, kita akan mengimplementasikan method-method dari interface shape melalui tipe data struct.

*/