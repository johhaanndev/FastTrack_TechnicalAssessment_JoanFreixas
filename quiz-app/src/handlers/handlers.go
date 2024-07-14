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
	responseQuestions := make([]models.GetQuestionsReponse, len(models.SampleQuestions))
	for i, q := range models.SampleQuestions {
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
		if answer == models.SampleQuestions[i].Correct {
			score++
		}
	}
	totalQuestions := len(models.SampleQuestions)
	percentage := float64(score) / float64(totalQuestions) * 100
	formattedPercentage := fmt.Sprintf("%.2f", percentage)

	var playerInfo = models.PlayerInfo{
		ID:    uuid.New().String(),
		Score: formattedPercentage,
	}

	models.Players = append(models.Players, playerInfo)

	if err := models.UpdatePlayersCsv(); err != nil {
		errorMessage := fmt.Sprintf("Error: '%s'", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": errorMessage})
		return
	}

	c.JSON(http.StatusOK, fmt.Sprintf("Correct answers: %s%%", formattedPercentage))
}

func ValidateRequest(c *gin.Context, answers []string) error {
	if len(answers) != len(models.SampleQuestions) {
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
