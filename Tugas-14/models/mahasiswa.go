package models

import "time"

type NilaiMahasiswa struct {
	// Nama, MataKuliah, IndeksNilai string
	Nama        string `json:"nama"`
	MataKuliah  string `json:"mataKuliah"`
	IndeksNilai string `json:"indeksNilai"`
	// Nilai, ID                     uint
	Nilai     uint      `json:"nilai"`
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
