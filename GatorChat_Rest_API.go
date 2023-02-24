package main

//API IS INTERFACED/TESTED USING POSTMAN
import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// CONNECT TO USER MESSAGES DATABASE
var userMessagesDsn = "swegroup39:8wWrp52ey^2^@tcp(gator-chat.mysql.database.azure.com:3306)/user_messages?parseTime=true&loc=Local&tls=true&charset=utf8mb4"
var userMessagesDb, messageErr = gorm.Open(mysql.Open(userMessagesDsn), &gorm.Config{})

// CONNECT TO USER ACCOUNTS DATABASE
var userAccountsDsn = "swegroup39:8wWrp52ey^2^@tcp(gator-chat.mysql.database.azure.com:3306)/user_accounts?parseTime=true&loc=Local&tls=true&charset=utf8mb4"
var userAccountsDb, accountErr = gorm.Open(mysql.Open(userAccountsDsn), &gorm.Config{})

// MESSAGE STRUCT USED FOR EACH MESSAGE TABLE ENTRY
type UserMessage struct {
	gorm.Model
	//THE ACTUAL MESSAGE CONTENT
	Message string `json:"message"`
	//USER ID OF WHOEVER SENT THE MESSAGE
	Sender_ID string `json:"sender_id"`
	//USER ID OF WHOEVER RECEIVED THE MESSAGE
	Receiver_ID string `json:"receiver_id"`

	//MAYBE ADD A USERMESSAGE SLICE TO KEEP TRACK OF GROUP CHATS?
}

// USER STRUCT FOR EACH USER TABLE ENTRY
type UserAccount struct {
	// THE USER'S USERNAME
	Username string `json:"username"`
	// THE USER'S PASSWORD
	Password string `json:"password"`
	// THE USER'S ID
	User_ID string `json:"user_id"`
	// A SLICE OF USER ID'S THAT REPRESENT THE PEOPLE THEY ARE IN A CURRENT CONVERSATION WITH
	// JSON.RAWMESSAGE IS A TYPE THAT ALLOWS FOR ARRAYS OF STRINGS
	Current_Conversations json.RawMessage `json:"current_conversations"`
}

// // GETS A MESSAGE BASED ON THE GORM ID
//
//	func getMessage(w http.ResponseWriter, r *http.Request) {
//		w.Header().Set("Content-Type", "application/json")
//		params := mux.Vars(r)
//		var message UserMessage
//		result := db.First(&message, params["id"])
//		if result.Error != nil {
//			http.Error(w, "Message not found.", http.StatusNotFound)
//			return
//		}
//		json.NewEncoder(w).Encode(message)
//	}

// RETRIEVES ALL MESSAGES BETWEEN TWO PEOPLE
// REQUEST NEEDS TO PASS IN SENDER ID AND RECEIVER ID
func getConversation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var messages []UserMessage
	result := userMessagesDb.Where("(sender_id = ? OR receiver_id = ?) AND (sender_id = ? OR receiver_id = ?)", params["id_1"], params["id_1"], params["id_2"], params["id_2"]).Find(&messages)
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
	result := userMessagesDb.Find(&messages)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(messages)
}

// MESSAGE PASSED IN MUST USE "%20" FOR SPACES
func searchMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var messages []UserMessage
	// LIKE REQUIRES % TO BE SURROUNDED AROUND THE MESSAGE TO TELL IT TO FIND MESSAGES THAT CONTAIN IT, REGARDLESS OF WHERE IT IS
	searchQuery := "%" + params["search"] + "%"
	_ = userMessagesDb.Where("Message LIKE ?", searchQuery).Find(&messages)
	if len(messages) == 0 {
		http.Error(w, "No messages found.", http.StatusNotFound)
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

	// VALIDATE THE INFORMATION
	// CHECK IF THE ID'S ARE NUMERIC
	_, err = strconv.Atoi(message.Sender_ID)

	if err != nil {
		http.Error(w, "Invalid Sender ID (NOT NUMERIC): "+message.Sender_ID+" is not a valid ID.", http.StatusBadRequest)
		return
	}

	_, err = strconv.Atoi(message.Receiver_ID)

	if err != nil {
		http.Error(w, "Invalid Receiver ID (NOT NUMERIC): "+message.Receiver_ID+" is not a valid ID.", http.StatusBadRequest)
		return
	}

	// CHECK IF THE ID LENGTH IS FOUR
	if len(message.Sender_ID) != 4 {
		http.Error(w, "Invalid Sender ID (NOT FOUR DIGITS): "+message.Sender_ID+" is not a valid ID.", http.StatusBadRequest)
		return
	}

	if len(message.Receiver_ID) != 4 {
		http.Error(w, "Invalid Receiver ID (NOT FOUR DIGITS): "+message.Receiver_ID+" is not a valid ID.", http.StatusBadRequest)
		return
	}

	if len(message.Message) == 0 {
		http.Error(w, "Invalid Message: Messages cannot be empty.", http.StatusBadRequest)
		return
	}

	//IF IT PASSES THE CHECKS, THEN IT CREATES
	result := userMessagesDb.Create(&message)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(message)
	json.NewEncoder(w).Encode("Message created successfully.")
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
	result := userMessagesDb.Model(&UserMessage{}).Where("sender_id = ? AND receiver_id = ? AND message = ?", params["id_1"], params["id_2"], params["inputMessage"]).Update("Message", message.Message)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	if result.RowsAffected == 0 {
		http.Error(w, "Message not found.", http.StatusNotFound)
		return
	}

	// GET THE ENTRY AGAIN (WITH THE UPDATED MESSAGE) TO RETURN BACK TO THE REQUESTER
	userMessagesDb.Model(&UserMessage{}).Where("sender_id = ? AND receiver_id = ? AND message = ?", params["id_1"], params["id_2"], message.Message).First(&message)

	json.NewEncoder(w).Encode(message)
	json.NewEncoder(w).Encode("Message edited successfully.")
}

// SOFT DELETES A MESSAGE
//
//	func deleteMessage(w http.ResponseWriter, r *http.Request) {
//		w.Header().Set("Content-Type", "application/json")
//		params := mux.Vars(r)
//		var userMessage UserMessage
//		result := db.Where("id = ?", params["id"]).Delete(&userMessage)
//		if result.RowsAffected == 0 {
//			http.NotFound(w, r)
//			return
//		}
//		json.NewEncoder(w).Encode("Message deleted successfully.")
//	}
//

// HARD DELETES ALL MESSAGES WITH THE MATCHING SENDER AND RECEIVER ID (EFFECTIVELY CLEARS AN ENTIRE CONVERSATION)
func deleteMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var userMessage UserMessage
	result := userMessagesDb.Where("sender_id = ? AND receiver_id = ?", params["id_1"], params["id_2"]).Unscoped().Delete(&userMessage)
	if result.RowsAffected == 0 {
		http.NotFound(w, r)
		return
	}
	json.NewEncoder(w).Encode("Messages deleted successfully.")
}

// HARD DELETES A SPECIFIC MESSAGE BASED ON SENDER ID, USER ID, AND THE SPECIFIC MESSAGE
func deleteSpecificMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var userMessage UserMessage
	result := userMessagesDb.Where("sender_id = ? AND receiver_id = ? AND message = ?", params["id_1"], params["id_2"], params["inputMessage"]).Unscoped().Delete(&userMessage)
	if result.RowsAffected == 0 {
		http.NotFound(w, r)
		return
	}
	json.NewEncoder(w).Encode("Message deleted successfully.")
}

func deleteTable(w http.ResponseWriter, r *http.Request) {
	//CLEARS THE ENTIRE DATABASE - THIS WILL ONLY EVER BE USED FOR TESTING PURPOSES
	//THIS RESETS THE GORM ID BACK TO 1
	err := userMessagesDb.Exec("TRUNCATE TABLE user_messages").Error
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode("Table was deleted.")
}

func createUserAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var userAccount UserAccount
	err := json.NewDecoder(r.Body).Decode(&userAccount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result := userAccountsDb.Create(&userAccount)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(userAccount)
	json.NewEncoder(w).Encode("User account created successfully.")
}

// GETS ALL MESSAGES IN DATABASE
func getAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var messages []UserAccount
	result := userAccountsDb.Find(&messages)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(messages)
}

func main() {

	//CHECK IF THERE WERE ERRORS WITH CONNECTING
	if messageErr != nil {
		panic("Error: Failed to connect to database.")
	}

	if accountErr != nil {
		panic("Error: Failed to connect to database.")
	}

	// INIT ROUTER
	r := mux.NewRouter()

	// AUTO MIGRATE CURRENTLY NOT WORKING...
	// err = db.AutoMigrate(&UserMessage{})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// TEXT ROUTE HANDLERS / ENDPOINTS

	//POST FUNCTIONS
	r.HandleFunc("/api/messages", createMessage).Methods("POST")

	//GET FUNCTIONS
	r.HandleFunc("/api/messages/{id_1}/{id_2}", getConversation).Methods("GET")
	r.HandleFunc("/api/messages", getAllMessages).Methods("GET")
	r.HandleFunc("/api/messages/{search}", searchMessage).Methods("GET")

	//PUT FUNCTIONS
	r.HandleFunc("/api/messages/{id_1}/{id_2}/{inputMessage}", editMessage).Methods("PUT")

	//DELETE FUNCTIONS
	r.HandleFunc("/api/messages/{id_1}/{id_2}", deleteMessage).Methods("DELETE")
	r.HandleFunc("/api/messages/{id_1}/{id_2}/{inputMessage}", deleteSpecificMessage).Methods("DELETE")
	r.HandleFunc("/api/messages/deleteTable", deleteTable).Methods("DELETE")

	r.HandleFunc("/api/users", createUserAccount).Methods("POST")
	r.HandleFunc("/api/users", getAllUsers).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}
