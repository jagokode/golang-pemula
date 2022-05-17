package main

import "fmt"

func pilihKota() {
	var kota string = "Jakarta"

switch kota {
	case "Bandung":
		// ...
	case "Semarang", "Yogyakarta":
		// ...
	case "Malang", "Surabaya":
		// ...
	case "Denpasar":
		// ...
	default:
		fmt.Println("Kota yang anda pilih tidak tersedia")
}
}
