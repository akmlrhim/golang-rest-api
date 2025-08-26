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
func IndexAuthor(w http.ResponseWriter, r *http.Request) {
	var author []models.Author

	if err := config.DB.Find(&author).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 200, "list author", author)
}

// Create
func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var author models.Author

	if err := json.NewDecoder(r.Body).Decode(&author); err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	// jalankan validasi
	if err := validators.CreateValidateAuthor(&author); err != nil {
		helper.Response(w, 400, err.Error(), nil)
		return
	}

	if err := config.DB.Create(&author).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 201, "success created author", nil)
}

// Detail
func DetailAuthor(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]

	id, _ := strconv.Atoi(idParams)

	var author models.Author

	if err := config.DB.First(&author, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, 404, "author not found", nil)
			return
		}

		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 200, "detail author", author)
}

// Update
func UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)

	var author models.Author
	if err := config.DB.First(&author, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, 404, "author not found", nil)
			return
		}
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	var input models.Author
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		helper.Response(w, 400, "invalid request body", nil)
		return
	}
	defer r.Body.Close()

	// validasi
	if err := validators.UpdateValidateAuthor(author.ID, &input); err != nil {
		helper.Response(w, 400, err.Error(), nil)
		return
	}

	// update hanya field yang ada di input (dan tidak kosong)
	if err := config.DB.Model(&author).Updates(input).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 200, "success updated author", nil)
}

// Delete
func DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]

	id, _ := strconv.Atoi(idParams)

	var author models.Author

	res := config.DB.Delete(&author, id)

	if res.Error != nil {
		helper.Response(w, 500, res.Error.Error(), nil)
	}

	if res.RowsAffected == 0 {
		helper.Response(w, 404, "author not found", nil)
	}

	helper.Response(w, 200, "success deleted author", nil)
}
