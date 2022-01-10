package main

import (
	"fmt"
	"strings"
)

func main() {
	//soal 1
	panjang := 12
	lebar := 4
	tinggi := 8

	luas := luasPersegiPanjang(panjang, lebar)
	keliling := kelilingPersegiPanjang(panjang, lebar)
	volume := volumeBalok(panjang, lebar, tinggi)
	//soal 2
	fmt.Println(luas)
	fmt.Println(keliling)
	fmt.Println(volume)
	//soal 2
	john := introduce("John", "laki-laki", "penulis", "30")
	fmt.Println(john) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"

	sarah := introduce("Sarah", "perempuan", "model", "28")
	fmt.Println(sarah) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"
	//soal 3
	var buah = []string{"semangka", "jeruk", "melon", "pepaya"}

	var buahFavoritJohn = buahFavorit("John", buah...)

	fmt.Println(buahFavoritJohn)
	//soal 4
	var dataFilm = []map[string]string{}
	// buatlah closure function disini
	var tambahDataFilm = func(title string, jam string, genre string, tahun string) {
		dataFilm = append(dataFilm, map[string]string{"title": title, "jam": jam, "genre": genre, "tahun": tahun})
	}
	tambahDataFilm("LOTR", "2 jam", "action", "1999")
	tambahDataFilm("avenger", "2 jam", "action", "2019")
	tambahDataFilm("spiderman", "2 jam", "action", "2004")
	tambahDataFilm("juon", "2 jam", "horror", "2004")

	for _, item := range dataFilm {
		fmt.Println(item)
	}
}

// soal 3
func buahFavorit(name string, buah ...string) string {
	var loadStringArrays []string
	var loadArrayToString string
	for _, v := range buah {
		loadStringArrays = append(loadStringArrays, fmt.Sprintf("\"%s\"", v))
	}
	loadArrayToString = strings.Join(loadStringArrays[:], ", ")

	return fmt.Sprintf("halo nama saya %s dan buah favorit saya adalah %s", name, loadArrayToString)
}

//soal 2
func introduce(name string, sex string, job string, age string) (loadString string) {
	var status string
	if sex == "laki-laki" {
		status = "Pak"
	} else if sex == "perempuan" {
		status = "Bu"
	}
	loadString = fmt.Sprintf("%s %s adalah seorang %s yang berusia %s tahun", status, name, job, age)
	return
}

//soal 1
func luasPersegiPanjang(panjang int, lebar int) int {
	return panjang * lebar
}

func kelilingPersegiPanjang(panjang int, lebar int) int {
	return 2 * (panjang + lebar)
}
func volumeBalok(panjang int, lebar int, tinggi int) int {
	return panjang * lebar * tinggi
}
