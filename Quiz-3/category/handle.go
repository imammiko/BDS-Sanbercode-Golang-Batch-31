package category

import (
	"BDS-Sanbercode-Golang-Batch-31/Quiz-3/models"
	"BDS-Sanbercode-Golang-Batch-31/Quiz-3/utils"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func GetCategory(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	category, err := GetAll(ctx)
	if err != nil {
		fmt.Println(err)
	}
	utils.ResponseJSON(w, category, http.StatusOK)
}
func GetByCategoryHandle(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var idCategory = ps.ByName("id")
	category, err := GetByCategory(ctx, idCategory)
	if err != nil {
		fmt.Println(err)
	}
	utils.ResponseJSON(w, category, http.StatusOK)
}

func PostCategory(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var mov models.Category
	if err := json.NewDecoder(r.Body).Decode(&mov); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}
	if err := Insert(ctx, mov); err != nil {
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
	var cat models.Category
	if err := json.NewDecoder(r.Body).Decode(&cat); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}
	var idCategory = ps.ByName("id")
	if err := Update(ctx, cat, idCategory); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"status": "succesfully",
	}
	utils.ResponseJSON(w, res, http.StatusCreated)
}

func DeleteCategory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var idCategory = ps.ByName("id")
	if err := Delete(ctx, idCategory); err != nil {
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
