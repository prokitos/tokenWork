package database

import (
	"mymod/internal/models"
	"mymod/internal/models/responses"
	"mymod/internal/models/tables"

	log "github.com/sirupsen/logrus"
)

type daoToken struct {
}

func (currentDao *daoToken) New() {
	GlobalPostgres.DaoToken = &daoToken{}
}

func (currentlDB *daoToken) curResponse() responses.ResponseToken {
	return responses.ResponseToken{}
}

func (currentDao *daoToken) UpdateData(data tables.Token) models.IResponse {

	var finded tables.Token
	finderData := tables.Token{GUID: data.GUID}
	result := GlobalPostgres.Instance.Find(&finded, finderData)
	if result.RowsAffected == 0 {
		if result := GlobalPostgres.Instance.Model(tables.Token{}).Create(&data); result.Error != nil {
			return currentDao.curResponse().BadUpdate()
		}
		return nil
	}

	if result := GlobalPostgres.Instance.Where("guid= ?", data.GUID).Updates(&data); result.Error != nil {
		return currentDao.curResponse().BadUpdate()
	}

	log.Debug("dao complete")
	return currentDao.curResponse().GoodUpdate()
}

func (currentDao *daoToken) ExistData(data tables.Token) (tables.Token, models.IResponse) {

	var finded tables.Token

	results := GlobalPostgres.Instance.Find(&finded, data)
	if results.Error != nil || results.RowsAffected == 0 {
		log.Debug("show record error!")
		return tables.Token{}, currentDao.curResponse().BadShow()
	}

	log.Debug("check complete")
	var array []tables.Token
	array = append(array, finded)

	return finded, currentDao.curResponse().GoodShow(array)
}
