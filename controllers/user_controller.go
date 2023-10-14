package controllers

import (
	"net/http"
	"time"

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

	// Mengkonversi Unix timestamp ke dalam format tanggal yang dapat dibaca
	for i, user := range users {
		users[i].CreatedAtFormatted = time.Unix(0, user.CreatedAt*int64(time.Millisecond)).Format("2023-01-01 15:04:05")
		users[i].UpdatedAtFormatted = time.Unix(0, user.UpdatedAt*int64(time.Millisecond)).Format("2023-01-01 15:04:05")
	}

	result := c.Status(http.StatusOK).JSON(fiber.Map{
		"data":    users,
		"message": "Success",
	})

	return result
}

func GetUserById(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.Users

	if err := database.DB.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"message": "Data tidak ditemukan!",
			})
		}

		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Terjadi kesalahan saat mengambil data!",
		})
	}

	// Mengonversi waktu CreatedAt dan UpdatedAt ke dalam format yang diinginkan
	user.CreatedAtFormatted = time.Unix(0, user.CreatedAt*int64(time.Millisecond)).Format("2006-01-02 15:04:05")
	user.UpdatedAtFormatted = time.Unix(0, user.UpdatedAt*int64(time.Millisecond)).Format("2006-01-02 15:04:05")

	result := c.Status(http.StatusOK).JSON(fiber.Map{
		"data":    user,
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
			CreatedAt: time.Now().UnixMilli(), // Mengatur created_at ke Unix timestamp dalam milidetik
			UpdatedAt: time.Now().UnixMilli(),
			DeletedAt: 0,
			IsDeleted: false,
		}
		// show data response convertan epoch time (unix) tanpa dimasukin ke database
		newUser.CreatedAtFormatted = time.Unix(0, newUser.CreatedAt*int64(time.Millisecond)).Format("2006-01-02 15:04:05")
		newUser.UpdatedAtFormatted = time.Unix(0, newUser.UpdatedAt*int64(time.Millisecond)).Format("2006-01-02 15:04:05")
		newUser.DeletedAtFormatted = ""

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
		Name:      UpdateUser.Name,
		Address:   UpdateUser.Address,
		Email:     UpdateUser.Email,
		UpdatedAt: time.Now().UnixMilli(),
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
	var updateRequests []models.Users // Sesuaikan dengan struktur Anda sendiri
	if err := c.BodyParser(&updateRequests); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// loc, err := time.LoadLocation("Asia/Jakarta")
	// if err != nil {
	// 	panic(err)
	// }

	for _, updateRequest := range updateRequests {

		update := models.Users{
			ID:        updateRequest.ID,
			Name:      updateRequest.Name,
			Address:   updateRequest.Address,
			Email:     updateRequest.Email,
			UpdatedAt: time.Now().UnixMilli(), // Menggunakan waktu yang sesuai dengan zona masing-masing (in unix time)
			DeletedAt: time.Now().UnixMilli(),
			IsDeleted: updateRequest.IsDeleted,
		}

		if updateRequest.IsDeleted {
			// Gunakan Unscoped untuk mengatur DeletedAt saat IsDeleted true
			if err := database.DB.Model(&update).
				Where("id = ?", updateRequest.ID).
				Update("Name", update.Name).
				Update("Address", update.Address).
				Update("Email", update.Email).
				Update("UpdatedAt", update.UpdatedAt).
				Update("DeletedAt", update.DeletedAt).
				Update("IsDeleted", true).Error; err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"message": "Tidak dapat mengupdate data dengan ID " + string(updateRequest.ID),
				})
			}
		} else {
			// Jika IsDeleted false, hanya mengupdate IsDeleted tanpa menyentuh DeletedAt
			if err := database.DB.Model(&update).
				Where("id = ?", updateRequest.ID).
				Update("Name", update.Name).
				Update("Address", update.Address).
				Update("Email", update.Email).
				Update("UpdatedAt", update.UpdatedAt).
				Update("DeletedAt", update.DeletedAt).
				Update("IsDeleted", false).Error; err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"message": "Tidak dapat mengupdate data dengan ID " + string(updateRequest.ID),
				})
			}
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
