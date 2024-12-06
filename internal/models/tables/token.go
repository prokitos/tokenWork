package tables

type Token struct {
	TokenId int    `json:"token_id" example:"12"  gorm:"unique;primaryKey;autoIncrement"`
	GUID    string `json:"guid" example:""`
	Refresh string `json:"refresh" example:""`
}

func (instance *Token) GetId() string         { return "" }
func (instance *Token) GetQueryId() error     { return nil }
func (instance *Token) GetQueryParams() error { return nil }
func (instance *Token) GetBodyParams() error  { return nil }
