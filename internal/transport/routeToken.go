package transport

import (
	"github.com/gofiber/fiber/v2"
)

// выдаёт аксес и рефреш токен по guid
func getToken(c *fiber.Ctx) error {

	return c.Status(200).JSON("good")
}

// делает рефреш токена. выдаёт новый аксес токен, проверяет совпадения и айпи
func refreshToken(c *fiber.Ctx) error {

	return c.Status(200).JSON("good")
}
