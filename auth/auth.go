package auth

import "github.com/dgrijalva/jwt-go"

type Auth interface {
	GenerateToken(userID int) (string, error)
}

type auth struct {
}

func InitAuthorization() *auth {
	return &auth{}
}

var SECRET_KEY = "BWASTARTUP_s3cr3T_k3Y"

func (a *auth) GenerateToken(userID int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}
