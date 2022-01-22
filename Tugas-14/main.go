package main

import (
	"BDS-Sanbercode-Golang-Batch-31/Tugas-14/mahasiswa"
	"BDS-Sanbercode-Golang-Batch-31/Tugas-14/models"
	"BDS-Sanbercode-Golang-Batch-31/Tugas-14/utils"

	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/mahasiswa", GetMahasiswa)
	router.POST("/mahasiswa", PostMahasiswa)
	router.PUT("/mahasiswa/:id", UpdateMahasiswa)
	router.DELETE("/mahasiswa/:id", DeleteMahasiswa)
	fmt.Println("server Runnning at Port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func DeleteMahasiswa(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var idMahasiswa = ps.ByName("id")
	if err := mahasiswa.Delete(ctx, idMahasiswa); err != nil {
		kesalahan := map[string]string{
			"error": fmt.Sprintf("%v", err),
		}
		utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"status": "succesfully",
	}
	utils.ResponseJSON(w, res, http.StatusOK)
}

func UpdateMahasiswa(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var mahasi models.NilaiMahasiswa
	if err := json.NewDecoder(r.Body).Decode(&mahasi); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}
	validation := ""

	if mahasi.Nilai < 0 {
		if len(validation) == 0 {
			validation = "nilai tidak boleh 0"
		} else {
			validation += " dan nilai tidak boleh 0"
		}
	} else if mahasi.Nilai > 100 {
		if len(validation) == 0 {
			validation = "nilai tidak boleh lebih dari 100"
		} else {
			validation += "dan nilai tidak boleh lebih dari 100"
		}
	}

	if len(validation) > 0 {
		http.Error(w, validation, http.StatusBadRequest)
		return
	}

	mahasi.IndeksNilai = conditionIndeksNilai(int(mahasi.Nilai))
	fmt.Println(mahasi.MataKuliah)
	var idMahasiswa = ps.ByName("id")
	if err := mahasiswa.Update(ctx, mahasi, idMahasiswa); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"status": "succesfully",
	}
	utils.ResponseJSON(w, res, http.StatusCreated)

}
func GetMahasiswa(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	//  mahasiswas,err :=
	mahasiswa, err := mahasiswa.GetAll(ctx)
	if err != nil {
		fmt.Println(err)
	}
	utils.ResponseJSON(w, mahasiswa, http.StatusOK)

}

func PostMahasiswa(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var mahasis models.NilaiMahasiswa
	if err := json.NewDecoder(r.Body).Decode(&mahasis); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}
	validation := ""

	if mahasis.Nilai < 0 {
		if len(validation) == 0 {
			validation = "nilai tidak boleh 0"
		} else {
			validation += " dan nilai tidak boleh 0"
		}
	} else if mahasis.Nilai > 100 {
		if len(validation) == 0 {
			validation = "nilai tidak boleh lebih dari 100"
		} else {
			validation += "dan nilai tidak boleh lebih dari 100"
		}
	}

	if len(validation) > 0 {
		http.Error(w, validation, http.StatusBadRequest)
		return
	}

	mahasis.IndeksNilai = conditionIndeksNilai(int(mahasis.Nilai))
	if err := mahasiswa.Insert(ctx, mahasis); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"status": "succesfully",
	}
	utils.ResponseJSON(w, res, http.StatusCreated)

}
func conditionIndeksNilai(aDataMahasiswa int) string {
	var indeks string
	if aDataMahasiswa >= 80 {
		indeks = "A"
	} else if aDataMahasiswa >= 70 && aDataMahasiswa < 80 {
		indeks = "B"
	} else if aDataMahasiswa >= 60 && aDataMahasiswa < 70 {
		indeks = "C"
	} else if aDataMahasiswa >= 50 && aDataMahasiswa < 60 {
		indeks = "D"
	} else if aDataMahasiswa < 50 {
		indeks = "E"
	}
	return indeks
}
