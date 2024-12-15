package database

import "mymod/internal/models/responses"

type daoAuth struct {
}

func (currentDao *daoAuth) New() {
	GlobalPostgres.DaoAuth = &daoAuth{}
}

func (currentlDB *daoAuth) curResponse() responses.ResponseToken {
	return responses.ResponseToken{}
}
