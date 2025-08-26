package controllers

import (
	"encoding/json"
	"errors"
	"golang-restApi/config"
	"golang-restApi/helper"
	"golang-restApi/helper/validators"
	"golang-restApi/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// Get All
func IndexBook(w http.ResponseWriter, r *http.Request) {
	var books []models.Book
	var bookResponse []models.BookResponse

	if err := config.DB.Joins("Author").Find(&books).Find(&bookResponse).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 200, "list books", bookResponse)
}

// Create Book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		helper.Response(w, 500, err.Error(), nil)
	}

	defer r.Body.Close()

	// jalankan validasi form
	if err := validators.CreateValidateBook(&book); err != nil {
		helper.Response(w, 400, err.Error(), nil)
		return
	}

	// cek apakah author ada
	var author models.Author
	if err := config.DB.First(&author, book.AuthorID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, 404, "author not found", nil)
			return
		}

		helper.Response(w, 500, err.Error(), nil)
		return
	}

	if err := config.DB.Joins("Author").Create(&book).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 201, "success created book", nil)
}

// Detail
func DetailBook(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]

	id, _ := strconv.Atoi(idParams)

	var book models.Book
	var bookResponse models.BookResponse

	if err := config.DB.Joins("Author").First(&book, id).First(&bookResponse).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, 404, "book not found", nil)
			return
		}

		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 200, "detail book", bookResponse)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)

	var book models.Book
	if err := config.DB.First(&book, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, 404, "book not found", nil)
			return
		}
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	var input models.Book
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		helper.Response(w, 400, "invalid request body", nil)
		return
	}
	defer r.Body.Close()

	// validasi form
	if err := validators.UpdateValidateBook(book.ID, &input); err != nil {
		helper.Response(w, 400, err.Error(), nil)
		return
	}

	// cek apakah author ada
	var author models.Author
	if err := config.DB.First(&author, book.AuthorID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, 404, "author not found", nil)
			return
		}

		helper.Response(w, 500, err.Error(), nil)
		return
	}

	if err := config.DB.Model(&book).Updates(input).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 200, "success updated book", nil)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]

	id, _ := strconv.Atoi(idParams)

	var book models.Book

	res := config.DB.Delete(&book, id)

	if res.Error != nil {
		helper.Response(w, 500, res.Error.Error(), nil)
	}

	if res.RowsAffected == 0 {
		helper.Response(w, 404, "book not found", nil)
	}

	helper.Response(w, 200, "success deleted book", nil)
}
