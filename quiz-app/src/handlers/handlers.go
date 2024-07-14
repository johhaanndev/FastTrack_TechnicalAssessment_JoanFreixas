package handlers

import (
	"fmt"
	"net/http"
	"quiz-app/src/models"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetQuestions(c *gin.Context) {
	responseQuestions := make([]models.GetQuestionsReponse, len(models.QuizQuestions))
	for i, q := range models.QuizQuestions {
		responseQuestions[i] = q.ToResponseQuestion()
	}
	c.IndentedJSON(http.StatusOK, responseQuestions)
}

func PostAnswers(c *gin.Context) {
	var receivedAnswers []string

	if err := c.BindJSON(&receivedAnswers); err != nil {
		return
	}

	if err := ValidateRequest(c, receivedAnswers); err != nil {
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

	if err := models.UpdateScoresCsv(); err != nil {
		errorMessage := fmt.Sprintf("Error: '%s'", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": errorMessage})
		return
	}

	allPlayers, err := models.ReadScoresCsv()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading players from CSV"})
		return
	}
	position := models.CalculateTopScorePercentage(scorePercentage, allPlayers)

	positionToString := fmt.Sprintf("%.2f", position)

	resultMessage := fmt.Sprintf("Correct answers: %s%%. You were better than %s%% of all quizzers", scoreToString, positionToString)
	c.JSON(http.StatusOK, resultMessage)
}

func ValidateRequest(c *gin.Context, answers []string) error {
	if len(answers) != len(models.QuizQuestions) {
		return fmt.Errorf("The answer array must have the same length as the questions")
	}

	regex := regexp.MustCompile("[^a-zA-Z]")
	for _, answer := range answers {
		if match := regex.FindString(answer); match != "" {
			return fmt.Errorf("Answer '%s' contains non-alphabetical characters", answer)
		}
	}

	return nil
}
