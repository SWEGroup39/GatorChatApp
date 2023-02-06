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

	//CONNECTS TO THE SQLITE DATABASE KNOWN AS "test.db"
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("Error: Failed to connect to database.")
	}

	//TABLE NAME IS THE PLURALIZED VERSION OF THE STRUCT NAME IN SNAKE CASE
	//CLEARS THE ENTIRE DATABASE
	err = db.Exec("DELETE FROM messages").Error
	if err != nil {
		panic(err)
	}

	//GENERATES A NEW TABLE IF ONE WAS NOT MADE OR MIGRATES ALL EXISTING DATA IF IT DOES EXIST
	db.AutoMigrate(&Message{})

	//GENERATE MESSAGE ENTRIES
	db.Create(&Message{Text: "Hello", SenderID: 12345, ReceiverID: 1359})

	//QUERY FOR THE FIRST ENTRY IN THE TABLE
	var textMessage Message
	db.First(&textMessage, 1)

	//UPDATE THE TEXT MESSAGE OF THE TEXTMESSAGE VARIABLE
	textMessage.Text = "Hi"
	db.Save(&textMessage)

	fmt.Println("The first message in the database.")
	fmt.Println("Message: ", textMessage.Text)
	fmt.Println("Message: ", textMessage.CreatedAt)

	//CLEAR THE VARIABLE
	textMessage = Message{}
	//LOCATE THE FIRST MESSAGE IN THE DATABASE THAT HAS THE MATCHING TEXT
	db.Where("Text = ?", "Hi").First(&textMessage)
	//%v IS A FORMAT VERB THAT WILL BE REPLACED WITH THE RESPECTIVE CONTENT AFTER THE COMMA
	fmt.Println("The first instance of 'Hi' in the database.")
	fmt.Printf("%v (%v)", textMessage.Text, textMessage.CreatedAt)
	fmt.Println()

	//PRINT ALL MESSAGES IN DATABASE
	var messages []Message
	db.Find(&messages)
	fmt.Println("All messages:")
	for _, message := range messages {
		fmt.Println("Message:", message.Text, "SenderID:", message.SenderID, "Time Stamp:", message.CreatedAt)
	}

}
