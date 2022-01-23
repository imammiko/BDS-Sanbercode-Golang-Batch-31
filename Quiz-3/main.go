package main

import (
	"BDS-Sanbercode-Golang-Batch-31/Quiz-3/book"
	"BDS-Sanbercode-Golang-Batch-31/Quiz-3/category"
	"BDS-Sanbercode-Golang-Batch-31/Quiz-3/shape"

	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/bangun-datar/:shape", shape.ShapeController)
	router.GET("/categories", category.GetCategory)
	router.POST("/categories", BasicAuth(category.PostCategory))
	router.PUT("/categories/:id", BasicAuth(category.UpdateCategory))
	router.DELETE("/categories/:id", BasicAuth(category.DeleteCategory))
	router.GET("/categories/:id/books", category.GetByCategoryHandle)

	router.GET("/books", book.GetBooks)
	router.POST("/books", BasicAuth(book.PostCategory))
	router.PUT("/books/:id", BasicAuth(book.UpdateCategory))
	router.DELETE("/books/:id", BasicAuth(book.DeleteBook))

	fmt.Println("Server Running at Port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
func BasicAuth(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// Get the Basic Authentication credentials
		user, password, hasAuth := r.BasicAuth()

		if hasAuth && (user == "admin" && password == "password") || (user == "editor" && password == "secret") || (user == "trainer" && password == "rahasia") {
			// Delegate request to the given handle
			h(w, r, ps)
		} else {
			// Request Basic Authentication otherwise
			w.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}
	}
}
