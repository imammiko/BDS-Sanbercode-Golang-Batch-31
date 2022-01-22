package movie

import (
	"BDS-Sanbercode-Golang-Batch-31/example14/config"
	"BDS-Sanbercode-Golang-Batch-31/example14/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
)

const (
	table          = "movie"
	layoutDateTime = "2006-01-02 15:04:05"
)

func GetAll(ctx context.Context) ([]models.Movie, error) {
	var movies []models.Movie
	var createdAt, updatedAt string
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("cant connect to MySql", err)
	}
	queryText := fmt.Sprintf("SELECT * FROM %v Order By created_at DESC", table)
	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}
	for rowQuery.Next() {
		var movie models.Movie
		var CreatedAt, UpdatedAt string
		if err = rowQuery.Scan(&movie.ID,
			&movie.Title,
			&movie.Year,
			&CreatedAt,
			&UpdatedAt); err != nil {
			return nil, err
		}

		movie.CreatedAt, err = time.Parse(layoutDateTime, createdAt)
		if err != nil {
			log.Fatal(err)
		}
		movie.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	return movies, nil
}

func Insert(ctx context.Context, movie models.Movie) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to mysql", err)
	}
	queryText := fmt.Sprintf("INSERT INTO %v (title, year, created_at, updated_at) values('%v',%v, NOW(), NOW())", table,
		movie.Title,
		movie.Year)
	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}
func Update(ctx context.Context, movie models.Movie, id string) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("UPDATE %v set title ='%s', year = %d, updated_at = NOW() where id = %s",
		table,
		movie.Title,
		movie.Year,
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
