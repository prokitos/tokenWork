package models

import (
	"mymod/internal/config"

	"github.com/gofiber/fiber/v2"
)

type IResponse interface {
	Validate() bool
	ToErrorBase() error
	ToErrorFiber(c *fiber.Ctx) error
}

type IDatabase interface {
	OpenConnection(config.MainConfig)
	StartMigration()
	GlobalSet()
}

type ITable interface {
	GetId() string
	GetQueryId() error
	GetQueryParams() error
	GetBodyParams() error
}
