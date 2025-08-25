package authcontroller

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
func Index(w http.ResponseWriter, r *http.Request) {
	var author []models.Author

	if err := config.DB.Find(&author).Error; err != nil {
		helper.Response(w, 200, err.Error(), nil)

	}

	helper.Response(w, 200, "List Authors", author)
}

// Create
func Create(w http.ResponseWriter, r *http.Request) {
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

	helper.Response(w, 200, "success created author", author)
}

// Detail
func Detail(w http.ResponseWriter, r *http.Request) {
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

func Update(w http.ResponseWriter, r *http.Request) {
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

}
