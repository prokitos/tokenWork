package models

import (
	"mymod/internal/config"
)

type IResponse interface {
	GetError() error
	Validate() bool
	BadUpdate() error
	BadShow() error
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
