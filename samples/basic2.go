// arrays in Go
// Fixed size (size = total elements)
// var variable_name = [size]variable type
// Add new elements by their index

package main

import "fmt"

func basic2() {
	var namaDepan string
	var namaBelakang string
	// var pemesanan [50]string
	pemesanan := [50]string{} // cara lain deklarasi array
	pemesanan[0] = namaDepan + " " + namaBelakang

	fmt.Println("Nama Depan: ")
	fmt.Scan(&namaDepan)
	fmt.Println("Nama Belakang: ")
	fmt.Scan(&namaBelakang)

	fmt.Printf("Info lengkap: %v\n", pemesanan)
	fmt.Printf("Tipe data: %T\n", pemesanan)
	fmt.Printf("Tiket atas nama %v\n", pemesanan[0])
	fmt.Printf("Panjang array: %v\n", len(pemesanan))
}
