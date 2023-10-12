package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/tedjoskb/go-restapi-fiber/database"
	"github.com/tedjoskb/go-restapi-fiber/models"
	"gorm.io/gorm"
)

func Index(c *fiber.Ctx) error {
	var books []models.Book
	database.DB.Find(&books)

	return c.Status(fiber.StatusOK).JSON(books)
}

func Show(c *fiber.Ctx) error {
	id := c.Params("id")
	var books models.Book
	if err := database.DB.First(&books, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"message": "data tidak ditemukan!",
			})
		}

		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "data tidak ditemukan!",
		})
	}

	return c.Status(fiber.StatusOK).JSON(books)
}

func Create(c *fiber.Ctx) error {
	var book models.Book
	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := database.DB.Create(&book).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(book)

}

func Update(c *fiber.Ctx) error {
	id := c.Params("id")
	var book models.Book
	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if database.DB.Where("id = ?", id).Updates(&book).RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "tidak dapat mengupdate data",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data Berhasil Diupdate",
	})
}

func Delete(c *fiber.Ctx) error {
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
