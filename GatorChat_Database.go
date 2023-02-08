package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DATABASE DETAILS
var dsn = "swegroup39:8wWrp52ey^2^@tcp(gator-chat.mysql.database.azure.com:3306)/user_messages?parseTime=true&tls=true&charset=utf8mb4"

// CONNECTS TO THE MySQL DATABASE ON A HOST AND PORT, USING THE DATABASE NAME, USERNAME, AND PASSWORD
var db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

// MAKE A USER STRUCT THAT STORES A MESSAGE STRUCT
// TABLE NAME IS THE PLURALIZED VERSION OF THE STRUCT NAME IN SNAKE CASE
type UserMessage struct {
	gorm.Model

	//THE ACTUAL MESSAGE CONTENT
	Message string

	//USER ID OF WHOEVER SENT THE MESSAGE
	Sender_ID string

	//USER ID OF WHOEVER RECEIVED THE MESSAGE
	Receiver_ID string
}

func main() {

	//CHECK IF THERE WERE ERRORS WITH CONNECTING
	if err != nil {
		panic("Error: Failed to connect to database.")
	}

	//CLEARS THE ENTIRE DATABASE - DO NOT ACTUALLY DO THIS; I AM CLEARING THIS SO THE TEST MESSAGES ARE NOT OVERWHELMING
	err = db.Exec("DELETE FROM user_messages").Error
	if err != nil {
		panic(err)
	}

	userExample := UserMessage{
		Sender_ID:   "1234",
		Receiver_ID: "4321",
		Message:     "Hello",
	}

	anotherExample := UserMessage{
		Sender_ID:   "4567",
		Receiver_ID: "7654",
		Message:     "Hi",
	}

	//GENERATE MESSAGE ENTRIES
	db.Create(&userExample)
	db.Create(&anotherExample)

	//GENERATES A NEW TABLE IF ONE WAS NOT MADE OR MIGRATES ALL EXISTING DATA IF IT DOES EXIST
	//db.AutoMigrate(&UserMessage{})

	//QUERY EXAMPLE: FIRST ENTRY IN THE TABLE
	var textMessage UserMessage
	db.First(&textMessage, 117)

	fmt.Println("The first message in the database:")
	fmt.Println("Message: ", textMessage.Message)
	fmt.Println("Time Created: ", textMessage.CreatedAt)

	//UPDATE THE TEXT MESSAGE OF THE TEXTMESSAGE VARIABLE
	textMessage.Message = "New Hi"
	db.Save(&textMessage)

	fmt.Println("Updated message:")
	fmt.Println("Message: ", textMessage.Message)

	//CLEAR THE VARIABLE - IMPORTANT BECAUSE OLD DATA WILL PERSIST IF NOT SPECIFICALLY UPDATED
	textMessage = UserMessage{}

	//QUERY EXAMPLE: LOCATE THE FIRST MESSAGE IN THE DATABASE THAT HAS THE MATCHING TEXT
	db.Where("Message = ?", "Hi").First(&textMessage)

	//%v IS A FORMAT VERB THAT WILL BE REPLACED WITH THE RESPECTIVE CONTENT AFTER THE COMMA
	fmt.Println("The first instance of 'Hi' in the database.")
	fmt.Printf("%v (%v)", textMessage.Message, textMessage.CreatedAt)
	fmt.Println()

	printMessages(db)

	//TEST DELETE (SOFT)
	textMessage = UserMessage{}
	db.Where("Message = ?", "Hi").First(&textMessage)
	fmt.Println("Soft deleting: ", textMessage.Message)
	db.Delete(&textMessage)

	printMessages(db)
}

// PRINT ALL MESSAGES IN DATABASE
func printMessages(db *gorm.DB) {
	//MAKE AN EMPTY CONTAINER TO STORE MESSAGES
	var messages []UserMessage
	db.Find(&messages)
	fmt.Println("All messages:")
	for _, message := range messages {
		fmt.Println("ID:", message.ID, "| Message:", message.Message, "| SenderID:", message.Sender_ID, "| Time Stamp:", message.CreatedAt)
	}
}

// THIS TAKES IN A MESSAGE AND THE DATA CAN BE RETRIEVED/UPDATED/DELETED THROUGH HTTP REQUESTS (ENABLES CREATING A REST API - NEXT STEP)
func insertMessage(userExample UserMessage) {
	http.HandleFunc("/createstuff", func(w http.ResponseWriter, r *http.Request) {
		createEntry(w, r, userExample)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func createEntry(w http.ResponseWriter, r *http.Request, userExample UserMessage) {
	db.Create(&userExample)

	if err := db.Create(&userExample).Error; err != nil {
		log.Fatalln((err))
	}

	json.NewEncoder(w).Encode(userExample)

	fmt.Println("Fields Added", userExample)
}
