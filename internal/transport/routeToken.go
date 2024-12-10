package transport

import (
	"mymod/internal/services"

	"github.com/gofiber/fiber/v2"
)

// выдаёт аксес и рефреш токен по guid
func getToken(c *fiber.Ctx) error {

	var GUID = c.Query("GUID", "")

	var temp services.TokenData
	temp.AddTimestamp()
	temp.AddGuid(GUID)
	acc, ref, err := temp.CreatePair()
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	var res services.TokenPair
	res.AccessToken = acc
	res.RefreshToken = ref

	return c.Status(200).JSON(res)
}

// делает рефреш токена. выдаёт новый аксес токен, проверяет совпадения и айпи
func refreshToken(c *fiber.Ctx) error {

	var access = c.Query("access", "")
	var refresh = c.Query("refresh", "")

	var temp services.TokenData
	acc, ref, err := temp.RefreshToken(access, refresh)
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	var res services.TokenPair
	res.AccessToken = acc
	res.RefreshToken = ref

	return c.Status(200).JSON(res)
}
