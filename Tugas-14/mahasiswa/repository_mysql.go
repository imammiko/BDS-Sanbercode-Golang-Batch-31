package mahasiswa

import (
	"BDS-Sanbercode-Golang-Batch-31/Tugas-14/config"
	"BDS-Sanbercode-Golang-Batch-31/Tugas-14/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
)

const (
	table          = "mahasiswa"
	layoutDateTime = "2006-01-02 15:04:05"
)

func GetAll(ctx context.Context) ([]models.NilaiMahasiswa, error) {
	var mahasiswas []models.NilaiMahasiswa
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("cant connect to mysql", err)
	}
	queryText := fmt.Sprintf("SELECT * FROM %v Order By created_at DESC", table)
	rowQuery, err := db.QueryContext(ctx, queryText)
	if err != nil {
		log.Fatal(err)
	}
	for rowQuery.Next() {
		var mahasiswa models.NilaiMahasiswa
		var createdAt, updatedAt string
		if err = rowQuery.Scan(
			&mahasiswa.ID,
			&mahasiswa.Nama,
			&mahasiswa.MataKuliah,
			&mahasiswa.IndeksNilai,
			&mahasiswa.Nilai,
			&createdAt,
			&updatedAt,
		); err != nil {
			return nil, err
		}
		mahasiswa.CreatedAt, err = time.Parse(layoutDateTime, createdAt)
		if err != nil {
			log.Fatal(err)
		}
		mahasiswa.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)
		if err != nil {
			log.Fatal()
		}
		mahasiswas = append(mahasiswas, mahasiswa)
	}
	return mahasiswas, nil
}

func Insert(ctx context.Context, mahasiswa models.NilaiMahasiswa) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to mysql", err)
	}
	queryText := fmt.Sprintf("INSERT INTO %v (nama, mataKuliah, indeksNilai, nilai,created_at,updated_at) values('%v','%v','%v',%v, NOW(), NOW())", table,
		mahasiswa.Nama, mahasiswa.MataKuliah, mahasiswa.IndeksNilai, mahasiswa.Nilai)
	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}

func Update(ctx context.Context, mahasiswa models.NilaiMahasiswa, id string) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("can 't connect to Mysql", err)
	}

	queryText := fmt.Sprintf("UPDATE %v set nama ='%s', mataKuliah = '%s',indeksNilai= '%s',nilai='%v', updated_at = NOW() where id = '%v'",
		table,
		mahasiswa.Nama,
		mahasiswa.MataKuliah,
		mahasiswa.IndeksNilai,
		mahasiswa.Nilai,
		id,
	)
	_, err = db.ExecContext(ctx, queryText)
	if err != nil {
		return err
	}
	return nil
}

func Delete(ctx context.Context, id string) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("can't connect to mysql", err)
	}
	queryText := fmt.Sprintf("DELETE FROM %v where id = %s", table, id)

	s, err := db.ExecContext(ctx, queryText)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	check, err := s.RowsAffected()
	fmt.Println(check)
	if check == 0 {
		return errors.New("id tidak ada")
	}

	if err != nil {
		fmt.Println(err.Error())
	}

	return nil
}
