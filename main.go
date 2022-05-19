package main

import (
	"booking-app/helper"
	"fmt"
	"strconv"
)

// Package Level Variables (Didefinisikan di atas semua function)
const tiketKonferensi = 50
var tiketTersedia uint = 50
var namaKonferensi = "Konferensi Golang" 
var pemesanan = make([]map[string]string, 0)
// var pemesanan = []string{}

func main() {
	salamPembuka()		

	for {
		namaDepan, namaBelakang, email, jumlahTiket := dataPendaftar()

		cekNamaLengkap, cekEmail, cekJumlahTiket := helper.ValidasiDataPendaftar(namaDepan, namaBelakang, email, jumlahTiket, tiketTersedia)

		if cekJumlahTiket && cekNamaLengkap && cekEmail {
			pesanTiket(jumlahTiket, namaDepan, namaBelakang, email)
		
			daftarNamaDepan := daftarPeserta()

			fmt.Printf("Daftar peserta : %v\n", daftarNamaDepan)

			if tiketTersedia == 0 {
				// end program
				fmt.Println("Tiket Konferensi sudah habis. Coba lagi tahun depan!")
				break
			}
		} else {
					if !cekNamaLengkap {
						fmt.Println("Nama Depan atau Nama Belakang minimal 2 karakter")
					}
					if !cekEmail {
						fmt.Println("Alamat email yang anda masukkan salah")
					}
					if !cekJumlahTiket {
						fmt.Println("Jumlah tiket yang anda pesan melebihi kuota yang tersedia")
					}
		}
	}
}

func salamPembuka() {
	fmt.Printf("Selamat Datang di Aplikasi Pendaftaran %v\n", namaKonferensi)
	fmt.Printf("Total tiket tersisa sejumlah %v dari total tiket sebanyak %v\n", tiketTersedia, tiketKonferensi)
	fmt.Println("Dapatkan tiket disini")		
}

func daftarPeserta() []string {
	daftarNamaDepan := []string{}
	for _, peserta := range pemesanan {
			// var daftarNama = strings.Fields(peserta)
			daftarNamaDepan = append(daftarNamaDepan, peserta["namaDepan"])
	}
	
	return daftarNamaDepan
}

func dataPendaftar() (string, string, string, uint) {
	var namaDepan string
		var namaBelakang string
		var email string 
		var jumlahTiket uint	 

	// get user input
		fmt.Println("Ketik nama depan anda (minimal 2 karakter): ")
		fmt.Scan(&namaDepan) 

		fmt.Println("Ketik nama belakang anda (minimal 2 karakter): ")
		fmt.Scan(&namaBelakang)

		fmt.Println("Ketik email anda (example@mail.com): ")
		fmt.Scan(&email)

		fmt.Println("Ketik jumlah tiket yang anda pesan: ")
		fmt.Scan(&jumlahTiket)

		return namaDepan, namaBelakang, email, jumlahTiket
}

func pesanTiket(jumlahTiket uint, namaDepan string, namaBelakang string, email string) {
	tiketTersedia = tiketTersedia - jumlahTiket

	// buat map untuk pemesan
	var dataPemesan = make(map[string]string) 
	dataPemesan["namaDepan"] = namaDepan
	dataPemesan["namaBelakang"] = namaBelakang
	dataPemesan["email"] = email
	dataPemesan["jumlahTiket"] = strconv.FormatUint(uint64(jumlahTiket), 10)

	pemesanan = append(pemesanan, dataPemesan)
	fmt.Printf("Daftar pemesanan: %v\n", pemesanan)
	// pemesanan = append(pemesanan, namaDepan + " " + namaBelakang)
	
	fmt.Printf("Terima kasih %v %v sudah memesan %v tiket. Anda akan menerima konfirmasi melalui email %v\n", namaDepan, namaBelakang, jumlahTiket, email)
	fmt.Printf("%v tiket tersedia untuk %v\n", tiketTersedia, namaKonferensi)
}