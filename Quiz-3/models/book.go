package models

import "time"

type (
	Book struct {
		Id          int       `json:"id"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		ImageUrl    string    `json:"image_url"`
		ReleaseYear int       `json:"release_year"`
		Price       string    `json:"price"`
		TotalPage   int       `json:"total_page"`
		Thickness   string    `json:"thickness"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
		Category_id int       `json:"category_id"`
	}
)
