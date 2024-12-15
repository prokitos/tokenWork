package responses

import (
	"errors"
	"mymod/internal/models/tables"

	"github.com/gofiber/fiber/v2"
)

type ResponseToken struct {
	Code        int            `json:"code" example:"status"`
	Description string         `json:"description" example:"description"`
	Data        []tables.Token `json:"auth,omitempty" example:"...."`
}

func (instance ResponseToken) BaseServerError() ResponseToken {
	instance.Code = 400
	instance.Description = "Internal Error"
	return instance
}

func (instance ResponseToken) BadUpdate() ResponseToken {
	instance.Code = 400
	instance.Description = "Update error"
	return instance
}

func (instance ResponseToken) BadShow() ResponseToken {
	instance.Code = 400
	instance.Description = "Data not exist"
	return instance
}

func (instance ResponseToken) GoodShow(data []tables.Token) ResponseToken {
	instance.Code = 200
	instance.Description = "Data not exist"
	instance.Data = data
	return instance
}

func (instance ResponseToken) GoodUpdate() ResponseToken {
	instance.Code = 200
	instance.Description = "Data updated"
	return instance
}

func (instance ResponseToken) CustomTokenError(text string) ResponseToken {
	instance.Code = 400
	instance.Description = text
	return instance
}

func (instance ResponseToken) Validate() bool {
	if instance.Code >= 200 && instance.Code < 300 {
		return true
	}
	return false
}

func (instance ResponseToken) ToErrorBase() error {
	return errors.New(instance.Description)
}

func (instance ResponseToken) ToErrorFiber(c *fiber.Ctx) error {
	return c.Status(instance.Code).JSON(instance)
}
