package models

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// модели для работы с токенами
type TokenAccessData struct {
	GUID  string
	Ip    string
	Email string
	jwt.StandardClaims
}

type TokenRefreshData struct {
	GUID         string
	AcceessToken string
	jwt.StandardClaims
}

///
///
///

type TokenData struct {
	accessToken  string
	refreshToken string
	payload      TokenPayloadData
	err          error
}

type TokenPayloadData struct {
	GUID  string
	Ip    string
	Email string
	Time  time.Time
}

type TokenPair struct {
	AccessToken  string
	RefreshToken string
}
