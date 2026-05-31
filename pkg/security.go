package pkg

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func Create_JWT_token(userID uint8) (string, error) {
	signKey := []byte("secret")

	type Claim struct {
		jwt.RegisteredClaims
		UserID int `json:"user_id"`
	}

	claims := Claim{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
		UserID: 1,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(signKey)

	return tokenString, err

}
