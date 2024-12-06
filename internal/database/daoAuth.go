package database

import (
	"mymod/internal/models"
	"mymod/internal/models/tables"

	"gorm.io/gorm"
)

var GlobalDaoAuth *daoAuth

type daoAuth struct {
	DB *gorm.DB
}

func (currentDao *daoAuth) New() {
	GlobalDaoAuth = &daoAuth{}
	GlobalDaoAuth.DB = GlobalPostgres.Instance
}

func (currentDao *daoAuth) CreateData(data tables.Auth) models.IResponse {
	return nil
}

func (currentDao *daoAuth) UpdateData(data tables.Auth) models.IResponse {
	return nil
}

func (currentDao *daoAuth) DeleteData(data tables.Auth) models.IResponse {
	return nil
}

func (currentDao *daoAuth) ShowData(data tables.Auth) models.IResponse {
	return nil
}
