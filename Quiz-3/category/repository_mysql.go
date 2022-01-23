package category

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
	table          = "category"
	layoutDateTime = "2006-01-02 15:04:05"
)

func GetAll(ctx context.Context) ([]models.Category, error) {
	var categories []models.Category

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
		var category models.Category
		var CreatedAt, UpdatedAt string
		if err = rowQuery.Scan(&category.Id,
			&category.Name,
			&CreatedAt,
			&UpdatedAt); err != nil {
			return nil, err
		}

		category.CreatedAt, err = time.Parse(layoutDateTime, CreatedAt)
		if err != nil {
			log.Fatal(err)
		}
		category.UpdatedAt, err = time.Parse(layoutDateTime, UpdatedAt)
		if err != nil {
			log.Fatal(err)
		}
		categories = append(categories, category)
	}
	return categories, nil
}
func GetByCategory(ctx context.Context, id string) ([]models.Book, error) {
	var books []models.Book
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("cant connect to MySql", err)
	}

	// queryText := fmt.Sprintf("SELECT * FROM %v Order By created_at DESC",table)
	queryText := fmt.Sprintf("SELECT * FROM %v WHERE category_id=%v", "book", id)
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

func Insert(ctx context.Context, category models.Category) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to mysql", err)
	}
	queryText := fmt.Sprintf("INSERT INTO %v (name, created_at, updated_at) values('%v', NOW(), NOW())", table,
		category.Name,
	)
	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}
func Update(ctx context.Context, category models.Category, id string) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("UPDATE %v set name ='%s', updated_at = NOW() where id = %s",
		table,
		category.Name,
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
