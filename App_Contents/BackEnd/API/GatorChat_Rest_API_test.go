package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// THIS TESTS THE ABILITY TO CREATE A MESSAGE IN THE MESSAGES DATABASE
func TestCreateMessage(t *testing.T) {
	// CREATE A NEW USERMESSAGE STRUCT THAT WILL BE USED TO TEST THE POST
	message := UserMessage{
		Sender_ID:   "9998",
		Receiver_ID: "9999",
		Message:     "Specific hello",

	}

	// TURN THE STRUCT INTO A JSON
	requestBody, err := json.Marshal(message)
	if err != nil {
		t.Fatalf("Failed to marshal message: %s", err)
		return
	}

	// CALL THE POST AND PASS IN THE USERMESSAGE STRUCT
	r, err := http.NewRequest("POST", "/messages", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatalf("Failed to create request: %s", err)
		return
	}

	// CREATES A RESPONSE RECORDER
	w := httptest.NewRecorder()

	// TEST THE POST BY PASSING IN THE RESPONSE RECORDER AND THE STRUCT
	createMessage(w, r)

	// VERIFY THAT THE CALL WAS SUCCESSFUL
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
		return
	}

	// TURN THE JSON BACK INTO A STRUCT
	var responseStruct UserMessage
	err = json.Unmarshal(w.Body.Bytes(), &responseStruct)
	if err != nil {
		t.Fatalf("Failed to unmarshal response body: %s", err)
		return
	}

	// CREATE THE EXPECTED USERMESSAGE STRUCT THAT SHOULD MATCH WITH THE ACTUAL STRUCT
	expectedResponse := UserMessage{
		// NOTE: THE GORM.MODEL VALUES ARE HARDCODED TO JUST BE EQUAL TO THE VALUES IN RESPONSESTRUCT
		// THIS WAS DONE BECAUSE IT IS NOT POSSIBLE TO HARDCODE THE TIME (TOO SPECIFIC)
		// THESE PARAMETERS ARE HANDLED BY GORM, SO IT CAN BE ASSUMED THAT THESE ARE ALWAYS VALID
		Model: gorm.Model{
			ID:        responseStruct.ID,
			CreatedAt: responseStruct.CreatedAt,
			UpdatedAt: responseStruct.UpdatedAt,
			DeletedAt: responseStruct.DeletedAt,
		},
		Sender_ID:   "9998",
		Receiver_ID: "9999",
		Message:     "Specific hello",
	}

	// CHECK IF THE EXPECTED RESPONSE IS EQUAL TO THE ACTUAL RESPONSE
	if !reflect.DeepEqual(responseStruct, expectedResponse) {
		t.Errorf("Expected the response body '%v', but got '%v'", expectedResponse, responseStruct)
		return
	}

	// REMOVE THAT MESSAGE ONCE TEST IS DONE
	// PASS IN THE GORM ID
	deleteTestMessage(responseStruct.ID)
}

// GENERATES TEST DATA (THIS FUNCTION IS SIMPLY CALLING A GORM COMMAND, SO IT IS ASSUMED TO ALWAYS WORK)
func createTestMessage(sender string, receiver string, mess string, image []byte) (uint, error) {
	testMessage := UserMessage{
		Sender_ID:   sender,
		Receiver_ID: receiver,
		Message:     mess,
		Image:       image,
	}

	result := userMessagesDb.Create(&testMessage)
	if result.Error != nil {
		return 0, result.Error
	}

	return testMessage.ID, nil
}

// CLEARS THE TEST DATA AFTER TEST
// PERFORMS A HARD DELETE AS WE DO NOT WANT THESE TEST MESSAGES BEING SOFT DELETED (WILL STAY AROUND)
func deleteTestMessage(messageID uint) {
	var userMessage UserMessage
	result := userMessagesDb.Where("ID = ?", messageID).Unscoped().Delete(&userMessage)
	if result.Error != nil {
		fmt.Println("Error with deleting the single test message.")
		return
	}
}

// THIS TEST RETRIEVES ALL THE MESSAGES BETWEEN TWO PEOPLE
func TestGetConversation(t *testing.T) {

	firstID, _ := createTestMessage("9998", "9999", "Specific message for TestGetConversation", []byte("test"))

	secondID, _ := createTestMessage("9999", "9998", "Specific other message for TestGetConversation", []byte("otherTest"))

	url := "/messages/" + "9998" + "/" + "9999"

	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatalf("Failed to create request: %s", err)
	}

	w := httptest.NewRecorder()

	vars := map[string]string{
		"id_1": "9998",
		"id_2": "9999",
	}

	r = mux.SetURLVars(r, vars)

	getConversation(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}

	var responseStruct []UserMessage
	err = json.Unmarshal(w.Body.Bytes(), &responseStruct)
	if err != nil {
		t.Fatalf("Failed to unmarshal response body: %s", err)
	}

	expectedResponse := []UserMessage{
		{
			Model: gorm.Model{
				ID:        firstID,
				CreatedAt: responseStruct[0].CreatedAt,
				UpdatedAt: responseStruct[0].UpdatedAt,
				DeletedAt: responseStruct[0].DeletedAt,
			},
			Sender_ID:   "9998",
			Receiver_ID: "9999",
			Message:     "Specific message for TestGetConversation",
			Image:       []byte("test"),
		},
		{
			Model: gorm.Model{
				ID:        secondID,
				CreatedAt: responseStruct[1].CreatedAt,
				UpdatedAt: responseStruct[1].UpdatedAt,
				DeletedAt: responseStruct[1].DeletedAt,
			},
			Sender_ID:   "9999",
			Receiver_ID: "9998",
			Message:     "Specific other message for TestGetConversation",
			Image:       []byte("otherTest"),
		},
	}

	if !reflect.DeepEqual(responseStruct, expectedResponse) {
		t.Errorf("Expected the response body '%v', but got '%v'", expectedResponse, responseStruct)
	}

	deleteTestMessage(firstID)
	deleteTestMessage(secondID)
}

// THIS TEST SEARCHES FOR A CREATED MESSAGE ACROSS ALL CONVERSATIONS
func TestSearchMessageAll(t *testing.T) {
	firstID, _ := createTestMessage("9998", "9999", "Specific message for TestSearchMessageAll", []byte("test"))
	secondID, _ := createTestMessage("9996", "9997", "Specific message for TestSearchMessageAll", []byte("otherTest"))

	searchMes := UserMessage{
		Message: "Specific message for TestSearchMessageAll",
	}

	requestBody, err := json.Marshal(searchMes)
	if err != nil {
		t.Fatalf("Failed to marshal message: %s", err)
	}

	r, err := http.NewRequest("POST", "/messages/searchAll", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatalf("Failed to create request: %s", err)
	}

	w := httptest.NewRecorder()

	searchAllConversations(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}

	var responseStruct []UserMessage
	err = json.Unmarshal(w.Body.Bytes(), &responseStruct)
	if err != nil {
		t.Fatalf("Failed to unmarshal response body: %s", err)
	}

	expectedResponse := []UserMessage{
		{
			Model: gorm.Model{
				ID:        firstID,
				CreatedAt: responseStruct[0].CreatedAt,
				UpdatedAt: responseStruct[0].UpdatedAt,
				DeletedAt: responseStruct[0].DeletedAt,
			},
			Sender_ID:   "9998",
			Receiver_ID: "9999",
			Message:     "Specific message for TestSearchMessageAll",
			Image:       []byte("test"),
		},
		{
			Model: gorm.Model{
				ID:        secondID,
				CreatedAt: responseStruct[1].CreatedAt,
				UpdatedAt: responseStruct[1].UpdatedAt,
				DeletedAt: responseStruct[1].DeletedAt,
			},
			Sender_ID:   "9996",
			Receiver_ID: "9997",
			Message:     "Specific message for TestSearchMessageAll",
			Image:       []byte("otherTest"),
		},
	}

	// CHECK IF THE EXPECTED RESPONSE IS EQUAL TO THE ACTUAL RESPONSE
	if !reflect.DeepEqual(responseStruct, expectedResponse) {
		t.Errorf("Expected the response body '%v', but got '%v'", expectedResponse, responseStruct)
	}

	deleteTestMessage(firstID)
	deleteTestMessage(secondID)
}

// THIS TEST SEARCHES FOR A SPECIFIC MESSAGE BETWEEN A SENDER AND USER
func TestSearchMessage(t *testing.T) {
	firstID, _ := createTestMessage("9998", "9999", "Specific message for TestSearchMessage", []byte("test"))
	secondID, _ := createTestMessage("9996", "9997", "Specific other message for TestSearchMessage", []byte("otherTest"))

	searchMes := UserMessage{
		Message: "Specific message for TestSearchMessage",
	}

	requestBody, err := json.Marshal(searchMes)
	if err != nil {
		t.Fatalf("Failed to marshal message: %s", err)
	}

	url := "/messages/" + "9998" + "/" + "9999" + "/" + "search"

	r, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatalf("Failed to create request: %s", err)
	}

	w := httptest.NewRecorder()

	vars := map[string]string{
		"id_1": "9998",
		"id_2": "9999",
	}

	r = mux.SetURLVars(r, vars)

	searchOneConversation(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}

	var responseStruct []UserMessage
	err = json.Unmarshal(w.Body.Bytes(), &responseStruct)
	if err != nil {
		t.Fatalf("Failed to unmarshal response body: %s", err)
	}

	expectedResponse := []UserMessage{
		{
			Model: gorm.Model{
				ID:        firstID,
				CreatedAt: responseStruct[0].CreatedAt,
				UpdatedAt: responseStruct[0].UpdatedAt,
				DeletedAt: responseStruct[0].DeletedAt,
			},
			Sender_ID:   "9998",
			Receiver_ID: "9999",
			Message:     "Specific message for TestSearchMessage",
			Image:       []byte("test"),
		},
	}

	// CHECK IF THE EXPECTED RESPONSE IS EQUAL TO THE ACTUAL RESPONSE
	if !reflect.DeepEqual(responseStruct, expectedResponse) {
		t.Errorf("Expected the response body '%v', but got '%v'", expectedResponse, responseStruct)
	}

	deleteTestMessage(firstID)
	deleteTestMessage(secondID)
}

// THIS TEST EDITS A CREATED MESSAGE
func TestEditMessage(t *testing.T) {
	firstID, _ := createTestMessage("9998", "9999", "This is a specific message for firstID in TestEditMessage", []byte("test"))

	newMes := UserMessage{
		Message: "Specific updated message for TestEditMessge",
	}

	requestBody, err := json.Marshal(newMes)
	if err != nil {
		t.Fatalf("Failed to marshal message: %s", err)
	}

	url := "/messages/" + fmt.Sprint(firstID)

	r, err := http.NewRequest("PUT", url, bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatalf("Failed to create request: %s", err)
	}

	w := httptest.NewRecorder()

	vars := map[string]string{
		"id": fmt.Sprint(firstID),
	}

	r = mux.SetURLVars(r, vars)

	editMessage(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}

	var responseStruct UserMessage
	err = json.Unmarshal(w.Body.Bytes(), &responseStruct)
	if err != nil {
		t.Fatalf("Failed to unmarshal response body: %s", err)
	}

	expectedResponse := UserMessage{
		Model: gorm.Model{
			ID:        firstID,
			CreatedAt: responseStruct.CreatedAt,
			UpdatedAt: responseStruct.UpdatedAt,
			DeletedAt: responseStruct.DeletedAt,
		},
		Sender_ID:   "9998",
		Receiver_ID: "9999",
		Message:     "Specific updated message for TestEditMessge",
		Image:       []byte("test"),
	}

	// CHECK IF THE EXPECTED RESPONSE IS EQUAL TO THE ACTUAL RESPONSE
	if !reflect.DeepEqual(responseStruct, expectedResponse) {
		t.Errorf("Expected the response body '%v', but got '%v'", expectedResponse, responseStruct)
	}

	deleteTestMessage(firstID)
}

// THIS TEST DELETES A CREATED MESSAGE
func TestDeleteSpecificMessage(t *testing.T) {
	firstID, _ := createTestMessage("9998", "9999", "This is a very specific message that can't possibly be accidentally replicated outside of this test", []byte("test"))

	url := "/messages/" + fmt.Sprint(firstID)
	r, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		t.Fatalf("Failed to create request: %s", err)
	}

	w := httptest.NewRecorder()

	vars := map[string]string{
		"id": fmt.Sprint(firstID),
	}

	r = mux.SetURLVars(r, vars)

	deleteSpecificMessage(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
		return
	}

	// TRY TO LOOK FOR THE MESSAGE AGAIN
	var userMessage UserMessage
	result := userMessagesDb.Where("id = ?", fmt.Sprint(firstID)).First(&userMessage)

	// THE MESSAGE SHOULD NOT BE IN THE DATABASE ANYMORE. IF IT CAN FIND IT, RETURN AN ERROR
	if result.Error == nil {
		t.Errorf("Expected message to be deleted, but it still exists.")
		return
	}

	deleteTestMessage(firstID)
}

// THIS TEST DELETES AN ENTIRE CONVERSATION BETWEEN TWO PEOPLE
func TestDeleteConversation(t *testing.T) {
	firstID, _ := createTestMessage("9998", "9999", "Specific message for firstID", []byte("test"))
	secondID, _ := createTestMessage("9999", "9998", "Specific message for secondID", []byte("otherTest"))

	url := "/messages/" + fmt.Sprint(firstID) + "/" + fmt.Sprint(secondID)
	r, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		t.Fatalf("Failed to create request: %s", err)
	}

	w := httptest.NewRecorder()

	vars := map[string]string{
		"id_1": "9998",
		"id_2": "9999",
	}

	r = mux.SetURLVars(r, vars)

	deleteConversation(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}

	// TRY TO LOOK FOR A MESSAGE
	var userMessage UserMessage
	result := userMessagesDb.Where("sender_id = ? AND receiver_id = ?", "9998", "9999").First(&userMessage)

	// THE CONVERSATION SHOULD NOT BE IN THE DATABASE ANYMORE. IF IT CAN FIND IT, RETURN AN ERROR
	if result.Error == nil {
		t.Errorf("Expected conversation to be deleted, but it still exists.")
		return
	}

	deleteTestMessage(firstID)
	deleteTestMessage(secondID)
}

// THIS TEST DELETES A MESSAGE BETWEEN TWO PEOPLE THEN UNDOES THE DELETE
func TestUndoDelete(t *testing.T) {
	// CREATE A MESSAGE
	firstID, _ := createTestMessage("9998", "9999", "Specific undo message for firstID", []byte("test"))

	// DELETE IT
	var userMessage UserMessage
	result := userMessagesDb.Where("ID = ?", firstID).Delete(&userMessage)
	if result.Error != nil {
		fmt.Println("Error with deleting the single test message.")
		return
	}

	// CALL THE UNDO FUNCTION
	url := "/messages/undo/" + "9998"
	r, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		t.Fatalf("Failed to create request: %s", err)
	}

	w := httptest.NewRecorder()

	vars := map[string]string{
		"id": "9998",
	}

	r = mux.SetURLVars(r, vars)

	//CALL UNDO
	undoDelete(w, r)

	//VERIFY IT IS RIGHT BY CHECKING THE MESSAGE'S DELETEDAT AS NULL

	// TURN THE JSON BACK INTO A STRUCT
	var responseStruct UserMessage
	err = json.Unmarshal(w.Body.Bytes(), &responseStruct)
	if err != nil {
		t.Fatalf("Failed to unmarshal response body: %s", err)
		return
	}

	expectedResponse := UserMessage{
		// NOTE: WE ARE TESTING IF DELETEDAT IS NULL NOW
		Model: gorm.Model{
			ID:        firstID,
			CreatedAt: responseStruct.CreatedAt,
			UpdatedAt: responseStruct.UpdatedAt,
			DeletedAt: gorm.DeletedAt{
				Time:  time.Time{},
				Valid: false,
			},
		},
		Sender_ID:   "9998",
		Receiver_ID: "9999",
		Message:     "Specific undo message for firstID",
		Image:       []byte("test"),
	}

	// CHECK IF THE EXPECTED RESPONSE IS EQUAL TO THE ACTUAL RESPONSE
	if !reflect.DeepEqual(responseStruct, expectedResponse) {
		t.Errorf("Expected the response body '%v', but got '%v'", expectedResponse, responseStruct)
		return
	}

	deleteTestMessage(firstID)
}

// THIS TEST CREATES A NEW USER ACCOUNT
func TestCreateUserAccount(t *testing.T) {
	user := UserAccount{
		Username:              "unitTestUser",
		Password:              "unitTestPassword",
		User_ID:               "9999",
		Email:                 "unitTest@ufl.edu",
		Full_Name:             "Test User",
		Phone_Number:          "(000) 000-0000",
		Current_Conversations: json.RawMessage([]byte("[]")),
	}

	requestBody, err := json.Marshal(user)
	if err != nil {
		t.Fatalf("Failed to marshal user: %s", err)
	}

	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatalf("Failed to create request: %s", err)
	}

	rr := httptest.NewRecorder()

	createUserAccount(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, rr.Code)
	}

	var responseStruct UserAccount
	err = json.Unmarshal(rr.Body.Bytes(), &responseStruct)
	if err != nil {
		t.Fatalf("Failed to unmarshal response body: %s", err)
	}

	expectedResponse := UserAccount{
		Username: "unitTestUser",
		// IT WILL RETURN THE SHA256 HASHED PASSWORD
		Password:              "b1b348465a1b06c150af3704f5a5f81466e77826f8351422db59b40c7a13f47e",
		User_ID:               "9999",
		Email:                 "unitTest@ufl.edu",
		Full_Name:             "Test User",
		Phone_Number:          "(000) 000-0000",
		Current_Conversations: json.RawMessage([]byte("[]")),
	}

	if !reflect.DeepEqual(responseStruct, expectedResponse) {
		t.Errorf("Expected the response body '%v', but got '%v'", expectedResponse, responseStruct)
	}

	deleteTestUser("9999")
}

// GENERATES TEST DATA (THIS FUNCTION IS SIMPLY CALLING A GORM COMMAND, SO IT IS ASSUMED TO ALWAYS WORK)
func createTestUser(username string, password string, email string, ID string, fullName string, phoneNumber string) {

	//HASH THE PASSWORD
	hashedPassword := sha256.Sum256([]byte(password))

	// CONVERT IT TO A HEX STRING
	encodedPassword := hex.EncodeToString(hashedPassword[:])

	user := UserAccount{
		Username:              username,
		Password:              encodedPassword,
		Email:                 email,
		User_ID:               ID,
		Full_Name:             fullName,
		Phone_Number:          phoneNumber,
		Current_Conversations: []byte(`[]`),
	}

	result := userAccountsDb.Create(&user)
	if result.Error != nil {
		fmt.Println("Error with creating a user.")
		return
	}
}

// CLEARS THE TEST DATA AFTER TEST
func deleteTestUser(ID string) {
	var user UserAccount
	result := userAccountsDb.Where("user_id = ?", ID).Unscoped().Delete(&user)
	if result.Error != nil {
		fmt.Println("Error with deleting the user.")
		return
	}
}

// THIS TEST DELETES A CONTACT FROM A USER'S LIST OF CURRENT CONVERSATIONS
func TestDeleteContact(t *testing.T) {
	createTestUser("unitTestUser", "unitTestPass", "unitTest@ufl.edu", "9999", "Test User", "(000) 000-0000")

	r, err := http.NewRequest("GET", "/api/9999/0000", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %s", err)
	}

	w := httptest.NewRecorder()

	vars := map[string]string{
		"id_1": "9999",
		"id_2": "0000",
	}

	r = mux.SetURLVars(r, vars)

	addConversation(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}

	r, err = http.NewRequest("DELETE", "/api/users/removeC/9999/0000", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %s", err)
	}

	w = httptest.NewRecorder()

	vars = map[string]string{
		"id_1": "9999",
		"id_2": "0000",
	}

	r = mux.SetURLVars(r, vars)

	deleteContact(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}

	var responseStruct UserAccount
	err = json.Unmarshal(w.Body.Bytes(), &responseStruct)
	if err != nil {
		t.Fatalf("Failed to unmarshal response body: %s", err)
	}

	expectedResponse := UserAccount{
		Username:              "unitTestUser",
		Password:              "f3632dec6bc0cead273d4301a8f13cb89e7ee0ef95175fd2c2ed7a7b6c0dac73",
		Email:                 "unitTest@ufl.edu",
		User_ID:               "9999",
		Full_Name:             "Test User",
		Phone_Number:          "(000) 000-0000",
		Current_Conversations: []byte(`[]`),
	}

	if !reflect.DeepEqual(responseStruct, expectedResponse) {
		t.Errorf("Expected the response body '%v', but got '%v'", expectedResponse, responseStruct)
	}

	deleteTestUser("9999")
}

// THIS TEST ADDS A NEW CONVERSATION
func TestAddConversation(t *testing.T) {
	createTestUser("unitTestUser", "unitTestPass", "unitTest@ufl.edu", "9999", "Test User", "(000) 000-0000")

	// THIS TEST IS ADDING ITS OWN ID TO ITSELF SINCE ADDING A DIFFERENT ONE WILL TRIGGER THE REVERSE FUNCTION (WHICH IS NOT BEING ACCOUNTED FOR HERE)
	// THIS IS SPECIFICALLY TESTING THE FUNCTIONALITY OF ADDING AN ID SO THIS IMPLEMENTATION WILL STILL PROVE THAT IT WORKS
	r, err := http.NewRequest("GET", "/api/9999/9999", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %s", err)
	}

	w := httptest.NewRecorder()

	vars := map[string]string{
		"id_1": "9999",
		"id_2": "9999",
	}

	r = mux.SetURLVars(r, vars)

	addConversation(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}

	var responseStruct UserAccount
	err = json.Unmarshal(w.Body.Bytes(), &responseStruct)
	if err != nil {
		t.Fatalf("Failed to unmarshal response body: %s", err)
	}

	expectedResponse := UserAccount{
		Username:              "unitTestUser",
		Password:              "f3632dec6bc0cead273d4301a8f13cb89e7ee0ef95175fd2c2ed7a7b6c0dac73",
		Email:                 "unitTest@ufl.edu",
		User_ID:               "9999",
		Full_Name:             "Test User",
		Phone_Number:          "(000) 000-0000",
		Current_Conversations: []byte(`["9999"]`),
	}

	// CHECK IF THE EXPECTED RESPONSE IS EQUAL TO THE ACTUAL RESPONSE
	if !reflect.DeepEqual(responseStruct, expectedResponse) {
		t.Errorf("Expected the response body '%v', but got '%v'", expectedResponse, responseStruct)
	}

	deleteTestUser("9999")
}

// THIS TEST EDITS THE USERNAME
func TestEditName(t *testing.T) {
	createTestUser("unitTestUser", "unitTestPass", "unitTest@ufl.edu", "9999", "Test User", "(000) 000-0000")

	newName := UserAccount{
		Username:              "uuunitTestUuuser",
	}

	requestBody, err := json.Marshal(newName)
	if err != nil {
		t.Fatalf("Failed to marshal message: %s", err)
	}

	r, err := http.NewRequest("PUT", "/api/users/updateN/9999", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatalf("Failed to create request: %s", err)
	}

	w := httptest.NewRecorder()

	vars := map[string]string{
		"id": "9999",
	}

	r = mux.SetURLVars(r, vars)

	editName(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}

	var responseStruct UserAccount
	err = json.Unmarshal(w.Body.Bytes(), &responseStruct)
	if err != nil {
		t.Fatalf("Failed to unmarshal response body: %s", err)
	}

	expectedResponse := UserAccount{
		Username:              "uuunitTestUuuser",
		Password:              "f3632dec6bc0cead273d4301a8f13cb89e7ee0ef95175fd2c2ed7a7b6c0dac73",
		User_ID:               "9999",
		Email:                 "unitTest@ufl.edu",
		Full_Name:             "Test User",
		Phone_Number:          "(000) 000-0000",
		Current_Conversations: json.RawMessage([]byte("[]")),
	}

	if !reflect.DeepEqual(responseStruct, expectedResponse) {
		t.Errorf("Expected the response body '%v', but got '%v'", expectedResponse, responseStruct)
	}

	deleteTestUser("9999")
}

// THIS TEST EDITS THE PASSWORD
func TestEditPass(t *testing.T) {
	createTestUser("unitTestUser", "unitTestPass", "unitTest@ufl.edu", "9999", "Test User", "(000) 000-0000")

	newPass := UserAccountConfirmPass{
		UserAccount: UserAccount{
			Password: "newTestPass",
		},
		OriginalPassword: "unitTestPass",
	}

	requestBody, err := json.Marshal(newPass)
	if err != nil {
		t.Fatalf("Failed to marshal message: %s", err)
	}

	r, err := http.NewRequest("PUT", "/api/users/updateP/9999", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatalf("Failed to create request: %s", err)
	}

	w := httptest.NewRecorder()

	vars := map[string]string{
		"id": "9999",
	}

	r = mux.SetURLVars(r, vars)

	editPass(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}

	var responseStruct UserAccount
	err = json.Unmarshal(w.Body.Bytes(), &responseStruct)
	if err != nil {
		t.Fatalf("Failed to unmarshal response body: %s", err)
	}

	expectedResponse := UserAccount{
		Username:              "unitTestUser",
		Password:              "c8eef775bc0e26d0fd2479eb35fdf0e568e6fb7ad36abd9b58198a1be248fe99",
		User_ID:               "9999",
		Email:                 "unitTest@ufl.edu",
		Full_Name:             "Test User",
		Phone_Number:          "(000) 000-0000",
		Current_Conversations: json.RawMessage([]byte("[]")),
	}

	if !reflect.DeepEqual(responseStruct, expectedResponse) {
		t.Errorf("Expected the response body '%v', but got '%v'", expectedResponse, responseStruct)
	}

	deleteTestUser("9999")
}

// THIS TEST RETURNS A USER (BASED ON EMAIL AND PASSWORD)
func TestGetUser(t *testing.T) {
	createTestUser("unitTestUser", "unitTestPass", "unitTest@ufl.edu", "9999", "Test User", "(000) 000-0000")

	user := UserAccount{
		Email:    "unitTest@ufl.edu",
		Password: "unitTestPass",
	}

	requestBody, err := json.Marshal(user)
	if err != nil {
		t.Fatalf("Failed to marshal message: %s", err)
	}

	r, err := http.NewRequest("POST", "/users/User", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatalf("Failed to create request: %s", err)
	}

	w := httptest.NewRecorder()

	getUser(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}

	var responseStruct UserAccount
	err = json.Unmarshal(w.Body.Bytes(), &responseStruct)
	if err != nil {
		t.Fatalf("Failed to unmarshal response body: %s", err)
	}

	expectedResponse := UserAccount{
		Username:              "unitTestUser",
		Password:              "f3632dec6bc0cead273d4301a8f13cb89e7ee0ef95175fd2c2ed7a7b6c0dac73",
		Email:                 "unitTest@ufl.edu",
		User_ID:               "9999",
		Full_Name:             "Test User",
		Phone_Number:          "(000) 000-0000",
		Current_Conversations: []byte(`[]`),
	}

	// CHECK IF THE EXPECTED RESPONSE IS EQUAL TO THE ACTUAL RESPONSE
	if !reflect.DeepEqual(responseStruct, expectedResponse) {
		t.Errorf("Expected the response body '%v', but got '%v'", expectedResponse, responseStruct)
	}

	deleteTestUser("9999")
}

// THIS TEST RETURNS A USER (BASED ON ID)
func TestGetUserByID(t *testing.T) {
	createTestUser("unitTestUser", "unitTestPass", "unitTest@ufl.edu", "9999", "Test User", "(000) 000-0000")

	r, err := http.NewRequest("GET", "/users/9999", bytes.NewBuffer(nil))
	if err != nil {
		t.Fatalf("Failed to create request: %s", err)
	}

	w := httptest.NewRecorder()

	vars := map[string]string{
		"id": "9999",
	}

	r = mux.SetURLVars(r, vars)

	getUserByID(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}

	var responseStruct UserAccount
	err = json.Unmarshal(w.Body.Bytes(), &responseStruct)
	if err != nil {
		t.Fatalf("Failed to unmarshal response body: %s", err)
	}

	expectedResponse := UserAccount{
		Username:              "unitTestUser",
		Password:              "f3632dec6bc0cead273d4301a8f13cb89e7ee0ef95175fd2c2ed7a7b6c0dac73",
		Email:                 "unitTest@ufl.edu",
		User_ID:               "9999",
		Full_Name:             "Test User",
		Phone_Number:          "(000) 000-0000",
		Current_Conversations: []byte(`[]`),
	}

	// CHECK IF THE EXPECTED RESPONSE IS EQUAL TO THE ACTUAL RESPONSE
	if !reflect.DeepEqual(responseStruct, expectedResponse) {
		t.Errorf("Expected the response body '%v', but got '%v'", expectedResponse, responseStruct)
	}

	deleteTestUser("9999")
}

// THIS TEST DELETES A USER
func TestDeleteUser(t *testing.T) {
	createTestUser("unitTestUser", "unitTestPass", "unitTest@ufl.edu", "9999", "Test User", "(000) 000-0000")

	r, err := http.NewRequest("DELETE", "/users/9999", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %s", err)
	}

	w := httptest.NewRecorder()

	vars := map[string]string{
		"id": "9999",
	}

	r = mux.SetURLVars(r, vars)

	deleteUser(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
		return
	}

	// TRY TO LOOK FOR THE USER AGAIN
	var user UserAccount
	result := userAccountsDb.Where("user_id = ?", "9999").First(&user)

	// THE USER SHOULD NOT BE IN THE DATABASE ANYMORE. IF IT CAN FIND IT, RETURN AN ERROR
	if result.Error == nil {
		t.Errorf("Expected user to be deleted, but it still exists.")
		return
	}
}

func TestEditFullName(t *testing.T) {
	createTestUser("unitTestUser", "unitTestPass", "unitTest@ufl.edu", "9999", "Test User", "(000) 000-0000")

	user := UserAccount{
		Full_Name: "New Name",
	}

	requestBody, err := json.Marshal(user)
	if err != nil {
		t.Fatalf("Failed to marshal message: %s", err)
	}

	r, err := http.NewRequest("PUT", "/users/updateFN/9999", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatalf("Failed to create request: %s", err)
	}

	w := httptest.NewRecorder()

	vars := map[string]string{
		"id": "9999",
	}

	r = mux.SetURLVars(r, vars)

	editFullName(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}

	var responseStruct UserAccount
	err = json.Unmarshal(w.Body.Bytes(), &responseStruct)
	if err != nil {
		t.Fatalf("Failed to unmarshal response body: %s", err)
	}

	expectedResponse := UserAccount{
		Username:              "unitTestUser",
		Password:              "f3632dec6bc0cead273d4301a8f13cb89e7ee0ef95175fd2c2ed7a7b6c0dac73",
		Email:                 "unitTest@ufl.edu",
		User_ID:               "9999",
		Full_Name:             "New Name",
		Phone_Number:          "(000) 000-0000",
		Current_Conversations: []byte(`[]`),
	}

	// CHECK IF THE EXPECTED RESPONSE IS EQUAL TO THE ACTUAL RESPONSE
	if !reflect.DeepEqual(responseStruct, expectedResponse) {
		t.Errorf("Expected the response body '%v', but got '%v'", expectedResponse, responseStruct)
	}

	deleteTestUser("9999")
}

func TestEditPhoneNumber(t *testing.T) {
	createTestUser("unitTestUser", "unitTestPass", "unitTest@ufl.edu", "9999", "Test User", "(000) 000-0000")

	user := UserAccount{
		Phone_Number: "(000) 000-0001",
	}

	requestBody, err := json.Marshal(user)
	if err != nil {
		t.Fatalf("Failed to marshal message: %s", err)
	}

	r, err := http.NewRequest("PUT", "/users/updatePN/9999", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatalf("Failed to create request: %s", err)
	}

	w := httptest.NewRecorder()

	vars := map[string]string{
		"id": "9999",
	}

	r = mux.SetURLVars(r, vars)

	editPhoneNumber(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}

	var responseStruct UserAccount
	err = json.Unmarshal(w.Body.Bytes(), &responseStruct)
	if err != nil {
		t.Fatalf("Failed to unmarshal response body: %s", err)
	}

	expectedResponse := UserAccount{
		Username:              "unitTestUser",
		Password:              "f3632dec6bc0cead273d4301a8f13cb89e7ee0ef95175fd2c2ed7a7b6c0dac73",
		Email:                 "unitTest@ufl.edu",
		User_ID:               "9999",
		Full_Name:             "Test User",
		Phone_Number:          "(000) 000-0001",
		Current_Conversations: []byte(`[]`),
	}

	// CHECK IF THE EXPECTED RESPONSE IS EQUAL TO THE ACTUAL RESPONSE
	if !reflect.DeepEqual(responseStruct, expectedResponse) {
		t.Errorf("Expected the response body '%v', but got '%v'", expectedResponse, responseStruct)
	}

	deleteTestUser("9999")
}

func TestSearchUser(t *testing.T) {
	createTestUser("unitTestUser", "unitTestPass", "unitTest@ufl.edu", "9999", "Test User", "(000) 000-0000")

	user := UserAccount{
		Username: "unitTestUser#9999",
	}

	requestBody, err := json.Marshal(user)
	if err != nil {
		t.Fatalf("Failed to marshal message: %s", err)
	}

	r, err := http.NewRequest("POST", "/users/search", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatalf("Failed to create request: %s", err)
	}

	w := httptest.NewRecorder()

	searchForUser(w, r)

	var responseStruct UserAccount
	err = json.Unmarshal(w.Body.Bytes(), &responseStruct)
	if err != nil {
		t.Fatalf("Failed to unmarshal response body: %s", err)
	}

	expectedResponse := UserAccount{
		Username:              "unitTestUser",
		Password:              "f3632dec6bc0cead273d4301a8f13cb89e7ee0ef95175fd2c2ed7a7b6c0dac73",
		Email:                 "unitTest@ufl.edu",
		User_ID:               "9999",
		Full_Name:             "Test User",
		Phone_Number:          "(000) 000-0000",
		Current_Conversations: []byte(`[]`),
	}

	// CHECK IF THE EXPECTED RESPONSE IS EQUAL TO THE ACTUAL RESPONSE
	if !reflect.DeepEqual(responseStruct, expectedResponse) {
		t.Errorf("Expected the response body '%v', but got '%v'", expectedResponse, responseStruct)
	}

	deleteTestUser("9999")
}

func TestGetRecentConvo(t *testing.T) {
	// SIMULATE A MESSAGE THAT WAS SENT FROM USER 9999 TO USER 9998
	firstID, _ := createTestMessage("9999", "9998", "Specific message for TestGetRecentConvo", []byte("test"))

	// THIS IS USER 9998
	createTestUser("unitTestUser", "unitTestPass", "unitTest@ufl.edu", "9998", "Test User", "(000) 000-0000")

	// SEARCH FOR THE LAST PERSON 9999 TALKED TO (9998)
	r, err := http.NewRequest("GET", "/messages/getRecent/user/9999", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %s", err)
	}

	w := httptest.NewRecorder()

	vars := map[string]string{
		"id": "9999",
	}

	r = mux.SetURLVars(r, vars)

	getMostRecentConvo(w, r)

	var responseStruct UserAccount
	err = json.Unmarshal(w.Body.Bytes(), &responseStruct)
	if err != nil {
		t.Fatalf("Failed to unmarshal response body: %s", err)
	}

	expectedResponse := UserAccount{
		Username:              "unitTestUser",
		Password:              "f3632dec6bc0cead273d4301a8f13cb89e7ee0ef95175fd2c2ed7a7b6c0dac73",
		Email:                 "unitTest@ufl.edu",
		User_ID:               "9998",
		Full_Name:             "Test User",
		Phone_Number:          "(000) 000-0000",
		Current_Conversations: []byte(`[]`),
	}

	// CHECK IF THE EXPECTED RESPONSE IS EQUAL TO THE ACTUAL RESPONSE
	if !reflect.DeepEqual(responseStruct, expectedResponse) {
		t.Errorf("Expected the response body '%v', but got '%v'", expectedResponse, responseStruct)
	}

	deleteTestMessage(firstID)
	deleteTestUser("9998")
}
