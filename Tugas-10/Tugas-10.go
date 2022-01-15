package main

import (
	"fmt"
	"math"
)

type segitgaSamaSisi struct {
	keliling int
	status   bool
}

func kelilingSegitigaSamaSisi(keliling int, status bool) string {
	var segitgaInterface interface{} = &segitgaSamaSisi{keliling, status}
	switch {
	case (segitgaInterface.(*segitgaSamaSisi).status == true):
		if segitgaInterface.(*segitgaSamaSisi).keliling == 0 {
			panic("Maaf anda belum menginput sisi dari segitiga sama sisi")
		} else {
			load := fmt.Sprintf("keliling segitiga sama sisinya dengan sisi %f cm adalah %d cm", math.Pow(float64(segitgaInterface.(*segitgaSamaSisi).keliling), 1.0/3), segitgaInterface.(*segitgaSamaSisi).keliling)
			return load

		}
	case (segitgaInterface.(*segitgaSamaSisi).status == false):
		if segitgaInterface.(*segitgaSamaSisi).keliling == 0 {
			panic("Maaf anda belum menginput sisi dari segitiga sama sisi")
		} else {
			load := fmt.Sprintf("%f", math.Pow(float64(segitgaInterface.(*segitgaSamaSisi).keliling), 1.0/3))
			return load
		}
	}
	return ""
}

func main() {
	//soal1
	defer deferFunc("Golang Backend Development", 2021)
	//soal2
	fmt.Println(kelilingSegitigaSamaSisi(4, true))

	fmt.Println(kelilingSegitigaSamaSisi(8, false))

	fmt.Println(kelilingSegitigaSamaSisi(0, true))

	fmt.Println(kelilingSegitigaSamaSisi(0, false))
}
func deferFunc(prefix string, year int) {
	fmt.Printf("%s %d", prefix, year)

}
