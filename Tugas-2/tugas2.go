package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// 	soal 1
	// tampilkan kalimat "Bootcamp Digital Skill Sanbercode Golang" yang tersusun dari gabungan variabel dalam setiap kata (5 variabel)
	var word, word1, word2, word3, word4 string
	word, word1, word2, word3, word4 = "Bootcamp ", "Digital ", "Skill ", "Sanbercode ", "Golang"
	fmt.Println(word + word1 + word2 + word3 + word4)
	// soal 2

	// terdapat variabel seperti di bawah ini:

	// halo := "Halo Dunia"

	// ganti kata "Dunia" menjadi "Golang" menggunakan packages strings
	halo := "Halo Dunia"
	halo = strings.Replace(halo, "Dunia", "Golang", 1)
	fmt.Println(halo)
	// 	soal 3

	// buatlah variabel-variabel seperti di bawah ini

	// var kataPertama = "saya";
	// var kataKedua = "senang";
	// var kataKetiga = "belajar";
	// var kataKeempat = "golang";
	// gabungkan variabel-variabel tersebut agar menghasilkan output

	// saya Senang belajaR GOLANG
	var kataPertama = "saya"
	var kataKedua = "senang"
	var kataKetiga = "belajar"
	var kataKeempat = "golang"
	fmt.Printf("%s %s %s %s ", kataPertama, kataKedua, kataKetiga, kataKeempat)

	// 	soal 4

	// buatlah variabel-variabel seperti di bawah ini

	// var angkaPertama= "8";
	// var angkaKedua= "5";
	// var angkaKetiga= "6";
	// var angkaKeempat = "7";
	// ubahlah variabel diatas menjadi integer dan jumlahkan semuanya

	var angkaPertama = "8"
	var angkaKedua = "5"
	var angkaKetiga = "6"
	var angkaKeempat = "7"

	angkaPertamaint, _ := strconv.Atoi(angkaPertama)
	angkaKeduaint, _ := strconv.Atoi(angkaKedua)
	angkaKetigaint, _ := strconv.Atoi(angkaKetiga)
	angkaKeempatint, _ := strconv.Atoi(angkaKeempat)

	fmt.Println(angkaPertamaint + angkaKeduaint + angkaKetigaint + angkaKeempatint)

	// soal 5

	// terdapat variabel seperti di bawah ini:

	// kalimat := "halo halo bandung"
	// angka := 2021

	// olah variabel diatas yang hasil output akhrinya adalah

	// "Hi Hi bandung" - 2021
	kalimat := "halo halo bandung"
	angka := 2021
	angkaString := strconv.Itoa(angka)
	kalimatHasil := strings.Replace(kalimat, "halo", "Hi", 2)
	fmt.Println("\"" + kalimatHasil + "\"" + " " + "-" + " " + angkaString)

}
