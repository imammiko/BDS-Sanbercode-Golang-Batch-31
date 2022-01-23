package book

import (
	"BDS-Sanbercode-Golang-Batch-31/Quiz-3/config"
	"BDS-Sanbercode-Golang-Batch-31/Quiz-3/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
)

const (
	table          = "book"
	layoutDateTime = "2006-01-02 15:04:05"
)

func GetAll(ctx context.Context) ([]models.Book, error) {
	var books []models.Book
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
		var book models.Book
		var CreatedAt, UpdatedAt string
		if err = rowQuery.Scan(&book.Id,
			&book.Title,
			&book.Description,
			&book.ImageUrl,
			&book.ReleaseYear,
			&book.Price,
			&book.TotalPage,
			&book.Thickness,
			&CreatedAt,
			&UpdatedAt,
			&book.Category_id); err != nil {
			return nil, err
		}

		book.CreatedAt, err = time.Parse(layoutDateTime, CreatedAt)
		if err != nil {
			log.Fatal(err)
		}
		book.UpdatedAt, err = time.Parse(layoutDateTime, UpdatedAt)
		if err != nil {
			log.Fatal(err)
		}
		books = append(books, book)
	}
	return books, nil
}

func Insert(ctx context.Context, book models.Book) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to mysql", err)
	}
	queryText := fmt.Sprintf("INSERT INTO %v (title, description,image_url,release_year,price,total_page,thickness,category_id,created_at, updated_at) values('%v','%v', '%v','%v','%v','%v','%v','%v',NOW(), NOW())", table,
		book.Title, book.Description, book.ImageUrl, book.ReleaseYear, book.Price, book.TotalPage, book.Thickness, book.Category_id,
	)
	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}

func Update(ctx context.Context, book models.Book, id string) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("UPDATE %v set title ='%s',description ='%s',image_url ='%s',release_year ='%d', price ='%s',total_page ='%d',thickness ='%s',category_id ='%d',updated_at = NOW() where id = %s",
		table,
		book.Title,
		book.Description,
		book.ImageUrl,
		book.ReleaseYear,
		book.Price,
		book.TotalPage,
		book.Thickness,
		book.Category_id,
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
