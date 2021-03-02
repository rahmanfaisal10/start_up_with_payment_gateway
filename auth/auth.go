package auth

import (
	"bwastartup/pkg/model"
	"errors"

	"github.com/dgrijalva/jwt-go"
)

type Auth interface {
	GenerateToken(user model.User) (string, error)
	ValidationToken(token string) (*jwt.Token, error)
}

type auth struct {
}

func InitAuthorization() *auth {
	return &auth{}
}

var SECRET_KEY = "BWASTARTUP_s3cr3T_k3Y"

func (a *auth) GenerateToken(user model.User) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = user.UUID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (a *auth) ValidationToken(encodeToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodeToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid Token")
		}

		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return token, err
	}
	return token, nil
}
