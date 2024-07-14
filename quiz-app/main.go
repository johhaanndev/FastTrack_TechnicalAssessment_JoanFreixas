package main

import (
	"quiz-app/src/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/questions", handlers.GetQuestions)
	router.POST("/answers", handlers.PostAnswers)

	router.Run("localhost:8080")
}
