package controllers

import (
	"encoding/json"
	"errors"
	"golang-restApi/config"
	"golang-restApi/helper"
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
		helper.Response(w, 200, err.Error(), nil)

	}

	helper.Response(w, 200, "List Authors", author)
}

// Create
func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var author models.Author
	var existingEmail models.Author
	var existingName models.Author

	if err := json.NewDecoder(r.Body).Decode(&author); err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	if author.Name == "" {
		helper.Response(w, 400, "name is required", nil)
		return
	}

	if author.Email == "" {
		helper.Response(w, 400, "email is required", nil)
		return
	}

	if author.Age == 0 {
		helper.Response(w, 400, "age is required and more than 0", nil)
		return
	}

	if author.Gender == "" {
		helper.Response(w, 400, "gender is required", nil)
	}

	if err := config.DB.Where("name = ?", author.Name).First(&existingName).Error; err == nil {
		helper.Response(w, 400, "Name already exists", nil)
		return
	}

	if err := config.DB.Where("email = ?", author.Email).First(&existingEmail).Error; err == nil {
		helper.Response(w, 400, "Email already exists", nil)
		return
	}

	if err := config.DB.Create(&author).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 201, "success created author", author)
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

	if input.Name == "" {
		helper.Response(w, 400, "name is required", nil)
		return
	}
	if input.Email == "" {
		helper.Response(w, 400, "email is required", nil)
		return
	}
	if input.Age <= 0 {
		helper.Response(w, 400, "age is required and must be greater than 0", nil)
		return
	}
	if input.Gender == "" {
		helper.Response(w, 400, "gender is required", nil)
		return
	}

	var existingName models.Author
	if err := config.DB.Where("name = ? AND id != ?", input.Name, id).First(&existingName).Error; err == nil {
		helper.Response(w, 400, "name already exists", nil)
		return
	}

	var existingEmail models.Author
	if err := config.DB.Where("email = ? AND id != ?", input.Email, id).First(&existingEmail).Error; err == nil {
		helper.Response(w, 400, "email already exists", nil)
		return
	}

	author.Name = input.Name
	author.Email = input.Email
	author.Age = input.Age
	author.Gender = input.Gender

	if err := config.DB.Save(&author).Error; err != nil {
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
