package handlers

import (
	"fmt"
	"net/http"
	"quiz-app/src/csvServices"
	"quiz-app/src/models"
	"quiz-app/src/validations"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetQuestions(c *gin.Context) {
	responseQuestions := make([]models.GetQuestionsReponse, len(models.QuizQuestions))
	for i, q := range models.QuizQuestions {
		responseQuestions[i] = q.ToGetQuestionsResponse()
	}
	c.IndentedJSON(http.StatusOK, responseQuestions)
}

func PostAnswers(c *gin.Context) {
	var receivedAnswers []string

	if err := c.BindJSON(&receivedAnswers); err != nil {
		return
	}

	if err := validations.ValidateRequest(c, receivedAnswers); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	score := 0

	for i, answer := range receivedAnswers {
		if answer == models.QuizQuestions[i].Correct {
			score++
		}
	}
	totalQuestions := len(models.QuizQuestions)
	scorePercentage := float64(score) / float64(totalQuestions) * 100
	scoreToString := fmt.Sprintf("%.2f", scorePercentage)

	var playerInfo = models.PlayerInfo{
		ID:    uuid.New().String(),
		Score: scoreToString,
	}

	models.Players = append(models.Players, playerInfo)
	allPlayers, err := csvServices.ReadScoresCsv()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading players from CSV"})
		return
	}

	topPercent := csvServices.CalculateTopScorePercentage(scorePercentage, allPlayers)
	if scorePercentage == 0.0 {
		topPercent = 0.0
	}
	topPercentToString := fmt.Sprintf("%.2f", topPercent)

	if err := csvServices.UpdateScoresCsv(); err != nil {
		errorMessage := fmt.Sprintf("Error: '%s'", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": errorMessage})
		return
	}

	resultMessage := fmt.Sprintf("Correct answers: %s%%. You were better than %s%% of all quizzers", scoreToString, topPercentToString)
	c.JSON(http.StatusOK, resultMessage)
}
