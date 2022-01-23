package shape

import (
	"BDS-Sanbercode-Golang-Batch-31/Quiz-3/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func ShapeController(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	q := r.URL.Query()
	hasilFormula := make(chan float64)
	shapeFlag := ps.ByName("shape")
	switch shapeFlag {
	case "segitiga-sama-sisi":
		hitung := q.Get("hitung")
		alas, err := strconv.ParseFloat(q.Get("alas"), 64)
		if err != nil {
			utils.ResponseJSON(w, "alas harus angka", http.StatusBadRequest)
			return
		}
		tinggi, err := strconv.ParseFloat(q.Get("tinggi"), 64)
		if err != nil {
			utils.ResponseJSON(w, "tinggi harus angka", http.StatusBadRequest)
			return
		}
		var segitgaSS Shape = SegitigaSamaSisi{alas: alas, tinggi: tinggi}
		switch hitung {
		case "luas":
			go segitgaSS.Luas(hasilFormula)
		case "keliling":
			go segitgaSS.Keliling(hasilFormula)
		default:
			utils.ResponseJSON(w, "tentukan rumus", http.StatusBadRequest)
			return
		}
		responseString := fmt.Sprintf("%s dari Segitga sama sisi adalah %f", hitung, <-hasilFormula)
		utils.ResponseJSON(w, responseString, http.StatusOK)
	case "persegi":
		hitung := q.Get("hitung")
		sisi, err := strconv.ParseFloat(q.Get("sisi"), 64)
		if err != nil {
			utils.ResponseJSON(w, "sisi harus angka", http.StatusBadRequest)
			return
		}
		var Shape Shape = Persegi{sisi: sisi}
		switch hitung {
		case "luas":
			go Shape.Luas(hasilFormula)
		case "keliling":
			go Shape.Keliling(hasilFormula)
		default:
			utils.ResponseJSON(w, "tentukan rumus", http.StatusBadRequest)
			return
		}
		responseString := fmt.Sprintf("%s dari persegi adalah %f", hitung, <-hasilFormula)
		utils.ResponseJSON(w, responseString, http.StatusOK)
	case "persegi-panjang":
		hitung := q.Get("hitung")
		panjang, err := strconv.ParseFloat(q.Get("panjang"), 64)
		if err != nil {
			utils.ResponseJSON(w, "panjang harus angka", http.StatusBadRequest)
			return
		}
		lebar, err := strconv.ParseFloat(q.Get("lebar"), 64)
		if err != nil {
			utils.ResponseJSON(w, "lebar harus angka", http.StatusBadRequest)
			return
		}
		var Shape Shape = PersegiPanjang{panjang: panjang, lebar: lebar}
		switch hitung {
		case "luas":
			go Shape.Luas(hasilFormula)
		case "keliling":
			go Shape.Keliling(hasilFormula)
		default:
			utils.ResponseJSON(w, "tentukan rumus", http.StatusBadRequest)
			return
		}
		responseString := fmt.Sprintf("%s dari Persegi adalah %f", hitung, <-hasilFormula)
		utils.ResponseJSON(w, responseString, http.StatusOK)
	case "lingkaran":
		hitung := q.Get("hitung")
		jariJari, err := strconv.ParseFloat(q.Get("jariJari"), 64)
		if err != nil {
			utils.ResponseJSON(w, "jariJari harus angka", http.StatusBadRequest)
			return
		}
		var Shape Shape = lingkaran{jariJari: jariJari}
		switch hitung {
		case "luas":
			go Shape.Luas(hasilFormula)
		case "keliling":
			go Shape.Keliling(hasilFormula)
		default:
			utils.ResponseJSON(w, "tentukan rumus", http.StatusBadRequest)
			return
		}
		responseString := fmt.Sprintf("%s dari lingkaran adalah %f", hitung, <-hasilFormula)
		utils.ResponseJSON(w, responseString, http.StatusOK)
	case "jajar-genjang":
		hitung := q.Get("hitung")
		sisi, err := strconv.ParseFloat(q.Get("sisi"), 64)
		if err != nil {
			utils.ResponseJSON(w, "sisi harus angka", http.StatusBadRequest)
			return
		}
		alas, err := strconv.ParseFloat(q.Get("alas"), 64)
		if err != nil {
			utils.ResponseJSON(w, "alas harus angka", http.StatusBadRequest)
			return
		}
		tinggi, err := strconv.ParseFloat(q.Get("tinggi"), 64)
		if err != nil {
			utils.ResponseJSON(w, "tinggi harus angka", http.StatusBadRequest)
			return
		}
		var Shape Shape = jajarGenjang{sisi: sisi, alas: alas, tinggi: tinggi}
		switch hitung {
		case "luas":
			go Shape.Luas(hasilFormula)
		case "keliling":
			go Shape.Keliling(hasilFormula)
		default:
			utils.ResponseJSON(w, "tentukan rumus", http.StatusBadRequest)
			return
		}
		responseString := fmt.Sprintf("%s dari jajar genjang adalah %f", hitung, <-hasilFormula)
		utils.ResponseJSON(w, responseString, http.StatusOK)
	default:
		utils.ResponseJSON(w, "bentuk shape anda salah", http.StatusBadRequest)
	}
}


