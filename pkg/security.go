package pkg

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const signKey = "secret"

type Claim struct {
	jwt.RegisteredClaims
	UserID uint8 `json:"user_id"`
}

func Create_JWT_token(userID uint8) (string, error) {

	claims := Claim{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
		UserID: 1,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(signKey))

	return tokenString, err

}

func ParseJWT(tokenString string) (Claim, error) {
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.ParseWithClaims(tokenString, &Claim{}, func(token *jwt.Token) (any, error) {
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(signKey), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		log.Fatal(err)
	}
	claims, ok := token.Claims.(*Claim)
	if !ok {
		return Claim{}, fmt.Errorf("token is not valid")
	} else {
		return *claims, nil
	}
}
