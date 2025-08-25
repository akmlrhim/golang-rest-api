package controllers

import (
	"encoding/json"
	"golang-restApi/config"
	"golang-restApi/helper"
	"golang-restApi/helper/validators"
	"golang-restApi/models"
	"net/http"
)

// Get All
func IndexBook(w http.ResponseWriter, r *http.Request) {
	var books []models.Book

	if err := config.DB.Joins("Author").Find(&books).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 200, "list books", books)
}

// Create Book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		helper.Response(w, 500, err.Error(), nil)
	}

	defer r.Body.Close()

	// jalankan validasi
	if err := validators.CreateValidateBook(&book); err != nil {
		helper.Response(w, 400, err.Error(), nil)
		return
	}

	if err := config.DB.Joins("Author").Create(&book).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 201, "success created book", book)
}

// Detail
func DetailBook(w http.ResponseWriter, r *http.Request) {

}
