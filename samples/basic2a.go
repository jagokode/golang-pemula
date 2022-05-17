// Slices in Go
// Slice is an abstraction of an Array
// Slices are more flexible and powerful: variable-length or get an sub-array of its own
// Slices are also index-based and have a size, but is resized when needed
// append : adds the element(s) at the end of the slice

package main

import "fmt"

func Slices() {
	// Slices
	var nama string 
	var pekerjaan string 
	// var pemesanan []string 
	pemesanan := []string{} // cara lain deklarasi slice
	// pemesanan[0] = nama + " " + pekerjaan
	pemesanan = append(pemesanan, nama + " " + pekerjaan)

	fmt.Println(pemesanan[0])
}