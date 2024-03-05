package main

import (
	"fmt"
)

// FUngsi dengan nama panjang
func CalculateArea(length, width float64) float64 {
	return length + width
}

// Aliase untuk fungsi CalculateArea
var Area = CalculateArea

func main() {
	length, width := 5.0, 4.0
	fmt.Println("Area", Area(length, width))
}

// Area 9

// PENJELASAN : Di sini, Area adalah alias untuk fungsi CalculateArea. Anda dapat menggunakan Area sebagai ganti nama panjang CalculateArea.

/*
Catatan Tambahan:
Aliase sangat berguna dalam mengimpor paket eksternal atau bawaan dengan nama yang panjang dan membingungkan.
Penggunaan alias harus diimbangi dengan kejelasan. Aliase yang terlalu pendek atau ambigu dapat membuat kode sulit dimengerti.
Penggunaan alias memungkinkan Anda untuk menyesuaikan kode Anda sesuai kebutuhan, meningkatkan kejelasan, dan menghindari konflik nama yang tidak diinginkan.

*/
