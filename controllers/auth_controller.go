package controllers

import (
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/tedjoskb/go-restapi-fiber/database"
	"github.com/tedjoskb/go-restapi-fiber/helper"
	"github.com/tedjoskb/go-restapi-fiber/models"
	"gorm.io/gorm"
)

func Login(ctx *fiber.Ctx) error {

	LoginRequest := new(models.LoginRequest)
	if err := ctx.BodyParser(&LoginRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	log.Println(LoginRequest)

	validate := validator.New()
	errValidate := validate.Struct(LoginRequest)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "login failed",
			"error":   errValidate.Error(),
		})
	}

	//check available user
	var user models.Users

	if err := database.DB.First(&user, "email =?", LoginRequest.Email).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
				"message": "User Not Found!",
				"error":   err.Error(),
			})
		}

		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Terjadi kesalahan saat mengambil data!",
			"error":   err.Error(),
		})
	}
	// check validation password

	isValid := helper.CheckPasswordHash(LoginRequest.Password, user.Password)
	if !isValid {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Wrong Credential!",
		})
	}

	return ctx.JSON(fiber.Map{
		"token": "secret",
	})

}
