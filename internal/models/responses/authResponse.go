package responses

import (
	"errors"
	"mymod/internal/models/tables"

	"github.com/gofiber/fiber/v2"
)

type ResponseAuth struct {
	Code        int           `json:"code" example:"status"`
	Description string        `json:"description" example:"description"`
	Data        []tables.Auth `json:"auth,omitempty" example:"...."`
}

func (instance ResponseAuth) BaseServerError() ResponseAuth {
	instance.Code = 400
	instance.Description = "Internal Error"
	return instance
}

func (instance ResponseAuth) BadUpdate() ResponseAuth {
	instance.Code = 400
	instance.Description = "Update error"
	return instance
}

func (instance ResponseAuth) BadShow() ResponseAuth {
	instance.Code = 400
	instance.Description = "Data not exist"
	return instance
}

func (instance ResponseAuth) GoodShow(data []tables.Auth) ResponseAuth {
	instance.Code = 200
	instance.Description = "Data not exist"
	instance.Data = data
	return instance
}

func (instance ResponseAuth) GoodUpdate() ResponseAuth {
	instance.Code = 200
	instance.Description = "Data updated"
	return instance
}

func (instance ResponseAuth) CustomTokenError(text string) ResponseAuth {
	instance.Code = 400
	instance.Description = text
	return instance
}

func (instance ResponseAuth) Validate() bool {
	if instance.Code >= 200 && instance.Code < 300 {
		return true
	}
	return false
}

func (instance ResponseAuth) ToErrorBase() error {
	return errors.New(instance.Description)
}

func (instance ResponseAuth) ToErrorFiber(c *fiber.Ctx) error {
	return c.Status(instance.Code).JSON(instance)
}
