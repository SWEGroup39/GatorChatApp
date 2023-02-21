package main

//API IS INTERFACED/TESTED USING POSTMAN

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

// // GETS A MESSAGE BASED ON THE GORM ID
// func getMessage(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	var message UserMessage
// 	result := db.First(&message, params["id"])
// 	if result.Error != nil {
// 		http.Error(w, "Message not found.", http.StatusNotFound)
// 		return
// 	}
// 	json.NewEncoder(w).Encode(message)
// }

// RETRIEVES ALL MESSAGES BETWEEN TWO PEOPLE
// REQUEST NEEDS TO PASS IN SENDER ID AND RECEIVER ID
func getConversation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var messages []UserMessage
	result := db.Where("(sender_id = ? OR receiver_id = ?) AND (sender_id = ? OR receiver_id = ?)", params["id_1"], params["id_1"], params["id_2"], params["id_2"]).Find(&messages)
	if result.Error != nil {
		http.Error(w, "Messages not found.", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(messages)
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
// REQUEST NEEDS TO PASS IN SENDER ID, RECEIVER ID, AND MESSAGE
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
// REQUEST NEEDS TO PASS IN SENDER ID, RECEIVER ID, AND MESSAGE
func editMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var message UserMessage
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result := db.Model(&UserMessage{}).Where("sender_id = ? AND receiver_id = ?", params["id_1"], params["id_2"]).Update("Message", message.Message)
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

// HARD DELETES A MESSAGE (ID CAN BE REUSED AFTER THIS)
func deleteMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var userMessage UserMessage
	result := db.Where("sender_id = ? AND receiver_id = ?", params["id_1"], params["id_2"]).Unscoped().Delete(&userMessage)
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
	r.HandleFunc("/api/messages/{id_1}/{id_2}", getConversation).Methods("GET")
	r.HandleFunc("/api/messages", getAllMessages).Methods("GET")
	r.HandleFunc("/api/messages", createMessage).Methods("POST")
	r.HandleFunc("/api/messages/{id_1}/{id_2}", editMessage).Methods("PUT")
	r.HandleFunc("/api/messages/{id_1}/{id_2}", deleteMessage).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
