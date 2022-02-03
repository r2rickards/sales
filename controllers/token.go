package controllers

import (
	"fmt"
	"go-mongodb-api/services"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type TokenData struct {
	Username string `json:"username"`
}

type ResponseData struct {
	Token        string                 `json:"token"`
	TokenDecoded map[string]interface{} `json:"tokenDecoded"`
}

func Token(ctx *gin.Context) {
	var token string = services.JWTAuthService().Encode(TokenData{Username: "test"})

	tokenDecoded, err := services.JWTAuthService().Decode(token)

	if tokenDecoded.Valid {
		claims := tokenDecoded.Claims.(jwt.MapClaims)

		ctx.JSON(200, gin.H{
			"data": ResponseData{Token: token, TokenDecoded: claims},
			"type": "SUCCESS",
		})
		fmt.Println(claims)
	} else {
		fmt.Println(err)
	}

}
