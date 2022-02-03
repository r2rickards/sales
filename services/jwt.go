package services

import (
	"fmt"
	"go-mongodb-api/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTService interface {
	Encode(data interface{}) string
	Decode(token string) (*jwt.Token, error)
}

type jwtServices struct {
	secretKey string
	issure    string
}

func JWTAuthService() JWTService {
	return &jwtServices{
		secretKey: utils.Config().Jwt.Secret,
		issure:    "Bikash",
	}
}

func (service *jwtServices) Encode(data interface{}) string {
	atClaims := jwt.MapClaims{}
	atClaims["data"] = data
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	t, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (service *jwtServices) Decode(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("invalid token %d", token.Header["alg"])
		}
		return []byte(service.secretKey), nil
	})
}
