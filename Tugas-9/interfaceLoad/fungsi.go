package interfaceload

import (
	"fmt"
	"math"
	"strings"
)

type SegitigaSamaSisi struct {
	Alas, Tinggi int
}

type PersegiPanjang struct {
	Panjang, Lebar int
}

type Tabung struct {
	JariJari, Tinggi float64
}

type Balok struct {
	Panjang, Lebar, Tinggi int
}

type Phone struct {
	Name, Brand string
	Year        int
	Colors      []string
}
type Persegi struct {
	Luas   int
	Status bool
}
type HitungBangunDatar interface {
	Luas() int
	Keliling() int
}

type HitungBangunRuang interface {
	Volume() float64
	LuasPermukaan() float64
}

type PhoneBrand interface {
	Print() string
}

func (s SegitigaSamaSisi) Luas() int {
	return s.Alas * s.Tinggi / 2
}

func (p Phone) Print() string {
	var loadStringPhone string
	loadStringPhone += fmt.Sprintf("name: %s\n", p.Name)
	loadStringPhone += fmt.Sprintf("brand: %s\n", p.Brand)
	loadStringPhone += fmt.Sprintf("year: %d\n", p.Year)
	loadStringPhone += fmt.Sprintf("colors: %s\n", strings.Join(p.Colors, ","))
	return loadStringPhone
}

func (s SegitigaSamaSisi) Keliling() int {
	var pytagoras float64
	var alas float64
	alas = float64(s.Alas) / 2
	pytagoras = math.Sqrt(math.Pow(alas, 2) + math.Pow(float64(s.Tinggi), 2))

	return (int(pytagoras) * 2) + s.Alas
}

func (p PersegiPanjang) Luas() int {
	return p.Panjang * p.Lebar
}

func (p PersegiPanjang) Keliling() int {
	return 2*p.Panjang + 2*p.Lebar
}
func (t Tabung) Volume() float64 {
	return math.Pi * math.Pow(t.JariJari, 2) * t.Tinggi
}
func (t Tabung) LuasPermukaan() float64 {
	return 2 * math.Pi * t.JariJari * (t.JariJari + t.Tinggi)
}
func (b Balok) Volume() float64 {
	return float64(b.Panjang * b.Lebar * b.Tinggi)
}
func (b Balok) LuasPermukaan() float64 {
	return float64(2 * (b.Panjang * b.Lebar * b.Tinggi))
}

func LuasPersegi(luas int, status bool) interface{} {
	var persegiInterface interface{} = &Persegi{luas, status}
	var load interface{}
	switch {
	case (persegiInterface.(*Persegi).Status == true):
		if persegiInterface.(*Persegi).Luas == 0 {
			load = fmt.Sprintf("Maaf anda belum menginput sisi dari persegi")
		} else {
			load = fmt.Sprintf("luas persegi dengan sisi %fcm adalah %fcm ", math.Sqrt(float64(persegiInterface.(*Persegi).Luas)), float64(persegiInterface.(*Persegi).Luas))
		}
	case (persegiInterface.(*Persegi).Status == false):
		if persegiInterface.(*Persegi).Luas == 0 {
			load = nil
		} else {
			load = fmt.Sprint(persegiInterface.(*Persegi).Luas)
		}
	}
	return load
}
func ArrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}
