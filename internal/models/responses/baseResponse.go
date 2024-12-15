package responses

import (
	"errors"
	"mymod/internal/models"

	"github.com/gofiber/fiber/v2"
)

type ResponseBase struct {
	Code        int             `json:"code" example:"status"`
	Description string          `json:"description" example:"description"`
	Data        []models.ITable `json:"data,omitempty" example:"...."`
}

func (instance ResponseBase) BaseServerError() ResponseBase {
	instance.Code = 400
	instance.Description = "Internal Error"
	return instance
}

func (instance ResponseBase) BadUpdate() ResponseBase {
	instance.Code = 400
	instance.Description = "Update error"
	return instance
}

func (instance ResponseBase) BadShow() ResponseBase {
	instance.Code = 400
	instance.Description = "Data not exist"
	return instance
}

func (instance ResponseBase) GoodShow(data []models.ITable) ResponseBase {
	instance.Code = 200
	instance.Description = "Data not exist"
	instance.Data = data
	return instance
}

func (instance ResponseBase) GoodUpdate() ResponseBase {
	instance.Code = 200
	instance.Description = "Data updated"
	return instance
}

func (instance ResponseBase) CustomTokenError(text string) ResponseBase {
	instance.Code = 400
	instance.Description = text
	return instance
}

func (instance ResponseBase) Validate() bool {
	if instance.Code >= 200 && instance.Code < 300 {
		return true
	}
	return false
}

func (instance ResponseBase) ToErrorBase() error {
	return errors.New(instance.Description)
}

func (instance ResponseBase) ToErrorFiber(c *fiber.Ctx) error {
	return c.Status(instance.Code).JSON(instance)
}
