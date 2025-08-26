package models

import "time"

type Book struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	AuthorID      uint      `json:"author_id"`
	Title         string    `gorm:"type:varchar(120);not null" json:"title"`
	Description   string    `gorm:"type:text;not null" json:"description"`
	PublishedYear string    `gorm:"type:varchar(4);not null" json:"published_year"`
	Published     string    `gorm:"type:varchar(120);not null" json:"published"`
	Pages         string    `gorm:"type:varchar(5);not null" json:"pages"`
	ISBN          string    `gorm:"type:varchar(70);not null" json:"isbn"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`

	Author Author `gorm:"foreignKey:AuthorID" json:"author"`
}

type BookResponse struct {
	ID            uint               `json:"id"`
	AuthorID      uint               `json:"-"`
	Author        AuthorBookResponse `gorm:"foreignKey:AuthorID" json:"author"`
	Title         string             `json:"title"`
	Description   string             `json:"description"`
	PublishedYear string             `json:"published_year"`
	Published     string             `json:"published"`
	Pages         string             `json:"pages"`
	ISBN          string             `json:"isbn"`
	CreatedAt     time.Time          `json:"created_at"`
	UpdatedAt     time.Time          `json:"updated_at"`
}
