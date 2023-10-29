package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/tedjoskb/go-restapi-fiber/helper"
)

func Auth(ctx *fiber.Ctx) error {
	token := ctx.Get("x-token")

	if token == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthenticated",
		})
	}

	auth := ctx.Get("Authorization")

	if auth == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Authorization header is missing",
		})
	}

	// Hapus "Bearer " dari token string
	tokenString := strings.TrimPrefix(auth, "Bearer ")

	// Mendekode token JWT dan memeriksa validitasnya
	validToken, err := helper.VerifyToken(tokenString)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid or expired token",
			"error":   err.Error(), // Menampilkan pesan kesalahan yang lebih spesifik
		})
	}

	claims, err := helper.DecodeToken(tokenString)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err,
		})
	}

	role, ok := claims["role"].(string)
	if !ok || role != "admin" {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "Forbidden Access",
		})
	}

	ctx.Locals("userInfo", claims)

	// Pastikan token valid sebelum lanjut ke handler berikutnya
	if validToken.Valid {
		return ctx.Next()
	} else {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid token",
		})
	}

}

func PermissionCreate(ctx *fiber.Ctx) error {
	return ctx.Next()
}
