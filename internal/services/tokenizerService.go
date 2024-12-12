package services

import (
	"errors"
	"mymod/internal/database"
	"mymod/internal/models/tables"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var accessKey = []byte("basic_key")
var refreshKey = []byte("super_secret_key")

func (instance *TokenData) AddTimestamp() {
	instance.payload.Time = time.Now()
}
func (instance *TokenData) AddGuid(guid string) {
	instance.payload.GUID = guid
}
func (instance *TokenData) AddIp(ip string) {
	instance.payload.Ip = ip
}
func (instance *TokenData) AddEmail(email string) {
	instance.payload.Email = email
}

func (instance *TokenData) CreatePair() (string, string, error) {

	accessToken, err := instance.createAccessToken()
	if err != nil {
		return "", "", err
	}

	refreshToken, err := instance.createRefreshToken()
	if err != nil {
		return "", "", err
	}

	refreshTokenSecured, err := instance.bcryptToken(refreshToken)
	if err != nil {
		return "", "", err
	}

	var data tables.Token
	data.GUID = instance.payload.GUID
	data.Refresh = refreshTokenSecured
	data.Stamp = instance.payload.Time
	err = database.GlobalPostgres.DaoToken.UpdateData(data)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (instance *TokenData) RefreshToken(access string, refresh string) (string, string, error) {

	accessClaim, err := instance.verifyAccessToken(access)
	if err != nil {
		return "", "", err
	}
	instance.AddGuid(accessClaim.Payload.GUID)

	refreshClaim, err := instance.verifyRefreshToken(refresh)
	if err != nil {
		return "", "", err
	}

	var data tables.Token
	data.GUID = instance.payload.GUID
	data.Stamp = accessClaim.Payload.Time
	res, err := database.GlobalPostgres.DaoToken.ExistData(data)
	if err != nil {
		return "", "", err
	}

	var databaseBcryptToken string = res.Refresh
	err = instance.bcryptCheck(refresh, databaseBcryptToken)
	if err != nil {
		return "", "", err
	}

	if accessClaim.Payload != refreshClaim.Payload {
		// делаем проверку на айпи и на таймстемп. но сейчас лень
		return "", "", errors.New("Access and refresh tokens do not match")
	}

	newAccess, newRefresh, err := instance.CreatePair()
	if err != nil {
		return "", "", err
	}

	return newAccess, newRefresh, nil
}

func (instance *TokenData) bcryptToken(refresh string) (string, error) {

	if len(refresh) > 72 {
		refresh = refresh[:72]
	}

	hashedToken, err := bcrypt.GenerateFromPassword([]byte(refresh), bcrypt.MinCost)

	if err != nil {
		return "", err
	}

	return string(hashedToken), nil
}
func (instance *TokenData) bcryptCheck(token string, cryptToken string) error {

	if len(token) > 72 {
		token = token[:72]
	}

	err := bcrypt.CompareHashAndPassword([]byte(cryptToken), []byte(token))
	if err != nil {
		return errors.New("bcrypt check failed")
	} else {
		return nil
	}
}
func (instance *TokenData) verifyAccessToken(access string) (*TokenAccessData, error) {

	claims := &TokenAccessData{}
	token, err := jwt.ParseWithClaims(access, claims, func(token *jwt.Token) (interface{}, error) {
		return accessKey, nil
	})

	if err != nil || !token.Valid {
		return claims, err
	}
	return claims, nil
}
func (instance *TokenData) verifyRefreshToken(refresh string) (*TokenAccessData, error) {

	claims := &TokenAccessData{}
	token, err := jwt.ParseWithClaims(refresh, claims, func(token *jwt.Token) (interface{}, error) {
		return refreshKey, nil
	})

	if err != nil || !token.Valid {
		return claims, err
	}
	return claims, nil
}

func (instance *TokenData) createAccessToken() (string, error) {
	claims := &TokenAccessData{
		Payload: instance.payload,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(3 * time.Minute).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return token.SignedString(accessKey)
}

func (instance *TokenData) createRefreshToken() (string, error) {
	claims := &TokenRefreshData{
		Payload: instance.payload,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(10 * time.Minute).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return token.SignedString(refreshKey)
}

///
///
///

type TokenPayloadData struct {
	GUID  string
	Ip    string
	Email string
	Time  time.Time
}

type TokenData struct {
	payload TokenPayloadData
}

type TokenAccessData struct {
	Payload TokenPayloadData
	jwt.StandardClaims
}

type TokenRefreshData struct {
	Payload TokenPayloadData
	jwt.StandardClaims
}
