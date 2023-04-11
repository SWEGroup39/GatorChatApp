package main

//API IS INTERFACED/TESTED USING POSTMAN
import (
	"github.com/Azure/azure-storage-blob-go/azblob"

	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/rs/cors"

	"crypto/sha256"
	"encoding/hex"
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
	// THE ACTUAL MESSAGE CONTENT
	Message string `json:"message"`
	// USER ID OF WHOEVER SENT THE MESSAGE
	Sender_ID string `json:"sender_id"`
	// USER ID OF WHOEVER RECEIVED THE MESSAGE
	Receiver_ID string `json:"receiver_id"`
	// THE IMAGE AS A BYTE ARRAY
	Image []byte `json:"image"`
}

// USER STRUCT FOR EACH USER TABLE ENTRY
type UserAccount struct {
	// THE USER'S USERNAME
	Username string `json:"username"`

	// THE USER'S PASSWORD
	Password string `json:"password"`

	// THE USER'S ID
	User_ID string `json:"user_id"`

	// THE USER'S EMAIL ADDRESS
	Email string `json:"email"`

	// THE USER'S FULL NAME
	Full_Name string `json:"full_name"`

	// THE USER'S PHONE NUMBER
	Phone_Number string `json:"phone_number"`

	// A SLICE OF USER ID'S THAT REPRESENT THE PEOPLE THEY ARE IN A CURRENT CONVERSATION WITH
	// JSON.RAWMESSAGE IS A TYPE THAT ALLOWS FOR ARRAYS OF STRINGS
	Current_Conversations json.RawMessage `json:"current_conversations"`
}

type UserAccountConfirmPass struct {
	// INCLUDE THE USERACCOUNT STRUCT
	UserAccount

	// ADD AN ADDITIONAL FIELD FOR THE OG PASSWORD (USED IN THE EDITPASS FUNCTION)
	OriginalPassword string `json:"original_pass"`
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

func getConversationLP(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting Conversations (GET (LONG POLL))")

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	// CREATE A CHANNEL THAT FACILITATES INFO BETWEEN API AND FRONTEND
	channel := make(chan []UserMessage)

	// CREATE A GOROUTINE THAT HANDLES THE GETCONVERSATION FUNCTIONALITY
	// THIS IS ACTUALLY DOING THE WORK
	go func() {
		for {
			// SEE IF THERE ARE NEW MESSAGES
			var messages []UserMessage
			result := userMessagesDb.Where("(sender_id = ? OR receiver_id = ?) AND (sender_id = ? OR receiver_id = ?)", params["id_1"], params["id_1"], params["id_2"], params["id_2"]).Find(&messages)
			if result.Error != nil {
				log.Println(result.Error)
			}

			// SEND CONVERSATION TO FRONTEND (WON'T SEND IF THERE ARE ZERO MESSAGES IN THE CONVERSATION)
			if len(messages) > 0 {
				channel <- messages
				return
			}

			// WAIT ONE SECOND THEN DO IT AGAIN
			time.Sleep(time.Second)
		}
	}()

	// LISTENS TO SEE IF A NEW MESSAGE IS AVAILABLE (GETTING IT FROM THE GOROUTINE)
	// ALSO SEES IF THE CONNECTION WAS CLOSED
	for {
		select {
		case messages := <-channel:
			// SEND NEW MESSAGE TO FRONTEND THROUGH JSON
			json.NewEncoder(w).Encode(messages)
			log.Println("Updated conversation successfully.")
			return
		case <-r.Context().Done():
			// CLOSE LONG POLLING IF FRONTEND CLOSES CONNECTION
			log.Println("Long polling closed.")
			return
		}
	}
}

// GETS ALL NON-DELETED MESSAGES IN DATABASE
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

// TEST FUNCTION TO SEE SOFT DELETED MESSAGES
func getDeletedMessages(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting Deleted Messages (GET)")
	w.Header().Set("Content-Type", "application/json")

	var messages []UserMessage
	result := userMessagesDb.Unscoped().Where("deleted_at IS NOT NULL").Find(&messages)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	if len(messages) == 0 {
		http.Error(w, "Deleted messages not found.", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(messages)
	log.Println("Got Deleted Messages successfully.")
}

// TEST FUNCTION TO REMOVE ALL SOFT DELETED MESSAGES
func deleteDeletedMessages(w http.ResponseWriter, r *http.Request) {
	log.Println("Deleting Deleted Messages (DELETE)")
	w.Header().Set("Content-Type", "application/json")

	var messages []UserMessage
	result := userMessagesDb.Unscoped().Where("deleted_at IS NOT NULL").Unscoped().Delete(&messages)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	if result.RowsAffected == 0 {
		http.Error(w, "Messages not found.", http.StatusNotFound)
		return
	}
	log.Println("Message(s) deleted successfully.")
}

func searchAllConversations(w http.ResponseWriter, r *http.Request) {
	log.Println("Searching for messsage in every conversation (GET)")
	w.Header().Set("Content-Type", "application/json")

	var message UserMessage
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var messages []UserMessage

	// USE REGEX TO FIND MESSAGE THAT IS A WHOLE WORD AND EXISTS EITHER AT THE VERY START, MIDDLE, OR VERY END OF MESSAGE
	// IT LOOKS FOR MESSAGES THAT HAVE MATCHING WORD AND SURROUNDED BY EITHER A START/END OR A NON-ALPHABETIC CHARACTER
	// QUOTEMETA USED FOR SPECIAL CHARACTERS LIKE "?"
	searchTerm := "(^|[^[:alnum:]])" + regexp.QuoteMeta(message.Message) + "([^[:alnum:]]*|$)"

	userMessagesDb.Where("Message RLIKE ?", searchTerm).Find(&messages)

	if len(messages) == 0 {
		http.Error(w, "No messages found.", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(messages)
	log.Println("Found message(s) successfully.")
}

func searchOneConversation(w http.ResponseWriter, r *http.Request) {
	log.Println("Searching for messsage in one conversation (POST)")
	w.Header().Set("Content-Type", "application/json")

	var message UserMessage
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	params := mux.Vars(r)
	var messages []UserMessage

	searchTerm := "(^|[^[:alpha:]])" + regexp.QuoteMeta(message.Message) + "([^[:alpha:]]|$)"

	// FIND MESSAGES IN INSTANCES ONLY WHERE THE ID'S MATCH
	userMessagesDb.Where("(sender_id = ? OR receiver_id = ?) AND (sender_id = ? OR receiver_id = ?) AND (Message RLIKE ?)", params["id_1"], params["id_1"], params["id_2"], params["id_2"], searchTerm).Find(&messages)

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

	imageString := base64.StdEncoding.EncodeToString(message.Image)

	imageData, err := base64.StdEncoding.DecodeString(imageString)
	if err != nil {
		http.Error(w, "Invalid image data: "+err.Error(), http.StatusBadRequest)
		return
	}

	message.Image = imageData

	//IF IT PASSES THE CHECKS, THEN IT CREATES
	result := userMessagesDb.Create(&message)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(message)
	log.Println("Message created successfully.")
}

// CREATE MESSAGE (WITH IMAGE)
func createMessageImage(w http.ResponseWriter, r *http.Request) {
	log.Println("Sending a Message (POST)")
	w.Header().Set("Content-Type", "application/json")

	var message UserMessage
	message.Sender_ID = r.FormValue("sender_id")
	message.Receiver_ID = r.FormValue("receiver_id")
	message.Message = r.FormValue("message")

	file, header, err := r.FormFile("image")

	if err != nil {
		http.Error(w, "Failed to retrieve image file: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	maxFileSize := 10 << 20
	if header.Size > int64(maxFileSize) {
		http.Error(w, "Image file is too large. Maximum size allowed is "+strconv.Itoa(maxFileSize)+" bytes", http.StatusBadRequest)
		return
	}

	// CREATE AN AZURE BLOB SERVICE CLIENT
	accountName := "gatorchatimages"
	accountKey := "Bcm07S+0vAEG45zbvUBQ3a7ujiXB1s9RxaJVAqCcKKTtbeGuu+22aQWyH8Ux++PAGBVkqfnSptcs+AStLXHyzg=="
	credential, err := azblob.NewSharedKeyCredential(accountName, accountKey)
	if err != nil {
		http.Error(w, "Failed to create shared key credential: "+err.Error(), http.StatusInternalServerError)
		return
	}
	p := azblob.NewPipeline(credential, azblob.PipelineOptions{})
	containerName := "images"
	blobURLParts := azblob.NewBlobURLParts(url.URL{
		Scheme: "https",
		Host:   accountName + ".blob.core.windows.net",
	})
	serviceURL := azblob.NewServiceURL(blobURLParts.URL(), p)
	containerURL := serviceURL.NewContainerURL(containerName)

	// UPLOAD THE IMAGE AS A BLOB
	blobName := uuid.New().String() + filepath.Ext(header.Filename)
	blobURL := containerURL.NewBlockBlobURL(blobName)
	ctx := context.Background()
	_, err = azblob.UploadStreamToBlockBlob(ctx, file, blobURL, azblob.UploadStreamToBlockBlobOptions{
		BlobHTTPHeaders: azblob.BlobHTTPHeaders{
			ContentType: header.Header.Get("Content-Type"),
		},
	})
	if err != nil {
		http.Error(w, "Failed to upload image file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// SET THE MESSAGE IMAGE URL TO THE BLOB URL
	message.Image = []byte(blobURL.String())

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
	log.Println("Message with image created successfully.")
}

// UPDATES ONLY THE MESSAGE FIELD IN THE ENTRY
// REQUEST NEEDS TO PASS IN UNIQUE GORM ID
func editMessage(w http.ResponseWriter, r *http.Request) {
	log.Println("Editing a Message (PUT)")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	var count int64
	userMessagesDb.Model(&UserMessage{}).Where("id = ?", params["id"]).Count(&count)
	if count == 0 {
		http.Error(w, "Message not found.", http.StatusNotFound)
		return
	}

	var message UserMessage
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result := userMessagesDb.Model(&UserMessage{}).Where("id = ?", params["id"]).Update("Message", message.Message)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	// GET THE ENTRY AGAIN (WITH THE UPDATED MESSAGE) TO RETURN BACK TO THE REQUESTER
	userMessagesDb.Model(&UserMessage{}).Where("id = ?", params["id"]).First(&message)

	json.NewEncoder(w).Encode(message)
	log.Println("Message edited successfully.")
}

func editName(w http.ResponseWriter, r *http.Request) {
	log.Println("Editing a Username (PUT)")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	var count int64
	userAccountsDb.Model(&UserAccount{}).Where("user_id = ?", params["id"]).Count(&count)
	if count == 0 {
		http.Error(w, "User not found.", http.StatusNotFound)
		return
	}

	var user UserAccount
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result := userAccountsDb.Model(&UserAccount{}).Where("user_id = ?", params["id"]).Update("Username", user.Username)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	userAccountsDb.Model(&UserAccount{}).Where("user_id = ?", params["id"]).First(&user)
	json.NewEncoder(w).Encode(user)
	log.Println("Username edited successfully.")
}

func editPass(w http.ResponseWriter, r *http.Request) {
	log.Println("Editing a Password (PUT)")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	var count int64
	userAccountsDb.Model(&UserAccount{}).Where("user_id = ?", params["id"]).Count(&count)
	if count == 0 {
		http.Error(w, "User not found.", http.StatusNotFound)
		return
	}

	var user UserAccountConfirmPass
	var temp UserAccount

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// HASH THE OG PASSWORD TO VERIFY THAT IT IS CORRECT
	hashedOriginalPassword := hashPassword(user.OriginalPassword)

	result := userAccountsDb.Model(&UserAccount{}).Where("user_id = ? AND password = ?", params["id"], hashedOriginalPassword).First(&temp)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		log.Println("Incorrect ID or Password")
		return
	}

	hashedNewPassword := hashPassword(user.Password)

	// CHECK IF THE NEW PASSWORD IS THE SAME AS THE ORIGINAL
	if user.Password == user.OriginalPassword {
		http.Error(w, "New password cannot be the same as the old password.", http.StatusBadRequest)
		log.Println("New password cannot be the same as the old password.")
		return
	}

	result = userAccountsDb.Model(&UserAccount{}).Where("user_id = ? AND password = ?", params["id"], hashedOriginalPassword).Update("Password", hashedNewPassword)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	userAccountsDb.Model(&UserAccount{}).Where("user_id = ?", params["id"]).First(&temp)
	json.NewEncoder(w).Encode(temp)
	log.Println("Password edited successfully.")
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

// HARD DELETES A SPECIFIC MESSAGE BASED ON UNIQUE GORM ID
func deleteSpecificMessage(w http.ResponseWriter, r *http.Request) {
	log.Println("Deleting a Message (DELETE)")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	checkRecentlyDeleted(params["id"], w, r)

	// NOW SOFT DELETE THIS MESSAGE AS IT IS NOW THE MOST RECENTLY DELETED MESSAGE
	var userMessage UserMessage
	result := userMessagesDb.Where("id = ?", params["id"]).Delete(&userMessage)

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

// INTERNAL FUNCTION, NOT REALLY A CRUD OPERATION
func checkRecentlyDeleted(gormID string, w http.ResponseWriter, r *http.Request) {
	log.Println("Checking If User Already Has Deleted Message (INTERNAL FUNCTION)")
	w.Header().Set("Content-Type", "application/json")
	// FIRST, CHECK IF THE SENDING USER (DERIVED FROM THE GORM ID'S MESSAGE) ALREADY HAS A STORED DELETED MESSAGE
	var user UserMessage

	result := userMessagesDb.Where("id = ?", gormID).Find(&user)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	var deletedMessage UserMessage

	result = userMessagesDb.Unscoped().Where("deleted_at IS NOT NULL AND Sender_ID = ?", user.Sender_ID).Find(&deletedMessage)

	// IF THEY DO, THEN HARD DELETE IT AS ONE USER CAN ONLY UNDO THEIR MOST RECENTLY DELETED MESSAGE
	if result.RowsAffected != 0 {
		result = userMessagesDb.Unscoped().Delete(&deletedMessage)
		if result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
			return
		}
		log.Println("Found a previously stored deleted message and permanently deleted it.")
	}
}

func undoDelete(w http.ResponseWriter, r *http.Request) {
	log.Println("Undoing a Delete (PUT)")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	var message UserMessage

	// SINCE THIS FUNCTION WILL ONLY BE AVAILABLE ONCE SOMEONE DELETES A MESSAGE, IT IS ASSUMED THAT IT WILL FIND A MESSAGE HERE
	result := userMessagesDb.Unscoped().Where("Sender_ID = ? AND Deleted_At IS NOT NULL", params["id"]).Find(&message)

	var ID = message.ID

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	// CHANGE THE MESSAGE ID'S DELETEDAT TO NULL TO BRING IT BACK
	result = userMessagesDb.Model(&UserMessage{}).Unscoped().Where("Sender_ID = ? AND Deleted_At IS NOT NULL", params["id"]).Update("Deleted_At", nil)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	if result.RowsAffected == 0 {
		http.Error(w, "Message not found.", http.StatusNotFound)
		return
	}

	// LOOK FOR THE MESSAGE AGAIN TO RETURN IT
	result = userMessagesDb.Unscoped().Where("id = ?", ID).Find(&message)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(message)
	log.Println("Message brought back successfully.")
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

	// GET THE NEXT VALID ID
	idResponse, err := getNextUserID()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// ONLY PUT IDRESPONSE IF THE ID WAS LEFT BLANK
	if userAccount.User_ID == "" {
		userAccount.User_ID = idResponse
	}

	// CHECK IF THE USER_ID IS NUMERIC AND FOUR DIGITS
	_, err = strconv.Atoi(userAccount.User_ID)

	if err != nil {
		http.Error(w, "Invalid User_ID (NOT NUMERIC): "+userAccount.User_ID+" is not a valid ID.", http.StatusBadRequest)
		return
	}

	if len(userAccount.User_ID) != 4 {
		http.Error(w, "Invalid User_ID (NOT FOUR DIGITS): "+userAccount.User_ID+" is not a valid ID.", http.StatusBadRequest)
		return
	}

	// CHECK IF THE EMAIL FITS THE EMAIL REGEX
	emailPattern := `^[a-zA-Z0-9._%+-]+@ufl.edu$`

	regex, err := regexp.Compile(emailPattern)

	if err != nil {
		http.Error(w, "Problem with compiling email regex pattern.", http.StatusBadRequest)
		return
	}

	if !regex.MatchString(userAccount.Email) {
		http.Error(w, "Invalid Email: "+userAccount.Email+" is not a valid email address. It must be an @ufl.edu email address.", http.StatusBadRequest)
		return
	}

	// SEARCHES IF THE USER EMAIL ALREADY EXISTS AND GIVES AN ERROR IF IT DOES
	var dupAccount UserAccount

	dup := userAccountsDb.Where("email = ?", userAccount.Email).First(&dupAccount)

	if dup.RowsAffected != 0 {
		http.Error(w, "Email already exists.", http.StatusBadRequest)
		return
	}

	// CHECK IF THE PHONE NUMBER MATCHES THE REGEX
	phonePattern := `^\(\d{3}\)\s\d{3}-\d{4}$`

	regex, err = regexp.Compile(phonePattern)

	if err != nil {
		http.Error(w, "Problem with compiling phone number regex pattern.", http.StatusBadRequest)
		return
	}

	if !regex.MatchString(userAccount.Phone_Number) {
		http.Error(w, "Invalid Phone Number: "+userAccount.Phone_Number+" is not a valid phone number format.", http.StatusBadRequest)
		return
	}

	// SEARCHES IF THE USER PHONE NUMBER ALREADY EXISTS AND GIVES AN ERROR IF IT DOES
	dup = userAccountsDb.Where("phone_number = ?", userAccount.Phone_Number).First(&dupAccount)

	if dup.RowsAffected != 0 {
		http.Error(w, "Phone number already exists.", http.StatusBadRequest)
		return
	}

	// HASH THE PASSWORD
	userAccount.Password = hashPassword(userAccount.Password)

	// CREATE THE USER IF IT PASSES ALL THE CHECKS
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
	log.Println("Getting a User (POST)")
	w.Header().Set("Content-Type", "application/json")

	var user UserAccount
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// HASH THE PASSWORD
	user.Password = hashPassword(user.Password)

	result := userAccountsDb.Model(&UserAccount{}).Where("email = ? AND password = ?", user.Email, user.Password).First(&user)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
	log.Println("Found user successfully.")
}

func addConversation(w http.ResponseWriter, r *http.Request) {
	log.Println("Adding an ID to a User's Conversation List (PUT)")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user UserAccount

	// FIND THE MATCHING USER STRUCT WITH THE MATCHING USER_ID
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

	// CHECK IF CONVERSATION ALREADY EXISTS
	for _, v := range conversationSlice {
		if v == params["id_2"] {
			http.Error(w, "Conversation already exists.", http.StatusBadRequest)
			return
		}
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

func deleteUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Deleting a User (DELETE)")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user UserAccount

	//FIND THE MATCHING USER STRUCT WITH THE MATCHING USER_ID AND DELETE IT
	result := userAccountsDb.Where("user_id = ?", params["id"]).Unscoped().Delete(&user)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	if result.RowsAffected == 0 {
		http.Error(w, "User not found.", http.StatusNotFound)
		return
	}
	log.Println("User deleted successfully.")
}

func getNextUserID() (string, error) {
	log.Println("Getting Next Valid ID (GET)")
	var currentID int = 1

	// GO UNTIL A VALID ID HAS BEEN FOUND
	for {
		// PAD NUMBER WITH ZEROS IF NOT FOUR DIGITS LONG
		paddedID := fmt.Sprintf("%04d", currentID)

		// CHECK IF THE ID IS ALREADY IN THE DATABASE
		var count int64

		userAccountsDb.Model(&UserAccount{}).Where("user_id = ?", paddedID).Count(&count)
		if count == 0 {
			// IF IT WAS NOT FOUND, THEN RETURN IT
			return paddedID, nil
		}

		// OTHERWISE INCREMENT ID TO THE NEXT ONE
		currentID++

		if currentID == 9996 {
			return "ERROR!", fmt.Errorf("max number of users reached")
		}
	}
}

func getUserByID(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting a User by ID (GET)")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	var user UserAccount

	result := userAccountsDb.Model(&UserAccount{}).Where("user_id = ?", params["id"]).First(&user)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
	log.Println("Found user successfully.")
}

// FUNCTION TO HASH THE PASSWORD
func hashPassword(rawPassword string) string {
	// USE SHA-256 TO HASH THE PASSWORD
	hashedPassword := sha256.Sum256([]byte(rawPassword))

	// CONVERT IT TO A HEX STRING
	encodedPassword := hex.EncodeToString(hashedPassword[:])

	return encodedPassword
}

func editFullName(w http.ResponseWriter, r *http.Request) {
	log.Println("Editing a User's Full Name (PUT)")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	var count int64
	userAccountsDb.Model(&UserAccount{}).Where("user_id = ?", params["id"]).Count(&count)
	if count == 0 {
		http.Error(w, "User not found.", http.StatusNotFound)
		return
	}

	var user UserAccount
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := userAccountsDb.Model(&UserAccount{}).Where("user_id = ?", params["id"]).Update("Full_Name", user.Full_Name)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	userAccountsDb.Model(&UserAccount{}).Where("user_id = ?", params["id"]).First(&user)

	json.NewEncoder(w).Encode(user)
	log.Println("User's full name edited successfully.")
}

func editPhoneNumber(w http.ResponseWriter, r *http.Request) {
	log.Println("Editing a User's Phone Number (PUT)")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	var count int64
	userAccountsDb.Model(&UserAccount{}).Where("user_id = ?", params["id"]).Count(&count)
	if count == 0 {
		http.Error(w, "User not found.", http.StatusNotFound)
		return
	}

	var user UserAccount
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// CHECK IF THE NEW PHONE NUMBER IS VALID
	// IF NOT, RETURN AN ERROR
	phonePattern := `^\(\d{3}\)\s\d{3}-\d{4}$`

	regex, err := regexp.Compile(phonePattern)

	if err != nil {
		http.Error(w, "Problem with compiling phone number regex pattern.", http.StatusBadRequest)
		return
	}

	if !regex.MatchString(user.Phone_Number) {
		http.Error(w, "Invalid Phone Number: "+user.Phone_Number+" is not a valid phone number format.", http.StatusBadRequest)
		return
	}

	// CHECK IF THE NEW PHONE NUMBER IS THE SAME AS THE CURRENT PHONE NUMBER
	var currentUser UserAccount
	userAccountsDb.Where("user_id = ?", params["id"]).First(&currentUser)

	if user.Phone_Number == currentUser.Phone_Number {
		json.NewEncoder(w).Encode(currentUser)
		return
	}

	// SEARCHES IF THE USER PHONE NUMBER ALREADY EXISTS AND GIVES AN ERROR IF IT DOES
	var dupAccount UserAccount
	dup := userAccountsDb.Where("phone_number = ?", user.Phone_Number).First(&dupAccount)

	if dup.RowsAffected != 0 && dupAccount.User_ID != currentUser.User_ID {
		http.Error(w, "Phone number already exists.", http.StatusBadRequest)
		return
	}

	result := userAccountsDb.Model(&UserAccount{}).Where("user_id = ?", params["id"]).Update("Phone_Number", user.Phone_Number)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	userAccountsDb.Model(&UserAccount{}).Where("user_id = ?", params["id"]).First(&user)

	json.NewEncoder(w).Encode(user)
	log.Println("User's phone number edited successfully.")
}

func searchForUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Searching for a User (POST)")
	w.Header().Set("Content-Type", "application/json")

	// THE SEARCH PATTERN FOR A USER WILL BE STORED IN THE "USERNAME" FIELD
	var user UserAccount
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// CHECK IF THE INPUT PATTERN IS VALID
	// IF NOT, RETURN AN ERROR
	searchPattern := `^[a-zA-Z\s]+#\d{4}$`

	regex, err := regexp.Compile(searchPattern)

	if err != nil {
		http.Error(w, "Problem with compiling search regex pattern.", http.StatusBadRequest)
		return
	}

	if !regex.MatchString(user.Username) {
		http.Error(w, "Invalid Search Pattern: "+user.Username+" is not a valid search pattern.", http.StatusBadRequest)
		return
	}

	// SPLIT USERNAME BY #
	usernameSplit := strings.Split(user.Username, "#")
	username := usernameSplit[0]
	userID := usernameSplit[1]

	// SEARCH FOR USER IN THE DATABASE
	result := userAccountsDb.Model(&UserAccount{}).Where("username = ? AND user_id = ?", username, userID).First(&user)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		log.Println("Unable to find a matching user.")
		return
	}

	json.NewEncoder(w).Encode(user)
	log.Println("Found user successfully.")
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

	// API FUNCTIONS FOR THE MESSAGES DATABASE

	// POST FUNCTIONS
	r.HandleFunc("/api/messages", createMessage).Methods("POST")
	r.HandleFunc("/api/messages/{id_1}/{id_2}/search", searchOneConversation).Methods("POST")
	r.HandleFunc("/api/messages/searchAll", searchAllConversations).Methods("POST")
	r.HandleFunc("/api/messages/image", createMessageImage).Methods("POST")

	// GET FUNCTIONS
	r.HandleFunc("/api/messages", getAllMessages).Methods("GET")
	r.HandleFunc("/api/messages/deleted", getDeletedMessages).Methods("GET")
	r.HandleFunc("/api/messages/{id_1}/{id_2}", getConversation).Methods("GET")
	r.HandleFunc("/api/messages/{id_1}/{id_2}/longPoll", getConversationLP).Methods("GET")

	// PUT FUNCTIONS
	r.HandleFunc("/api/messages/{id}", editMessage).Methods("PUT")
	r.HandleFunc("/api/messages/undo/{id}", undoDelete).Methods("PUT")

	// DELETE FUNCTIONS
	r.HandleFunc("/api/messages/deleteDeleted", deleteDeletedMessages).Methods("DELETE")
	r.HandleFunc("/api/messages/deleteTable", deleteTable).Methods("DELETE")
	r.HandleFunc("/api/messages/{id_1}/{id_2}", deleteConversation).Methods("DELETE")
	r.HandleFunc("/api/messages/{id}", deleteSpecificMessage).Methods("DELETE")

	// API FUNCTIONS FOR THE USERS DATABASE

	// POST FUNCTIONS
	r.HandleFunc("/api/users", createUserAccount).Methods("POST")
	r.HandleFunc("/api/users/User", getUser).Methods("POST")
	r.HandleFunc("/api/users/search", searchForUser).Methods("POST")

	// PUT FUNCTION
	r.HandleFunc("/api/users/updateN/{id}", editName).Methods("PUT")
	r.HandleFunc("/api/users/updateP/{id}", editPass).Methods("PUT")
	r.HandleFunc("/api/users/updateFN/{id}", editFullName).Methods("PUT")
	r.HandleFunc("/api/users/updatePN/{id}", editPhoneNumber).Methods("PUT")
	r.HandleFunc("/api/users/{id_1}/{id_2}", addConversation).Methods("PUT")

	// GET FUNCTIONS
	r.HandleFunc("/api/users", getAllUsers).Methods("GET")
	r.HandleFunc("/api/users/{id}", getUserByID).Methods("GET")

	//DELETE FUNCTION
	r.HandleFunc("/api/users/{id}", deleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", corsHandler(r)))
}
