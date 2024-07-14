package validations

import (
	"fmt"
	"quiz-app/src/models"
	"regexp"

	"github.com/gin-gonic/gin"
)

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
