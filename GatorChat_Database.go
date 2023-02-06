package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Message struct {
	//GORM.MODEL INCLUDES TIMESTAMPS AS WELL AS AN ID
	gorm.Model

	//USER ID OF WHOEVER SENT THE MESSAGE
	SenderID uint

	//USER ID OF WHOEVER RECEIVED THE MESSAGE
	ReceiverID uint

	//THE ACTUAL MESSAGE CONTENT
	Text string
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("Error: Failed to connect to database.")
	}

	db.AutoMigrate(&Message{})

	//GENERATE MESSAGE ENTRIES
	db.Create(&Message{Text: "Test Message", SenderID: 12345, ReceiverID: 1359})

	//QUERY FOR THE FIRST ENTRY IN THE TABLE
	var textMessage Message
	db.First(&textMessage, 1)

	textMessage.Text = "Replacement Text"
	db.Save(&textMessage)

	fmt.Println("db.First(&textMessage, 1) returns: ")
	fmt.Println("Message: ", textMessage.Text)
	fmt.Println("Message: ", textMessage.CreatedAt)
}
