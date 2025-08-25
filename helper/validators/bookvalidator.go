package validators

import (
	"errors"
	"golang-restApi/config"
	"golang-restApi/models"

	"gorm.io/gorm"
)

func CreateValidateBook(book *models.Book) error {
	if book.Title == "" {
		return errors.New("title is required")
	}

	if book.AuthorID == 0 {
		return errors.New("author_id is required")
	}

	if book.Description == "" {
		return errors.New("description is required")
	}

	if book.PublishedYear == "" {
		return errors.New("published year is required")
	}

	if book.Published == "" {
		return errors.New("published is required")
	}

	if book.ISBN == "" {
		return errors.New("isbn is required")
	}

	if book.Pages == "" {
		return errors.New("pages is required")
	}

	// check author
	var author models.Author
	if err := config.DB.First(&author, book.AuthorID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("author not found")
		}
		return err
	}

	var existing models.Book
	if err := config.DB.Where("title = ?", book.Title).First(&existing).Error; err == nil {
		return errors.New("title already exists")
	}

	return nil
}
