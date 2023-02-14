package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Connects to the MySQL database on a host and port, using the database name, username, and password
var dsn = "swegroup39:8wWrp52ey^2^@tcp(gator-chat.mysql.database.azure.com:3306)/user_messages?parseTime=true&tls=true&charset=utf8mb4"
var db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

type UserMessage struct {
	gorm.Model
	Message    string `json:"message"`
	SenderID   string `json:"sender_id"`
	ReceiverID string `json:"receiver_id"`
}

// Get a specific message by ID
func getMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var message UserMessage
	result := db.First(&message, params["id"])
	if result.Error != nil {
		http.Error(w, "Message not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(message)
}

// Create a new message
func createMessage(w http.ResponseWriter, r *http.Request) {
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

// Edit a message
func editMessage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var message UserMessage
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result := db.Model(&UserMessage{}).Where("id = ?", params["id"]).Updates(&message)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	if result.RowsAffected == 0 {
		http.Error(w, "Message not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(message)
}

// Delete a Message
func deleteMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var userMessage UserMessage
	result := db.Where("id = ?", params["id"]).Delete(&userMessage)
	if result.RowsAffected == 0 {
		http.NotFound(w, r)
		return
	}
	json.NewEncoder(w).Encode("Message deleted successfully")
}

func main() {
	// Init Router
	r := mux.NewRouter()

	// // Connect to MySQL database
	// err = db.AutoMigrate(&UserMessage{})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Text Route Handlers / Endpoints
	r.HandleFunc("/api/messages/{id}", getMessage).Methods("GET")
	r.HandleFunc("/api/messages", createMessage).Methods("POST")
	r.HandleFunc("/api/messages/{id}", editMessage).Methods("PUT")
	r.HandleFunc("/api/messages/{id}", deleteMessage).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
