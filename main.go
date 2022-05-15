package main

import "fmt"

func main() {
	namaKonferensi := "Konferensi Golang" 
	const tiketKonferensi = 50
	var tiketTersedia uint = 50
	
	fmt.Printf("Selamat Datang di Aplikasi Pendaftaran %v\n", namaKonferensi)
	fmt.Printf("Total tiket tersisa sejumlah %v dari total tiket sebanyak %v\n", tiketTersedia, tiketKonferensi)
	fmt.Println("Dapatkan tiket disini")

	var namaDepan string
	var namaBelakang string
	var email string 
	var jumlahTiket uint
	// get user input
	fmt.Println("Ketik nama depan anda: ")
	fmt.Scan(&namaDepan) // using pointer
	fmt.Println("Ketik nama belakang anda: ")
	fmt.Scan(&namaBelakang)
	fmt.Println("Ketik email anda: ")
	fmt.Scan(&email)
	fmt.Println("Ketik jumlah tiket yang anda pesan: ")
	fmt.Scan(&jumlahTiket)

	fmt.Printf("Terima kasih %v %v sudah memesan %v tiket. Anda akan menerima konfirmasi melalui email %v", namaDepan, namaBelakang, jumlahTiket, email)
}

