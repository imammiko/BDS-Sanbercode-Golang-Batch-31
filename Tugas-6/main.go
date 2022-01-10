package main

import (
	"fmt"
	"math"
)

func main() {
	var luasLigkaran float64
	var kelilingLingkaran float64
	//1
	luasKelilingLingkaran(&luasLigkaran, &kelilingLingkaran, 3)
	fmt.Println(luasLigkaran, kelilingLingkaran)
	//2
	//soal 2

	var sentence string
	introduce(&sentence, "John", "laki-laki", "penulis", "30")

	fmt.Println(sentence) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"
	introduce(&sentence, "Sarah", "perempuan", "model", "28")

	fmt.Println(sentence) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"

	//soal 3
	var buah = []string{}
	loadForLoop(&buah, "Jeruk", "Semangka", "Mangga", "Strawberry", "Durian", "Manggis", "Alpukat")
	for k, v := range buah {
		fmt.Printf("%d. %s\n", k+1, v)
	}
	// soal4
	var tambahDataFilm = func(title string, jam string, genre string, tahun string, dataFilm *[]map[string]string) {
		*dataFilm = append(*dataFilm, map[string]string{"title": title, "jam": jam, "genre": genre, "tahun": tahun})
	}
	var dataFilm = []map[string]string{}

	tambahDataFilm("LOTR", "2 jam", "action", "1999", &dataFilm)
	tambahDataFilm("avenger", "2 jam", "action", "2019", &dataFilm)
	tambahDataFilm("spiderman", "2 jam", "action", "2004", &dataFilm)
	tambahDataFilm("juon", "2 jam", "horror", "2004", &dataFilm)
	for k, v := range dataFilm {
		fmt.Print(k + 1)
		fmt.Print(".")
		for k, v := range v {
			fmt.Printf("%s:%s\n", k, v)
		}

		fmt.Println("")
	}
}
func loadForLoop(buah *[]string, arrayLoad ...string) {

	*buah = arrayLoad
}

func luasKelilingLingkaran(luas *float64, keliling *float64, radius float64) {
	*luas = float64(math.Pi * math.Pow(radius, 2))
	*keliling = float64(math.Pi * radius * 2)

}

func introduce(sentence *string, name string, sex string, job string, age string) {
	var status string
	if sex == "laki-laki" {
		status = "Pak"
	} else if sex == "perempuan" {
		status = "Bu"
	}
	*sentence = fmt.Sprintf("%s %s adalah seorang %s yang berusia %s tahun", status, name, job, age)
	return
}
