package validators

import (
	"errors"
	"golang-restApi/config"
	"golang-restApi/models"
)

func CreateValidateAuthor(author *models.Author) error {
	if author.Name == "" {
		return errors.New("name is required")
	}

	if author.Email == "" {
		return errors.New("email is required")
	}

	if author.Age == 0 {
		return errors.New("age is required")
	}

	if author.Gender == "" {
		return errors.New("gender is required")
	}

	var existingName models.Author
	if err := config.DB.Where("name = ?", author.Name).First(&existingName).Error; err == nil {
		return errors.New("name already exists")
	}

	var existingEmail models.Author
	if err := config.DB.Where("email = ?", author.Email).First(&existingEmail).Error; err == nil {
		return errors.New("email already exists")
	}

	return nil
}

func UpdateValidateAuthor(id uint, author *models.Author) error {
	if author.Name == "" {
		return errors.New("name is required")
	}

	if author.Email == "" {
		return errors.New("email is required")
	}

	if author.Age == 0 {
		return errors.New("age is required")
	}

	if author.Gender == "" {
		return errors.New("gender is required")
	}

	var existingName models.Author
	if err := config.DB.Where("name = ? AND id != ?", author.Name, id).First(&existingName).Error; err == nil {
		return errors.New("name already exists")
	}

	var existingEmail models.Author
	if err := config.DB.Where("email = ? AND id != ?", author.Email, id).First(&existingEmail).Error; err == nil {
		return errors.New("email already exists")
	}

	return nil
}
