package main

import (
	"fmt"
	"strconv"
)

func main() {
	// soal 1
	var panjangPersegiPanjang string = "8"
	var lebarPersegiPanjang string = "5"

	var alasSegitiga string = "6"
	var tinggiSegitiga string = "7"

	var luasPersegiPanjang int
	var kelilingPersegiPanjang int
	var luasSegitiga int
	panjang, _ := strconv.Atoi(panjangPersegiPanjang)
	lebar, _ := strconv.Atoi(lebarPersegiPanjang)
	alas, _ := strconv.Atoi(alasSegitiga)
	tinggi, _ := strconv.Atoi(tinggiSegitiga)
	luasPersegiPanjang = panjang * lebar
	kelilingPersegiPanjang = 2 * (panjang + lebar)
	luasSegitiga = (alas * tinggi) / 2
	fmt.Println(luasPersegiPanjang)
	fmt.Println(kelilingPersegiPanjang)
	fmt.Println(luasSegitiga)
	// soal 2
	var nilaiJohn = 80
	var nilaiDoe = 50
	var indeks string
	switch {
	case nilaiJohn >= 80:
		indeks = "A"
	case nilaiJohn >= 70 && nilaiJohn < 80:
		indeks = "B"
	case nilaiJohn >= 60 && nilaiJohn < 70:
		indeks = "C"
	case nilaiJohn >= 50 && nilaiJohn < 60:
		indeks = "D"
	case nilaiJohn < 50:
		indeks = "E"
	}
	fmt.Printf("nilai si Doe %s ", indeks)
	fmt.Println("")
	switch {
	case nilaiDoe >= 80:
		indeks = "A"
	case nilaiDoe >= 70 && nilaiDoe < 80:
		indeks = "B"
	case nilaiDoe >= 60 && nilaiDoe < 70:
		indeks = "C"
	case nilaiDoe >= 50 && nilaiDoe < 60:
		indeks = "D"
	case nilaiDoe < 50:
		indeks = "E"
	}
	fmt.Printf("nilai si Doe %s ", indeks)
	fmt.Println("")
	// 3
	var tanggal = 26
	var bulan = 12
	var tahun = 1995
	var bulanWord string
	switch bulan {
	case 1:
		bulanWord = "Januari"
	case 2:
		bulanWord = "Ferbuari"
	case 3:
		bulanWord = "Maret"
	case 4:
		bulanWord = "April"
	case 5:
		bulanWord = "Mei"
	case 6:
		bulanWord = "Juni"
	case 7:
		bulanWord = "Juli"
	case 8:
		bulanWord = "Agustus"
	case 9:
		bulanWord = "September"
	case 10:
		bulanWord = "Oktober"
	case 11:
		bulanWord = "November"
	case 12:
		bulanWord = "Desember"

	}
	fmt.Printf("%d %s %d", tanggal, bulanWord, tahun)
	fmt.Println("")
	//4
	var generasi string
	switch {
	case tahun >= 1944 && tahun <= 1964:
		generasi = "Baby Boomer"
	case tahun >= 1965 && tahun <= 1979:
		generasi = "X"
	case tahun >= 1980 && tahun <= 1994:
		generasi = "Y"
	case tahun >= 1994 && tahun <= 2015:
		generasi = "Z"
	}
	fmt.Printf("kelahiran saya generasi %s", generasi)
}
