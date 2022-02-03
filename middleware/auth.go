package middleware

import (
	"go-mongodb-api/services"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type authHeader struct {
	IDToken string `header:"Authorization"`
}

type invalidArgument struct {
	Field string `json:"field"`
	Value string `json:"value"`
	Tag   string `json:"tag"`
	Param string `json:"param"`
}

func AuthUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		h := authHeader{}

		if err := c.ShouldBindHeader(&h); err != nil {
			if errs, ok := err.(validator.ValidationErrors); ok {
				var invalidArgs []invalidArgument

				for _, err := range errs {
					invalidArgs = append(invalidArgs, invalidArgument{
						err.Field(),
						err.Value().(string),
						err.Tag(),
						err.Param(),
					})
				}

				c.JSON(500, gin.H{
					"code":        "INVALID_REQUEST_PARAMETERS",
					"type":        "ERROR",
					"invalidArgs": invalidArgs,
				})
				c.Abort()
				return

			}

			c.JSON(500, gin.H{
				"code": "UNKNOWN_ERROR",
				"type": "ERROR",
			})
			c.Abort()
			return

		}

		idTokenHeader := strings.Split(h.IDToken, "Bearer ")

		if len(idTokenHeader) < 2 {
			c.JSON(500, gin.H{
				"code": "TOKEN_NOT_FOUND",
				"type": "ERROR",
			})
			c.Abort()
			return
		}

		_, err := services.JWTAuthService().Decode(idTokenHeader[1])
		if err != nil {
			c.JSON(500, gin.H{
				"code": "TOKEN_INVALID",
				"type": "ERROR",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
