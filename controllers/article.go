package controllers

import (
	"go-mongodb-api/models"
	"go-mongodb-api/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetArticles(ctx *gin.Context) {
	res, err := services.Database().Collection("articles").Find(ctx.Request.Context(), bson.M{})

	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "INTERNAL_ERROR",
			"type":    "ERROR",
		})
	}

	var results []models.Article

	if err = res.All(ctx, &results); err != nil {
		ctx.JSON(500, gin.H{
			"message": "INTERNAL_ERROR",
			"type":    "ERROR",
		})
	}

	ctx.JSON(200, gin.H{
		"data": results,
		"type": "SUCCESS",
	})
}

func PostArticles(ctx *gin.Context) {
	if len(ctx.PostForm("title")) > 0 && len(ctx.PostForm("description")) > 0 {
		article := models.Article{
			Title:       ctx.PostForm("title"),
			Description: ctx.PostForm("description"),
		}

		_, err := services.Database().Collection("articles").InsertOne(ctx.Request.Context(), article)

		if err != nil {
			ctx.JSON(500, gin.H{
				"message": "INTERNAL_ERROR",
				"type":    "ERROR",
			})
		}

		ctx.JSON(201, gin.H{
			"code": "ARTICLE_CREATED",
			"type": "SUCCESS",
		})
		return
	}
	ctx.JSON(422, gin.H{
		"code": "MISSING_FIELDS",
		"type": "ERROR",
	})
}
