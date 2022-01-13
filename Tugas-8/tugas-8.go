package main

import (
	"fmt"
	"math"
	"strings"
)

type segitigaSamaSisi struct {
	alas, tinggi int
}

type persegiPanjang struct {
	panjang, lebar int
}

type tabung struct {
	jariJari, tinggi float64
}

type balok struct {
	panjang, lebar, tinggi int
}

type hitungBangunDatar interface {
	luas() int
	keliling() int
}

type hitungBangunRuang interface {
	volume() float64
	luasPermukaan() float64
}

func (s segitigaSamaSisi) luas() int {
	return s.alas * s.tinggi / 2
}

type phone struct {
	name, brand string
	year        int
	colors      []string
}

type phoneBrand interface {
	print() string
}

type persegi struct {
	luas   int
	status bool
}

func (p phone) print() string {
	var loadStringPhone string
	loadStringPhone += fmt.Sprintf("name: %s\n", p.name)
	loadStringPhone += fmt.Sprintf("brand: %s\n", p.brand)
	loadStringPhone += fmt.Sprintf("year: %d\n", p.year)
	loadStringPhone += fmt.Sprintf("colors: %s\n", strings.Join(p.colors, ","))
	return loadStringPhone
}

func (s segitigaSamaSisi) keliling() int {
	var pytagoras float64
	var alas float64
	alas = float64(s.alas) / 2
	pytagoras = math.Sqrt(math.Pow(alas, 2) + math.Pow(float64(s.tinggi), 2))

	return (int(pytagoras) * 2) + s.alas
}

func (p persegiPanjang) luas() int {
	return p.panjang * p.lebar
}

func (p persegiPanjang) keliling() int {
	return 2*p.panjang + 2*p.lebar
}
func (t tabung) volume() float64 {
	return math.Pi * math.Pow(t.jariJari, 2) * t.tinggi
}
func (t tabung) luasPermukaan() float64 {
	return 2 * math.Pi * t.jariJari * (t.jariJari + t.tinggi)
}
func (b balok) volume() float64 {
	return float64(b.panjang * b.lebar * b.tinggi)
}
func (b balok) luasPermukaan() float64 {
	return float64(2 * (b.panjang * b.lebar * b.tinggi))
}

func luasPersegi(luas int, status bool) interface{} {
	var persegiInterface interface{} = &persegi{luas, status}
	var load interface{}
	switch {
	case (persegiInterface.(*persegi).status == true):
		if persegiInterface.(*persegi).luas == 0 {
			load = fmt.Sprintf("Maaf anda belum menginput sisi dari persegi")
		} else {
			load = fmt.Sprintf("luas persegi dengan sisi %fcm adalah %fcm ", math.Sqrt(float64(persegiInterface.(*persegi).luas)), float64(persegiInterface.(*persegi).luas))
		}
	case (persegiInterface.(*persegi).status == false):
		if persegiInterface.(*persegi).luas == 0 {
			load = nil
		} else {
			load = fmt.Sprint(persegiInterface.(*persegi).luas)
		}
	}
	return load
}
func main() {
	//soal 1
	var bangunDatar hitungBangunDatar
	bangunDatar = segitigaSamaSisi{6, 12}
	fmt.Println("==== segitiga sama sisi")
	fmt.Println(bangunDatar.luas())
	fmt.Println(bangunDatar.keliling())

	bangunDatar = persegiPanjang{6, 12}
	fmt.Println("===== persegi panjang")
	fmt.Println(bangunDatar.keliling())
	fmt.Println(bangunDatar.luas())

	var bangunRuang hitungBangunRuang
	fmt.Println("=== bangun tabung")
	bangunRuang = tabung{float64(12), float64(6)}
	fmt.Println(bangunRuang.luasPermukaan())
	fmt.Println(bangunRuang.volume())
	fmt.Println("==== bangun balok")
	bangunRuang = balok{20, 20, 20}
	fmt.Println(bangunRuang.luasPermukaan())
	fmt.Println(bangunRuang.volume())
	// soal 2
	fmt.Println("=========")
	var phoneProperty phoneBrand
	phoneProperty = phone{"Samsung Galaxy Note 20", "Samsung Galaxy Note 20", 2020, []string{"Mystic Bronze", "Mystic White", "Mystic Black"}}
	fmt.Println(phoneProperty.print())
	//soal 3
	fmt.Println(luasPersegi(4, true))

	fmt.Println(luasPersegi(8, false))

	fmt.Println(luasPersegi(0, true))

	fmt.Println(luasPersegi(0, false))
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
	loadString = arrayToString(arrayload, "+")
	fmt.Printf("%s %s %d", prefix, loadString, sum)
	// tulis jawaban anda disini

}
func arrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}
