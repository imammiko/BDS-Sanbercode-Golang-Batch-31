package shape

import "math"

type Shape interface {
	Keliling(ch chan<- float64)
	Luas(ch chan<- float64)
}
type SegitigaSamaSisi struct {
	alas, tinggi float64
}
type Persegi struct {
	sisi float64
}

type PersegiPanjang struct {
	panjang, lebar float64
}
type lingkaran struct {
	jariJari float64
}
type jajarGenjang struct {
	sisi, alas, tinggi float64
}

func (s SegitigaSamaSisi) Luas(ch chan<- float64) {
	luas := 0.5 * s.alas * s.tinggi
	ch <- luas
}

func (s SegitigaSamaSisi) Keliling(ch chan<- float64) {
	keliling := s.alas * 3
	ch <- keliling
}

func (p Persegi) Luas(ch chan<- float64) {
	luas := math.Pow(p.sisi, 2)
	ch <- luas
}

func (p Persegi) Keliling(ch chan<- float64) {
	keliling := 4 * p.sisi
	ch <- keliling
}

func (p PersegiPanjang) Luas(ch chan<- float64) {
	luas := p.panjang * p.lebar
	ch <- luas
}
func (p PersegiPanjang) Keliling(ch chan<- float64) {
	keliling := 2*p.lebar + 2*p.panjang
	ch <- keliling
}
func (l lingkaran) Luas(ch chan<- float64) {
	luas := math.Pi * l.jariJari * l.jariJari
	ch <- luas
}

func (l lingkaran) Keliling(ch chan<- float64) {
	keliling := math.Pi * 2 * l.jariJari
	ch <- keliling
}
func (j jajarGenjang) Luas(ch chan<- float64) {
	luas := j.sisi * j.tinggi
	ch <- luas
}
func (j jajarGenjang) Keliling(ch chan<- float64) {
	keliling := 2*j.alas + 2*j.tinggi
	ch <- keliling
}
