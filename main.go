package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	InitDatabase()

	r.POST("/create", createMessage)
	r.GET("/get/messages/:account_id", getMessages)
	r.GET("/search", searchMessages)

	r.Run()
}

func createMessage(c *gin.Context) {
	var message Message
	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := DB.Create(&message); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, message)
}

func getMessages(c *gin.Context) {
	accountID := c.Param("account_id")
	var messages []Message
	if result := DB.Where("account_id = ?", accountID).Find(&messages); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, messages)
}

func searchMessages(c *gin.Context) {
	var messages []Message
	if messageID := c.Query("message_id"); messageID != "" {
		DB.Where("message_id IN ?", messageID).Find(&messages)
	}
	if senderNumber := c.Query("sender_number"); senderNumber != "" {
		DB.Where("sender_number IN ?", senderNumber).Find(&messages)
	}
	if receiverNumber := c.Query("receiver_number"); receiverNumber != "" {
		DB.Where("receiver_number IN ?", receiverNumber).Find(&messages)
	}

	c.JSON(http.StatusOK, messages)
}
