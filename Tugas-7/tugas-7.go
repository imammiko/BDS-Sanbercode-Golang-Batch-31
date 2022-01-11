package main

import (
	"fmt"
)

type Buah struct {
	Nama       string
	Warna      string
	AdaBijinya bool
	Harga      int
}
type segitiga struct {
	alas, tinggi int
}

type persegi struct {
	sisi int
}

type persegiPanjang struct {
	panjang, lebar int
}
type phone struct {
	name, brand string
	year        int
	colors      []string
}
type movie struct {
	title    string
	genre    string
	duration int
	year     int
}

func main() {
	// soal 1

	fmt.Println("Nama   Warna   Ada Bijinya    Harga")
	Nanas := Buah{"nanas", "Kuning", false, 9000}
	Nanas.print()
	Jeruk := Buah{"Jeruk", "Warna", true, 8000}
	Jeruk.print()
	Semangka := Buah{"Semangka", "Hijau & Merah", true, 10000}
	Semangka.print()
	Pisang := Buah{"Pisang", "Kuning", false, 5000}
	Pisang.print()
	// soal 2
	luasSegitiga := segitiga{1, 7}
	fmt.Println(luasSegitiga.luas())
	luasPersegi := persegi{2}
	fmt.Println(luasPersegi.luas())
	luasPersegiPanjang := persegiPanjang{3, 4}
	fmt.Println(luasPersegiPanjang.luas())
	//soal 3
	phone := phone{"nokia 6600", "nokia", 2005, []string{}}
	phone.TambahColor("black", "yellow")
	fmt.Println(phone)

	var dataFilm = []movie{}

	tambahDataFilm("LOTR", 120, "action", 1999, &dataFilm)
	tambahDataFilm("avenger", 120, "action", 2019, &dataFilm)
	tambahDataFilm("spiderman", 120, "action", 2004, &dataFilm)
	tambahDataFilm("juon", 120, "horror", 2004, &dataFilm)
	for i := 0; i < len(dataFilm); i++ {
		var duration float64
		duration = float64(dataFilm[i].duration / 60)
		fmt.Printf("%d. ", i+1)
		fmt.Printf("title: %s\n", dataFilm[i].title)
		fmt.Printf("duration: %f jam\n", duration)
		fmt.Printf("year: %d", dataFilm[i].year)
	}
	fmt.Println(dataFilm)
}

func tambahDataFilm(title string, jam int, genre string, tahun int, dataFilm *[]movie) {

	*dataFilm = append(*dataFilm, movie{title, genre, jam, tahun})
}

func (b Buah) print() {
	var StatusAda string
	if b.AdaBijinya {
		StatusAda = "Ada"
	} else {
		StatusAda = "Tidak Ada"
	}
	fmt.Printf("%s %s %s %d\n", b.Nama, b.Warna, StatusAda, b.Harga)

}
func (s segitiga) luas() float64 {
	return (float64(s.alas) * float64(s.tinggi)) / 2
}
func (s persegi) luas() float64 {
	return float64(s.sisi) * float64(s.sisi)
}
func (p persegiPanjang) luas() float64 {
	return float64(p.lebar) * float64(p.panjang)
}

func (p *phone) TambahColor(warna ...string) {
	p.colors = append(p.colors, warna...)
	fmt.Println(p.colors)
}
