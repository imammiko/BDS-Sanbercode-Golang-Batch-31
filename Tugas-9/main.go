package main

import (
	"BDS-Sanbercode-Golang-Batch-31/Tugas-9/interfaceload"
	"fmt"
)

func main() {
	//soal 1
	var bangunDatar interfaceload.HitungBangunDatar

	bangunDatar = interfaceload.SegitigaSamaSisi{6, 12}
	fmt.Println("==== segitiga sama sisi")
	fmt.Println(bangunDatar.Luas())
	fmt.Println(bangunDatar.Keliling())

	bangunDatar = interfaceload.PersegiPanjang{6, 12}
	fmt.Println("===== persegi panjang")
	fmt.Println(bangunDatar.Keliling())
	fmt.Println(bangunDatar.Luas())

	var bangunRuang interfaceload.HitungBangunRuang
	fmt.Println("=== bangun tabung")
	bangunRuang = interfaceload.Tabung{float64(12), float64(6)}
	fmt.Println(bangunRuang.LuasPermukaan())
	fmt.Println(bangunRuang.Volume())
	fmt.Println("==== bangun balok")
	bangunRuang = interfaceload.Balok{20, 20, 20}
	fmt.Println(bangunRuang.LuasPermukaan())
	fmt.Println(bangunRuang.Volume())
	// soal 2
	fmt.Println("=========")
	var phoneProperty interfaceload.PhoneBrand
	phoneProperty = interfaceload.Phone{"Samsung Galaxy Note 20", "Samsung Galaxy Note 20", 2020, []string{"Mystic Bronze", "Mystic White", "Mystic Black"}}
	fmt.Println(phoneProperty.Print())
	//soal 3
	fmt.Println(interfaceload.LuasPersegi(4, true))

	fmt.Println(interfaceload.LuasPersegi(8, false))

	fmt.Println(interfaceload.LuasPersegi(0, true))

	fmt.Println(interfaceload.LuasPersegi(0, false))
	//soal 4
	var prefix interface{} = "hasil penjumlahan dari "

	var kumpulanAngkaPertama interface{} = []int{6, 8}

	var kumpulanAngkaKedua interface{} = []int{12, 14}
	var sum int
	var loadString string

	for i := 0; i < len(kumpulanAngkaKedua.([]int)); i++ {
		sum += kumpulanAngkaKedua.([]int)[i]
		sum += kumpulanAngkaPertama.([]int)[i]

	}
	arrayload := append(kumpulanAngkaPertama.([]int), kumpulanAngkaKedua.([]int)...)
	loadString = interfaceload.ArrayToString(arrayload, "+")
	fmt.Printf("%s %s %d", prefix, loadString, sum)
	// tulis jawaban anda disini

}
