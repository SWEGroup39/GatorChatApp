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

	"github.com/rs/cors"
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

// RETRIEVES ALL MESSAGES BETWEEN TWO PEOPLE
// REQUEST NEEDS TO PASS IN SENDER ID AND RECEIVER ID
func getConversation(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting Conversations (GET)")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	var messages []UserMessage
	result := userMessagesDb.Where("(sender_id = ? OR receiver_id = ?) AND (sender_id = ? OR receiver_id = ?)", params["id_1"], params["id_1"], params["id_2"], params["id_2"]).Find(&messages)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	if len(messages) == 0 {
		http.Error(w, "Conversation not found.", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(messages)
	log.Println("Got Conversation successfully.")
}

// GETS ALL MESSAGES IN DATABASE
func getAllMessages(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting All Messages (GET)")
	w.Header().Set("Content-Type", "application/json")

	var messages []UserMessage
	result := userMessagesDb.Find(&messages)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	if len(messages) == 0 {
		http.Error(w, "Messages not found.", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(messages)
	log.Println("Got All Messages sucessfully.")
}

// MESSAGE PASSED IN MUST USE "%20" FOR SPACES
func searchMessage(w http.ResponseWriter, r *http.Request) {
	log.Println("Searching for messsage (GET)")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	var messages []UserMessage

	// ADD % SO IT CAN FIND IT ANYWHERE
	searchTerm := "%" + params["search"] + "%"

	// USE REGEX TO FIND MESSAGE THAT IS A WHOLE WORD AND EXISTS EITHER AT THE VERY START, MIDDLE, OR VERY END OF MESSAGE
	userMessagesDb.Where("Message RLIKE CONCAT('(^|\\s)', ? ,'(\\s|$)') OR Message LIKE ?", params["search"], searchTerm).Find(&messages)

	if len(messages) == 0 {
		http.Error(w, "No messages found.", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(messages)
	log.Println("Found message(s) successfully.")
}

// CREATES A NEW ENTRY IN THE DATABASE
// REQUEST NEEDS TO PASS IN SENDER ID, RECEIVER ID, AND MESSAGE
func createMessage(w http.ResponseWriter, r *http.Request) {
	log.Println("Sending a Message (POST)")
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
	log.Println("Message created successfully.")
}

// UPDATES ONLY THE MESSAGE FIELD IN THE ENTRY
// REQUEST NEEDS TO PASS IN SENDER ID, RECEIVER ID, AND MESSAGE
func editMessage(w http.ResponseWriter, r *http.Request) {
	log.Println("Editing a Message (PUT)")
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
	log.Println("Message edited successfully.")
}

// HARD DELETES ALL MESSAGES WITH THE MATCHING SENDER AND RECEIVER ID (EFFECTIVELY CLEARS AN ENTIRE CONVERSATION)
func deleteConversation(w http.ResponseWriter, r *http.Request) {
	log.Println("Deleting a Conversation (DELETE)")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	var userMessage UserMessage
	result := userMessagesDb.Where("(sender_id = ? OR receiver_id = ?) AND (sender_id = ? OR receiver_id = ?)", params["id_1"], params["id_1"], params["id_2"], params["id_2"]).Unscoped().Delete(&userMessage)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	if result.RowsAffected == 0 {
		http.Error(w, "Conversation not found.", http.StatusNotFound)
		return
	}

	log.Println("Conversation deleted successfully.")
}

// HARD DELETES A SPECIFIC MESSAGE BASED ON SENDER ID, USER ID, AND THE SPECIFIC MESSAGE
func deleteSpecificMessage(w http.ResponseWriter, r *http.Request) {
	log.Println("Deleting a Message (DELETE)")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	var userMessage UserMessage
	result := userMessagesDb.Where("sender_id = ? AND receiver_id = ? AND message = ?", params["id_1"], params["id_2"], params["inputMessage"]).Unscoped().Delete(&userMessage)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	if result.RowsAffected == 0 {
		http.Error(w, "Message not found.", http.StatusNotFound)
		return
	}
	log.Println("Message deleted successfully.")
}

func deleteTable(w http.ResponseWriter, r *http.Request) {
	//CLEARS THE ENTIRE DATABASE - THIS WILL ONLY EVER BE USED FOR TESTING PURPOSES
	//THIS RESETS THE GORM ID BACK TO 1
	log.Println("Deleting the Entire Database (DELETE)")
	err := userMessagesDb.Exec("TRUNCATE TABLE user_messages").Error
	if err != nil {
		panic(err)
	}
	log.Println("Database deleted successfully.")
}

func createUserAccount(w http.ResponseWriter, r *http.Request) {
	log.Println("Creating a User Account (POST)")
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
	log.Println("User account created successfully.")
}

// GETS ALL MESSAGES IN DATABASE
func getAllUsers(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting All Users (GET)")
	w.Header().Set("Content-Type", "application/json")
	var users []UserAccount
	result := userAccountsDb.Find(&users)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	if len(users) == 0 {
		http.Error(w, "Users not found.", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(users)
	log.Println("Found All Users successfully.")
}

func getUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting a User (GET)")
	w.Header().Set("Content-Type", "application/json")

	var user UserAccount
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := userAccountsDb.Model(&UserAccount{}).Where("username = ? AND password = ?", user.Username, user.Password).First(&user)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
	log.Println("Found User successfully.")
}

func addConversation(w http.ResponseWriter, r *http.Request) {
	log.Println("Adding an ID to a User's Conversation List (PUT)")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user UserAccount

	//FIND THE MATCHING USER STRUCT WITH THE MATCHING USER_ID
	result := userAccountsDb.Model(&UserAccount{}).Where("user_id = ?", params["id_1"]).First(&user)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	// CREATE A SLICE OF STRINGS
	var conversationSlice []string
	// CONVERT THE JSON INTO A STRING SLICE
	err := json.Unmarshal([]byte(user.Current_Conversations), &conversationSlice)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// ADD THE NEW ID TO THE STRING SLICE
	conversationSlice = append(conversationSlice, params["id_2"])

	// CONVERT THE STRING SLICE BACK INTO A JSON AND UPDATE THE USER'S CONVERSATION JSON
	conversationJSON, _ := json.Marshal(conversationSlice)

	// REPLACE THE JSON WITH A NEW JSON THAT CONTAINS THE ADDED ID
	result = userAccountsDb.Model(&UserAccount{}).Where("user_id = ?", params["id_1"]).Update("Current_Conversations", string(conversationJSON))

	if result.Error != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//GET THE UPDATED USER STRUCT
	userAccountsDb.Model(&UserAccount{}).Where("user_id = ?", params["id_1"]).First(&user)

	json.NewEncoder(w).Encode(user)
	log.Println("ID added successfully.")
}

func main() {
	log.Println("Connecting to API...")

	//CHECK IF THERE WERE ERRORS WITH CONNECTING
	if messageErr != nil {
		panic("Error: Failed to connect to messages database.")
	}

	if accountErr != nil {
		panic("Error: Failed to connect to users database.")
	}

	// INIT ROUTER
	r := mux.NewRouter()

	log.Println("API Connected.")

	corsHandler := cors.Default().Handler

	// AUTO MIGRATE CURRENTLY NOT WORKING...
	// err = db.AutoMigrate(&UserMessage{})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// API FUNCTIONS FOR THE MESSAGES DATABASE

	// POST FUNCTIONS
	r.HandleFunc("/api/messages", createMessage).Methods("POST")

	// GET FUNCTIONS
	r.HandleFunc("/api/messages/{id_1}/{id_2}", getConversation).Methods("GET")
	r.HandleFunc("/api/messages", getAllMessages).Methods("GET")
	r.HandleFunc("/api/messages/{search}", searchMessage).Methods("GET")

	// PUT FUNCTIONS
	r.HandleFunc("/api/messages/{id_1}/{id_2}/{inputMessage}", editMessage).Methods("PUT")

	// DELETE FUNCTIONS
	r.HandleFunc("/api/messages/{id_1}/{id_2}", deleteConversation).Methods("DELETE")
	r.HandleFunc("/api/messages/{id_1}/{id_2}/{inputMessage}", deleteSpecificMessage).Methods("DELETE")
	r.HandleFunc("/api/messages/deleteTable", deleteTable).Methods("DELETE")

	// API FUNCTIONS FOR THE USERS DATABASE

	// POST FUNCTIONS
	r.HandleFunc("/api/users", createUserAccount).Methods("POST")

	// PUT FUNCTIONS
	r.HandleFunc("/api/users/{id_1}/{id_2}", addConversation).Methods("PUT")

	// GET FUNCTIONS
	r.HandleFunc("/api/users", getAllUsers).Methods("GET")
	r.HandleFunc("/api/users/User", getUser).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", corsHandler(r)))
}
