package models

import (
	"mymod/internal/config"

	"github.com/gofiber/fiber/v2"
)

type IResponse interface {
	GetError(c *fiber.Ctx) error
	Validate() bool
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
