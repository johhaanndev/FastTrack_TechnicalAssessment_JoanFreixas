package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestPostAnswers_WithInvalidInput_Returns400(t *testing.T) {
	r := SetUpRouter()
	r.POST("/answers", PostAnswers)
	requestBody := []string{"1", "2", "a", "a", "a"}
	bodyJson, _ := json.Marshal(requestBody)
	request, _ := http.NewRequest("POST", "/answers", bytes.NewBuffer(bodyJson))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
