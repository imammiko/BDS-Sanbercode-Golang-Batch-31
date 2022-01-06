package main

import "fmt"

func main() {
	//1

	for i := 1; i <= 20; i++ {
		var attach string
		if (i%3 == 0) && (i%2 != 0) {
			attach = "I Love Coding"
		} else if i%2 == 0 {
			attach = "berkualitas"
		} else {
			attach = "santai"
		}
		fmt.Printf("%d - %s \n", i, attach)
	}

	//2
	var loadString string
	for i := 1; i <= 7; i++ {
		for j := 0; j < i; j++ {
			loadString += "#"
		}
		loadString += "\n"
	}
	fmt.Println(loadString)

	//3
	var kalimat = [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}
	fmt.Println(kalimat[2:])
	//4
	var sayuran = []string{}
	sayuran = append(sayuran, "Bayam", "Buncis", "Kangkung", "Kubis", "Seledri", "Tauge", "Timun")
	for i, v := range sayuran {
		fmt.Printf("%d. %s \n", i, v)
	}
	//5
	var satuan = map[string]int{
		"Panjang": 7,
		"Lebar":   4,
		"Tinggi":  6,
	}
	satuan["Volume Balok"] = satuan["Panjang"] * satuan["Lebar"] * satuan["Tinggi"]
	for k, v := range satuan {

		fmt.Printf("%s = %d \n", k, v)
	}

}
