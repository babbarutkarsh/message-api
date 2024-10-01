package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Message represents the message structure in the database
type Message struct {
	ID             uint   `gorm:"primaryKey"`
	AccountID      string `json:"account_id"`
	MessageID      string `json:"message_id"`
	SenderNumber   string `json:"sender_number"`
	ReceiverNumber string `json:"receiver_number"`
}

// InitDatabase initializes the database connection
func InitDatabase() error {
	// Retrieve environment variables
	host := os.Getenv("DB_HOST")
	portStr := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")
	sslmode := os.Getenv("DB_SSLMODE")

	// Convert port to integer
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatal("Failed to convert port to integer: ", err)
	}
	// Construct the DSN
	dsn := "host=" + host + " port=" + strconv.Itoa(port) + " user=" + user +
		" dbname=" + dbname + " password=" + password + " sslmode=" + sslmode

	// Connect to the PostgreSQL database
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
		return err
	}

	// Auto migrate the database schema
	err = DB.AutoMigrate(&Message{})
	if err != nil {
		log.Fatal("Failed to migrate database schema: ", err)
		return err
	}

	return nil
}

func main() {
	r := gin.Default()

	// Initialize the database connection
	if err := InitDatabase(); err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	// API endpoints
	r.POST("/create", createMessage)
	r.GET("/get/messages/:account_id", getMessages)
	r.GET("/search", searchMessages)

	// Start the server
	r.Run(":8080")
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
	// Search by message_id
	if messageID := c.Query("message_id"); messageID != "" {
		DB.Where("message_id = ?", messageID).Find(&messages)
	}
	// Search by sender_number
	if senderNumber := c.Query("sender_number"); senderNumber != "" {
		DB.Where("sender_number = ?", senderNumber).Find(&messages)
	}
	// Search by receiver_number
	if receiverNumber := c.Query("receiver_number"); receiverNumber != "" {
		DB.Where("receiver_number = ?", receiverNumber).Find(&messages)
	}
	// Return the results
	c.JSON(http.StatusOK, messages)
}
