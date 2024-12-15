package transport

import (
	"mymod/internal/database"
	"mymod/internal/models/tables"

	"github.com/gofiber/fiber/v2"
)

func SetHandlers(instance *fiber.App) {

	instance.Get("/getToken", getToken)
	instance.Get("/refreshToken", refreshToken)

	instance.Get("/Token", testResponse)

}

func testResponse(c *fiber.Ctx) error {

	var data tables.Token
	data.GUID = c.Query("GUID", "")
	_, resp := database.GlobalPostgres.DaoToken.ExistData(data)

	return resp.ToErrorFiber(c)
}
