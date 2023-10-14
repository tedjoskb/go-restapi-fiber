package controllers

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/tedjoskb/go-restapi-fiber/database"
	"github.com/tedjoskb/go-restapi-fiber/models"
	"gorm.io/gorm"
)

func GetUserAll(c *fiber.Ctx) error {
	var users []models.Users
	database.DB.Table("users").Select("id, name, email, address, created_at, updated_at, is_deleted, deleted_at").
		Where("is_deleted = ?", false).
		Find(&users)

	result := c.Status(http.StatusOK).JSON(fiber.Map{
		"data":    users,
		"message": "Success",
	})

	return result
}

func GetUserById(c *fiber.Ctx) error {
	id := c.Params("id")
	var users models.Users
	if err := database.DB.First(&users, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"message": "data tidak ditemukan!",
			})
		}

		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "data tidak ditemukan!",
		})
	}

	result := c.Status(http.StatusOK).JSON(fiber.Map{
		"data":    users,
		"message": "Success",
	})

	return result
}

func CreateUser(c *fiber.Ctx) error {
	var userCreates []models.UserCreate
	if err := c.BodyParser(&userCreates); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	validate := validator.New()

	// Loop melalui array pengguna yang akan dibuat
	var newUsers []models.Users
	var validationErrors []error
	for _, userCreate := range userCreates {
		errValidate := validate.Struct(userCreate)
		if errValidate != nil {
			validationErrors = append(validationErrors, errValidate)
			continue
		}

		newUser := models.Users{
			Name:      userCreate.Name,
			Email:     userCreate.Email,
			Address:   userCreate.Address,
			IsDeleted: false,
		}
		newUsers = append(newUsers, newUser)
	}

	if len(validationErrors) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Some users failed validation",
			"errors":  validationErrors,
		})
	}

	// Simpan pengguna yang valid ke database
	if err := database.DB.Create(&newUsers).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	result := c.Status(http.StatusOK).JSON(fiber.Map{
		"data":    newUsers,
		"message": "Users created successfully",
	})

	return result

}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var UpdateUser models.UserUpdate
	if err := c.BodyParser(&UpdateUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	Update := models.Users{
		Name:    UpdateUser.Name,
		Address: UpdateUser.Address,
		Email:   UpdateUser.Email,
	}

	if database.DB.Where("id = ?", id).Updates(&Update).RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "tidak dapat mengupdate data",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data Berhasil Diupdate",
	})
}

func UpdateMultipleUsers(c *fiber.Ctx) error {
	var updateRequests []models.Users // Ubah dengan struct Anda sendiri
	if err := c.BodyParser(&updateRequests); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	for _, updateRequest := range updateRequests {
		update := models.Users{
			Name:    updateRequest.Name,
			Address: updateRequest.Address,
		}

		if database.DB.Where("id = ?", updateRequest.ID).Updates(&update).RowsAffected == 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Tidak dapat mengupdate data dengan ID " + string(updateRequest.ID),
			})
		}
	}

	return c.JSON(fiber.Map{
		"message": "Data Berhasil Diupdate",
	})
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var book models.Book
	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if database.DB.Where("id = ?", id).Delete(&book).RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "tidak dapat menghapus data",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data Berhasil Menghapus",
	})
}
