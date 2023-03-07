package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// THIS TESTS THE ABILITY TO CREATE A MESSAGE IN THE MESSAGES DATABASE
func TestCreateMessage(t *testing.T) {
	// CREATE A NEW USERMESSAGE STRUCT THAT WILL BE USED TO TEST THE POST
	message := UserMessage{
		Sender_ID:   "0001",
		Receiver_ID: "0002",
		Message:     "Hello",
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
		// ALSO, THE ID IS ALWAYS ALTERING WHICH MAKES HARDCODING THE ACTUAL ID DIFFICULT
		// THESE PARAMETERS ARE HANDLED BY GORM, SO IT CAN BE ASSUMED THAT THESE ARE ALWAYS VALID
		Model: gorm.Model{
			ID:        responseStruct.ID,
			CreatedAt: responseStruct.CreatedAt,
			UpdatedAt: responseStruct.UpdatedAt,
			DeletedAt: responseStruct.DeletedAt,
		},
		Sender_ID:   "0001",
		Receiver_ID: "0002",
		Message:     "Hello",
	}

	// CHECK IF THE EXPECTED RESPONSE IS EQUAL TO THE ACTUAL RESPONSE
	if !reflect.DeepEqual(responseStruct, expectedResponse) {
		t.Errorf("Expected the response body '%v', but got '%v'", expectedResponse, responseStruct)
		return
	}

	// REMOVE THAT MESSAGE ONCE TEST IS DONE
	deleteTestMessage("0001", "0002", "Hello")
}

// GENERATES TEST DATA (THIS FUNCTION IS SIMPLY CALLING A GORM COMMAND, SO IT IS ASSUMED TO ALWAYS WORK)
func createTestMessage(sender string, receiver string, mess string) {
	testMessage := UserMessage{
		Sender_ID:   sender,
		Receiver_ID: receiver,
		Message:     mess,
	}

	result := userMessagesDb.Create(&testMessage)
	if result.Error != nil {
		fmt.Println("Error with creating single test message.")
		return
	}
}

// CLEARS THE TEST DATA AFTER TEST
func deleteTestMessage(sender string, receiver string, mess string) {
	var userMessage UserMessage
	result := userMessagesDb.Where("sender_id = ? AND receiver_id = ? AND message = ?", sender, receiver, mess).Unscoped().Delete(&userMessage)
	if result.Error != nil {
		fmt.Println("Error with deleting the single test message.")
		return
	}
}

// THIS TEST RETRIEVES ALL THE MESSAGES BETWEEN TWO PEOPLE
func TestGetConversation(t *testing.T) {

	createTestMessage("0001", "0002", "Testing")
	createTestMessage("0002", "0001", "Testing_2")

	r, err := http.NewRequest("GET", "/messages/0001/0002", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %s", err)
	}

	w := httptest.NewRecorder()

	vars := map[string]string{
		"id_1": "0001",
		"id_2": "0002",
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
				ID:        responseStruct[0].ID,
				CreatedAt: responseStruct[0].CreatedAt,
				UpdatedAt: responseStruct[0].UpdatedAt,
				DeletedAt: responseStruct[0].DeletedAt,
			},
			Sender_ID:   "0001",
			Receiver_ID: "0002",
			Message:     "Testing",
		},
		{
			Model: gorm.Model{
				ID:        responseStruct[1].ID,
				CreatedAt: responseStruct[1].CreatedAt,
				UpdatedAt: responseStruct[1].UpdatedAt,
				DeletedAt: responseStruct[1].DeletedAt,
			},
			Sender_ID:   "0002",
			Receiver_ID: "0001",
			Message:     "Testing_2",
		},
	}

	if !reflect.DeepEqual(responseStruct, expectedResponse) {
		t.Errorf("Expected the response body '%v', but got '%v'", expectedResponse, responseStruct)
	}

	deleteTestMessage("0001", "0002", "Testing")
	deleteTestMessage("0002", "0001", "Testing_2")
}

// THIS TEST SEARCHES FOR A CREATED MESSAGE ACROSS ALL CONVERSATIONS
func TestSearchMessageAll(t *testing.T) {
	createTestMessage("0001", "0002", "Testing")
	createTestMessage("0003", "0004", "Testing")

	searchMes := UserMessage{
		Message: "Testing",
	}

	requestBody, err := json.Marshal(searchMes)
	if err != nil {
		t.Fatalf("Failed to marshal message: %s", err)
	}

	r, err := http.NewRequest("GET", "/messages/searchAll", bytes.NewBuffer(requestBody))
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
				ID:        responseStruct[0].ID,
				CreatedAt: responseStruct[0].CreatedAt,
				UpdatedAt: responseStruct[0].UpdatedAt,
				DeletedAt: responseStruct[0].DeletedAt,
			},
			Sender_ID:   "0001",
			Receiver_ID: "0002",
			Message:     "Testing",
		},
		{
			Model: gorm.Model{
				ID:        responseStruct[1].ID,
				CreatedAt: responseStruct[1].CreatedAt,
				UpdatedAt: responseStruct[1].UpdatedAt,
				DeletedAt: responseStruct[1].DeletedAt,
			},
			Sender_ID:   "0003",
			Receiver_ID: "0004",
			Message:     "Testing",
		},
	}

	// CHECK IF THE EXPECTED RESPONSE IS EQUAL TO THE ACTUAL RESPONSE
	if !reflect.DeepEqual(responseStruct, expectedResponse) {
		t.Errorf("Expected the response body '%v', but got '%v'", expectedResponse, responseStruct)
	}

	deleteTestMessage("0001", "0002", "Testing")
	deleteTestMessage("0003", "0004", "Testing")
}

func TestSearchMessage(t *testing.T) {
	createTestMessage("0001", "0002", "Testing")
	createTestMessage("0003", "0004", "Testing")

	searchMes := UserMessage{
		Message: "Testing",
	}

	requestBody, err := json.Marshal(searchMes)
	if err != nil {
		t.Fatalf("Failed to marshal message: %s", err)
	}

	r, err := http.NewRequest("GET", "/messages/0001/0002/search", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatalf("Failed to create request: %s", err)
	}

	w := httptest.NewRecorder()

	vars := map[string]string{
		"id_1":         "0001",
		"id_2":         "0002",
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
				ID:        responseStruct[0].ID,
				CreatedAt: responseStruct[0].CreatedAt,
				UpdatedAt: responseStruct[0].UpdatedAt,
				DeletedAt: responseStruct[0].DeletedAt,
			},
			Sender_ID:   "0001",
			Receiver_ID: "0002",
			Message:     "Testing",
		},
	}

	// CHECK IF THE EXPECTED RESPONSE IS EQUAL TO THE ACTUAL RESPONSE
	if !reflect.DeepEqual(responseStruct, expectedResponse) {
		t.Errorf("Expected the response body '%v', but got '%v'", expectedResponse, responseStruct)
	}

	deleteTestMessage("0001", "0002", "Testing")
	deleteTestMessage("0003", "0004", "Testing")
}

// THIS TEST EDITS A CREATED MESSAGE
func TestEditMessage(t *testing.T) {
	createTestMessage("0001", "0002", "Testing")

	newMes := UserMessage{
		Message: "Update",
	}

	requestBody, err := json.Marshal(newMes)
	if err != nil {
		t.Fatalf("Failed to marshal message: %s", err)
	}

	r, err := http.NewRequest("PUT", "/messages/0001/0002/Testing", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatalf("Failed to create request: %s", err)
	}

	w := httptest.NewRecorder()

	vars := map[string]string{
		"id_1":         "0001",
		"id_2":         "0002",
		"inputMessage": "Testing",
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
			ID:        responseStruct.ID,
			CreatedAt: responseStruct.CreatedAt,
			UpdatedAt: responseStruct.UpdatedAt,
			DeletedAt: responseStruct.DeletedAt,
		},
		Sender_ID:   "0001",
		Receiver_ID: "0002",
		Message:     "Update",
	}

	// CHECK IF THE EXPECTED RESPONSE IS EQUAL TO THE ACTUAL RESPONSE
	if !reflect.DeepEqual(responseStruct, expectedResponse) {
		t.Errorf("Expected the response body '%v', but got '%v'", expectedResponse, responseStruct)
	}

	deleteTestMessage("0001", "0002", "Update")
}

// THIS TEST DELETES A CREATED MESSAGE
func TestDeleteSpecificMessage(t *testing.T) {
	createTestMessage("0001", "0002", "Testing")

	r, err := http.NewRequest("DELETE", "/messages/0001/0002/Testing", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %s", err)
	}

	w := httptest.NewRecorder()

	vars := map[string]string{
		"id_1":         "0001",
		"id_2":         "0002",
		"inputMessage": "Testing",
	}

	r = mux.SetURLVars(r, vars)

	deleteSpecificMessage(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
		return
	}

	// TRY TO LOOK FOR THE MESSAGE AGAIN
	var userMessage UserMessage
	result := userMessagesDb.Where("sender_id = ? AND receiver_id = ? AND message = ?", "0001", "0002", "Testing").First(&userMessage)

	// THE MESSAGE SHOULD NOT BE IN THE DATABASE ANYMORE. IF IT CAN FIND IT, RETURN AN ERROR
	if result.Error == nil {
		t.Errorf("Expected message to be deleted, but it still exists.")
		return
	}
}

// THIS TEST DELETES AN ENTIRE CONVERSATION BETWEEN TWO PEOPLE
func TestDeleteConversation(t *testing.T) {
	createTestMessage("0001", "0002", "Testing")
	createTestMessage("0002", "0001", "Testing_2")

	r, err := http.NewRequest("DELETE", "/messages/0001/0002", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %s", err)
	}

	w := httptest.NewRecorder()

	vars := map[string]string{
		"id_1": "0001",
		"id_2": "0002",
	}

	r = mux.SetURLVars(r, vars)

	deleteConversation(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}

	// TRY TO LOOK FOR A MESSAGE
	var userMessage UserMessage
	result := userMessagesDb.Where("sender_id = ? AND receiver_id = ?", "0001", "0002").First(&userMessage)

	// THE CONVERSATION SHOULD NOT BE IN THE DATABASE ANYMORE. IF IT CAN FIND IT, RETURN AN ERROR
	if result.Error == nil {
		t.Errorf("Expected conversation to be deleted, but it still exists.")
		return
	}
}

// THIS TEST CREATES A NEW USER ACCOUNT
func TestCreateUserAccount(t *testing.T) {
	user := UserAccount{
		Username:              "user",
		Password:              "pass",
		User_ID:               "0001",
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
		Username:              "user",
		Password:              "pass",
		User_ID:               "0001",
		Current_Conversations: json.RawMessage([]byte("[]")),
	}

	if !reflect.DeepEqual(responseStruct, expectedResponse) {
		t.Errorf("Expected the response body '%v', but got '%v'", expectedResponse, responseStruct)
	}

	deleteTestUser("0001")
}

// GENERATES TEST DATA (THIS FUNCTION IS SIMPLY CALLING A GORM COMMAND, SO IT IS ASSUMED TO ALWAYS WORK)
func createTestUser(username string, password string, ID string) {
	user := UserAccount{
		Username:              username,
		Password:              password,
		User_ID:               ID,
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

// THIS TEST ADDS A NEW CONVERSATION
func TestAddConversation(t *testing.T) {
	createTestUser("user", "pass", "0001")

	r, err := http.NewRequest("PUT", "/api/0001/0002", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %s", err)
	}

	w := httptest.NewRecorder()

	vars := map[string]string{
		"id_1": "0001",
		"id_2": "0002",
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
		Username:              "user",
		Password:              "pass",
		User_ID:               "0001",
		Current_Conversations: []byte(`["0002"]`),
	}

	// CHECK IF THE EXPECTED RESPONSE IS EQUAL TO THE ACTUAL RESPONSE
	if !reflect.DeepEqual(responseStruct, expectedResponse) {
		t.Errorf("Expected the response body '%v', but got '%v'", expectedResponse, responseStruct)
	}

	deleteTestUser("0001")
}

// THIS TEST RETURNS A USER
func TestGetUser(t *testing.T) {
	createTestUser("user", "pass", "0001")

	user := UserAccount{
		Username: "user",
		Password: "pass",
	}

	requestBody, err := json.Marshal(user)
	if err != nil {
		t.Fatalf("Failed to marshal message: %s", err)
	}

	r, err := http.NewRequest("GET", "/users/User", bytes.NewBuffer(requestBody))
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
		Username:              "user",
		Password:              "pass",
		User_ID:               "0001",
		Current_Conversations: []byte(`[]`),
	}

	// CHECK IF THE EXPECTED RESPONSE IS EQUAL TO THE ACTUAL RESPONSE
	if !reflect.DeepEqual(responseStruct, expectedResponse) {
		t.Errorf("Expected the response body '%v', but got '%v'", expectedResponse, responseStruct)
	}

	deleteTestUser("0001")
}

// THIS TEST DELETES A USER
func TestDeleteUser(t *testing.T) {
	createTestUser("user", "pass", "0001")

	r, err := http.NewRequest("DELETE", "/users/0001", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %s", err)
	}

	w := httptest.NewRecorder()

	vars := map[string]string{
		"id": "0001",
	}

	r = mux.SetURLVars(r, vars)

	deleteUser(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
		return
	}

	// TRY TO LOOK FOR THE USER AGAIN
	var user UserAccount
	result := userAccountsDb.Where("user = ?", "0001").First(&user)

	// THE USER SHOULD NOT BE IN THE DATABASE ANYMORE. IF IT CAN FIND IT, RETURN AN ERROR
	if result.Error == nil {
		t.Errorf("Expected message to be deleted, but it still exists.")
		return
	}
}
