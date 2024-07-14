package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateMessage(t *testing.T) {
	router := setupRouter()

	message := Message{
		AccountID:      "test_account",
		MessageID:      "test_message_id",
		SenderNumber:   "1234567890",
		ReceiverNumber: "0987654321",
	}
	jsonValue, _ := json.Marshal(message)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/create", bytes.NewBuffer(jsonValue))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
func setupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/create", createMessage)
	r.GET("/get/messages/:account_id", getMessages)
	r.GET("/search", searchMessages)
	return r
}
