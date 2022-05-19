package helper

import "strings" 

func ValidasiDataPendaftar(namaDepan string, namaBelakang string, email string, jumlahTiket uint, tiketTersedia uint) (bool, bool, bool) {
	// input validation
		cekNamaLengkap := len(namaDepan) >= 2 && len(namaBelakang) >= 2
		cekEmail := strings.Contains(email, "@")
		cekJumlahTiket := jumlahTiket <= tiketTersedia

		return cekEmail, cekNamaLengkap, cekJumlahTiket
}