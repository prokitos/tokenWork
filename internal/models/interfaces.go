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
type IDatabaseDao interface {
	CreateData(ITable) IResponse
	DeleteData(ITable) IResponse
	UpdateData(ITable) IResponse
	ShowData(ITable) IResponse
}

type ITable interface {
	GetId() string
	GetQueryId() error
	GetQueryParams() error
	GetBodyParams() error
}

type IToken interface {
	AddAccess()
	AddRefresh()
	addPayload()
	GetError()
	CreatePair()
	VerifyToken()
	RefreshToken()
}
