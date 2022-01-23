package book

import (
	"BDS-Sanbercode-Golang-Batch-31/Quiz-3/helper"
	"BDS-Sanbercode-Golang-Batch-31/Quiz-3/models"
	"BDS-Sanbercode-Golang-Batch-31/Quiz-3/utils"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/julienschmidt/httprouter"
)

func GetBooks(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	books, err := GetAll(ctx)
	if err != nil {
		fmt.Println(err)
	}
	utils.ResponseJSON(w, books, http.StatusOK)
}

func PostCategory(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var bok models.Book
	if err := json.NewDecoder(r.Body).Decode(&bok); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}
	validation := ""

	_, errUri := url.ParseRequestURI(bok.ImageUrl)
	if errUri != nil {
		validation += "alamat gambar anda salah\n"
	}
	if !(bok.ReleaseYear >= 1980 && bok.ReleaseYear <= 2021) {
		validation += "alamat gambar anda salah\n"
	}
	if len(validation) > 0 {
		http.Error(w, validation, http.StatusBadRequest)
		return
	}
	bok.Thickness = helper.KonversiKetebalan(bok.TotalPage)

	if err := Insert(ctx, bok); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"status": "succesfully",
	}
	utils.ResponseJSON(w, res, http.StatusCreated)

}

func UpdateCategory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var bok models.Book
	if err := json.NewDecoder(r.Body).Decode(&bok); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}
	var idBook = ps.ByName("id")
	validation := ""

	_, errUri := url.ParseRequestURI(bok.ImageUrl)
	if errUri != nil {
		validation += "alamat gambar anda salah\n"
	}
	if !(bok.ReleaseYear >= 1980 && bok.ReleaseYear <= 2021) {
		validation += "alamat gambar anda salah\n"
	}
	if len(validation) > 0 {
		http.Error(w, validation, http.StatusBadRequest)
		return
	}
	bok.Thickness = helper.KonversiKetebalan(bok.TotalPage)

	if err := Update(ctx, bok, idBook); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"status": "succesfully",
	}
	utils.ResponseJSON(w, res, http.StatusCreated)
}

func DeleteBook(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var idbook = ps.ByName("id")
	if err := Delete(ctx, idbook); err != nil {
		kesalahan := map[string]string{
			"error": fmt.Sprintf("%v", err),
		}
		utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"status": "Succesfully",
	}
	utils.ResponseJSON(w, res, http.StatusOK)
}
