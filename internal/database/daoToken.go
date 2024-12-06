package database

import (
	"mymod/internal/models"
	"mymod/internal/models/tables"

	"gorm.io/gorm"
)

var GlobalDaoToken *daoToken

type daoToken struct {
	DB *gorm.DB
}

func (currentDao *daoToken) New() {
	GlobalDaoToken = &daoToken{}
	GlobalDaoToken.DB = GlobalPostgres.Instance
}

func (currentDao *daoToken) CreateData(data tables.Token) models.IResponse {
	return nil
}

func (currentDao *daoToken) UpdateData(data tables.Token) models.IResponse {
	return nil
}

func (currentDao *daoToken) DeleteData(data tables.Token) models.IResponse {
	return nil
}

func (currentDao *daoToken) ShowData(data tables.Token) models.IResponse {
	return nil
}
