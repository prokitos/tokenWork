package tables

type Auth struct {
	UserId int    `json:"user_id" example:"12"  gorm:"unique;primaryKey;autoIncrement"`
	GUID   string `json:"guid" example:""`
	Email  string `json:"email" example:""`
}

func (instance *Auth) GetId() string         { return "" }
func (instance *Auth) GetQueryId() error     { return nil }
func (instance *Auth) GetQueryParams() error { return nil }
func (instance *Auth) GetBodyParams() error  { return nil }
