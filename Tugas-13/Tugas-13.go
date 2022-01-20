package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type NilaiMahasiswa struct {
	// Nama, MataKuliah, IndeksNilai string
	Nama        string `json:"nama"`
	MataKuliah  string `json:"mataKuliah"`
	IndeksNilai string `json:"indeksNilai"`
	// Nilai, ID                     uint
	Nilai uint `json:"nilai"`
	ID    uint `json:"id"`
}
type Mahasiswas struct {
	mahasiswas []NilaiMahasiswa
}

// Fungi Log yang berguna sebagai middleware
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uname, pwd, ok := r.BasicAuth()
		if !ok {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Username atau Password tidak boleh kosong"))
			return
		}

		if uname == "admin" && pwd == "admin" {
			next.ServeHTTP(w, r)
			return
		}
		w.Write([]byte("Username atau Password tidak sesuai"))
		return
	})
}

// Fungsi GetMovie untuk mengampilkan text string di browser
func (m *Mahasiswas) PostMahasiswa(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var TakeNilaiMahasiswa NilaiMahasiswa
	if r.Method == "POST" {
		if r.Header.Get("Content-Type") == "application/json" {
			// parse dari json
			decodeJSON := json.NewDecoder(r.Body)
			if err := decodeJSON.Decode(&TakeNilaiMahasiswa); err != nil {
				log.Fatal(err)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
			}

		} else {
			// Nama        string `json:"nama"`
			// MataKuliah  string `json:"mataKuliah"`
			// IndeksNilai string `json:"indeksNilai"`
			// // Nilai, ID                     uint
			// Nilai uint `json:"nilai"`
			// ID    uint `json:"id"`

			nama := r.PostFormValue("nama")
			mataKuliah := r.PostFormValue("mataKuliah")
			indeksNilai := r.PostFormValue("indeksNilai")
			getId := r.PostFormValue("id")
			getNilai := r.PostFormValue("nilai")
			id, _ := strconv.Atoi(getId)
			nilai, _ := strconv.Atoi(getNilai)
			TakeNilaiMahasiswa = NilaiMahasiswa{
				Nama:        nama,
				MataKuliah:  mataKuliah,
				IndeksNilai: indeksNilai,
				Nilai:       uint(nilai),
				ID:          uint(id),
			}
		}
		// fmt.Println(TakeNilaiMahasiswa)
		addMahasiswa(m, TakeNilaiMahasiswa)

	} else if r.Method == "GET" {
		fmt.Println(m)
		DataMahasiswas, err := json.Marshal(m.mahasiswas)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(DataMahasiswas)
		return
	}
}

func addMahasiswa(allDataMahasiswa *Mahasiswas, getDataMahasiswa NilaiMahasiswa) {
	conditionIndeksNilai(&getDataMahasiswa)
	putId(allDataMahasiswa, &getDataMahasiswa)
	allDataMahasiswa.mahasiswas = append(allDataMahasiswa.mahasiswas, getDataMahasiswa)
	fmt.Println(allDataMahasiswa, "line 88")
}
func conditionIndeksNilai(aDataMahasiswa *NilaiMahasiswa) {
	if aDataMahasiswa.Nilai >= 80 {
		aDataMahasiswa.IndeksNilai = "A"
	} else if aDataMahasiswa.Nilai >= 70 && aDataMahasiswa.Nilai < 80 {
		aDataMahasiswa.IndeksNilai = "B"
	} else if aDataMahasiswa.Nilai >= 60 && aDataMahasiswa.Nilai < 70 {
		aDataMahasiswa.IndeksNilai = "C"
	} else if aDataMahasiswa.Nilai >= 50 && aDataMahasiswa.Nilai < 60 {
		aDataMahasiswa.IndeksNilai = "D"
	} else if aDataMahasiswa.Nilai < 50 {
		aDataMahasiswa.IndeksNilai = "E"
	}
}
func putId(allDataMahasiswa *Mahasiswas, aDataMahasiswa *NilaiMahasiswa) {
	var maxValueId = 0
	if len(allDataMahasiswa.mahasiswas) == 0 {
		aDataMahasiswa.ID = 1
	} else {

		for _, v := range allDataMahasiswa.mahasiswas {
			if v.ID > uint(maxValueId) {
				maxValueId = int(v.ID)
			}
		}
		aDataMahasiswa.ID = uint(maxValueId) + 1
	}

}

func main() {
	// konfigurasi server
	var nilaiNilaiMahasiswa Mahasiswas
	server := &http.Server{
		Addr: ":8080",
	}

	// routing
	http.Handle("/mahasiswa", Auth(http.HandlerFunc(nilaiNilaiMahasiswa.PostMahasiswa)))

	// jalankan server
	fmt.Println("server running at http://localhost:8080")
	fmt.Printf("%p", &nilaiNilaiMahasiswa)

	server.ListenAndServe()
}
