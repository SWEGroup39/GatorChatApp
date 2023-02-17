package main

//API IS INTERFACED USING POSTMAN

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// CONNECT TO MYSQL DATABASE USING MICROSOFT AZURE
var dsn = "swegroup39:8wWrp52ey^2^@tcp(gator-chat.mysql.database.azure.com:3306)/user_messages?parseTime=true&loc=Local&tls=true&charset=utf8mb4"
var db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

// MESSAGE STRUCT USED FOR EACH TABLE ENTRY
type UserMessage struct {
	gorm.Model

	//THE ACTUAL MESSAGE CONTENT
	Message string `json:"message"`

	//USER ID OF WHOEVER SENT THE MESSAGE
	Sender_ID string `json:"sender_id"`

	//USER ID OF WHOEVER RECEIVED THE MESSAGE
	Receiver_ID string `json:"receiver_id"`
}

// GETS A MESSAGE BASED ON THE GORM ID
func getMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var message UserMessage
	result := db.First(&message, params["id"])
	if result.Error != nil {
		http.Error(w, "Message not found.", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(message)
}

// GETS ALL MESSAGES IN DATABASE
func getAllMessages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var messages []UserMessage
	result := db.Find(&messages)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(messages)
}

// CREATES A NEW ENTRY IN THE DATABASE
func createMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var message UserMessage
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result := db.Create(&message)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(message)
}

// UPDATES ONLY THE MESSAGE FIELD IN THE ENTRY
func editMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var message UserMessage
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result := db.Model(&UserMessage{}).Where("id = ?", params["id"]).Update("Message", message.Message)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	if result.RowsAffected == 0 {
		http.Error(w, "Message not found.", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(message)
}

// SOFT DELETES A MESSAGE
//func deleteMessage(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	params := mux.Vars(r)
//	var userMessage UserMessage
//	result := db.Where("id = ?", params["id"]).Delete(&userMessage)
//	if result.RowsAffected == 0 {
//		http.NotFound(w, r)
//		return
//	}
//	json.NewEncoder(w).Encode("Message deleted successfully.")
//}

// HARD DELETES A MESSAGE (the ID can be reused after this)
func deleteMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var userMessage UserMessage
	result := db.Where("id = ?", params["id"]).Unscoped().Delete(&userMessage)
	if result.RowsAffected == 0 {
		http.NotFound(w, r)
		return
	}
	json.NewEncoder(w).Encode("Message deleted successfully.")
}

func main() {
	// INIT ROUTER
	r := mux.NewRouter()

	// AUTO MIGRATE CURRENTLY NOT WORKING...

	// err = db.AutoMigrate(&UserMessage{})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// TEXT ROUTE HANDLERS / ENDPOINTS
	r.HandleFunc("/api/messages/{id}", getMessage).Methods("GET")
	r.HandleFunc("/api/messages", getAllMessages).Methods("GET")
	r.HandleFunc("/api/messages", createMessage).Methods("POST")
	r.HandleFunc("/api/messages/{id}", editMessage).Methods("PUT")
	r.HandleFunc("/api/messages/{id}", deleteMessage).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
