package database

import (
	"mymod/internal/models"
	"mymod/internal/models/tables"
)

type daoAuth struct {
}

func (currentDao *daoAuth) New() {
	GlobalPostgres.DaoAuth = &daoAuth{}
}

func (currentDao *daoAuth) CreateData(data tables.Auth) models.IResponse {
	return nil
}
