package main

import (
	"fmt"
	"math"
	"sort"
	"sync"
	"time"
)

type phone struct {
	brandPhone []string
}

func (p *phone) Sort() {
	var wg sync.WaitGroup

	var phones = []string{"Xiaomi", "Asus", "Iphone", "Samsung", "Oppo", "Realme", "Vivo"}
	p.brandPhone = append(p.brandPhone, phones...)

	sort.Strings(p.brandPhone)
	for i, v := range p.brandPhone {
		wg.Add(1)
		go print(i+1, v)
		time.Sleep(500 * time.Millisecond)
		wg.Done()
	}
	wg.Wait()

}

func print(index int, value string) {
	fmt.Printf("%d. %s \n", index, value)
}

func getMovies(moviesChannel chan string, movies ...string) {
	for i, v := range movies {
		moviesChannel <- fmt.Sprintf("%d. %s", i+1, v)
	}
	close(moviesChannel)
}

func keliling(rujiRuji int, ch chan float64) {
	kelilingValue := 2 * math.Pi * float64(rujiRuji)
	ch <- kelilingValue
}
func luas(rujiRuji int, ch chan float64) {
	luas := math.Phi * math.Pow(float64(rujiRuji), 2)
	ch <- luas
}
func volume(rujiRuji int, tinggi int, ch chan float64) {
	volume := math.Phi * math.Pow(float64(rujiRuji), 2) * float64(tinggi)
	ch <- volume
}

func kelilingPersegiPanjang(panjang int, lebar int, ch chan int) {
	keliling := 2 * (panjang * lebar)
	ch <- keliling
}
func luasPersegiPanjang(panjang int, lebar int, ch chan int) {
	luas := panjang * lebar
	ch <- luas
}
func volumePersegiPanjang(panjang int, lebar int, tinggi int, ch chan int) {
	volume := panjang * lebar * tinggi
	ch <- volume
}
func main() {
	//soal 1
	merekPhone := phone{}
	merekPhone.Sort()
	// soal 2
	var movies = []string{"Harry Potter", "LOTR", "SpiderMan", "Logan", "Avengers", "Insidious", "Toy Story"}

	moviesChannel := make(chan string)

	go getMovies(moviesChannel, movies...)

	for value := range moviesChannel {
		fmt.Println(value)
	}
	// soal 3
	jari := []int{8, 14, 20}
	var tinggiTabung int
	tinggiTabung = 10
	value := make(chan float64)
	go keliling(jari[0], value)
	fmt.Println(<-value)
	go keliling(jari[1], value)
	fmt.Println(<-value)
	go keliling(jari[2], value)
	fmt.Println(<-value)
	go luas(jari[0], value)
	fmt.Println(<-value)
	go luas(jari[1], value)
	fmt.Println(<-value)
	go luas(jari[2], value)
	fmt.Println(<-value)
	go volume(jari[0], tinggiTabung, value)
	fmt.Println(<-value)
	go volume(jari[1], tinggiTabung, value)
	fmt.Println(<-value)
	go volume(jari[2], tinggiTabung, value)
	fmt.Println(<-value)
	//soal4
	panjang := 4
	lebar := 5
	tinggi := 6

	var keliling = make(chan int)
	go kelilingPersegiPanjang(panjang, lebar, keliling)
	var luas = make(chan int)
	go luasPersegiPanjang(panjang, lebar, luas)
	var volume = make(chan int)
	go volumePersegiPanjang(panjang, lebar, tinggi, volume)

	for i := 0; i < 3; i++ {
		select {
		case keliling := <-keliling:
			fmt.Println("keliling balok:", keliling)
		case luas := <-luas:
			fmt.Println("luas balok:", luas)
		case volume := <-volume:
			fmt.Println("volume balok", volume)
		}
	}
}
