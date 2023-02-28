package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func TestCreateMessage(t *testing.T) {
	// CREATE A NEW USERMESSAGE STRUCT THAT WILL BE USED TO TEST THE POST
	message := UserMessage{
		Sender_ID:   "1234",
		Receiver_ID: "5678",
		Message:     "test",
	}

	// TURN THE STRUCT INTO A JSON
	requestBody, err := json.Marshal(message)
	if err != nil {
		t.Fatalf("Failed to marshal message: %s", err)
	}

	// CALL THE POST AND PASS IN THE USERMESSAGE STRUCT
	req, err := http.NewRequest("POST", "/messages", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatalf("Failed to create request: %s", err)
	}

	// CREATES A RESPONSE RECORDER
	rr := httptest.NewRecorder()

	// TEST THE POST BY PASSING IN THE RESPONSE RECORDER AND THE STRUCT
	createMessage(rr, req)

	// VERIFY THAT THE CALL WAS SUCCESSFUL
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, rr.Code)
	}

	// TURN THE JSON BACK INTO A STRUCT
	var responseStruct UserMessage
	err = json.Unmarshal(rr.Body.Bytes(), &responseStruct)
	if err != nil {
		t.Fatalf("Failed to unmarshal response body: %s", err)
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
		Sender_ID:   "1234",
		Receiver_ID: "5678",
		Message:     "test",
	}

	// CHECK IF THE EXPECTED RESPONSE IS EQUAL TO THE ACTUAL RESPONSE
	if !reflect.DeepEqual(responseStruct, expectedResponse) {
		t.Errorf("Expected the response body '%v', but got '%v'", expectedResponse, responseStruct)
	}
}

func TestGetConversation(t *testing.T) {
	r, err := http.NewRequest("GET", "/messages/1234/5678", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %s", err)
	}

	w := httptest.NewRecorder()

	vars := map[string]string{
		"id_1": "1234",
		"id_2": "5678",
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
			Sender_ID:   "1234",
			Receiver_ID: "5678",
			Message:     "test",
		},
	}

	if !reflect.DeepEqual(responseStruct, expectedResponse) {
		t.Errorf("Expected the response body '%v', but got '%v'", expectedResponse, responseStruct)
	}
}

func TestSearchMessage(t *testing.T) {
	r, err := http.NewRequest("GET", "/messages/test", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %s", err)
	}

	w := httptest.NewRecorder()

	vars := map[string]string{
		"search": "test",
	}

	r = mux.SetURLVars(r, vars)

	searchMessage(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}

	// this is what the request returns
	var responseStruct []UserMessage
	err = json.Unmarshal(w.Body.Bytes(), &responseStruct)
	if err != nil {
		t.Fatalf("Failed to unmarshal response body: %s", err)
	}

	//we need to make a matching datatype, in this case a slice of usermessage structs

	expectedResponse := []UserMessage{
		{
			Model: gorm.Model{
				ID:        responseStruct[0].ID,
				CreatedAt: responseStruct[0].CreatedAt,
				UpdatedAt: responseStruct[0].UpdatedAt,
				DeletedAt: responseStruct[0].DeletedAt,
			},
			Sender_ID:   "1234",
			Receiver_ID: "5678",
			Message:     "test",
		},
	}

	// CHECK IF THE EXPECTED RESPONSE IS EQUAL TO THE ACTUAL RESPONSE
	if !reflect.DeepEqual(responseStruct, expectedResponse) {
		t.Errorf("Expected the response body '%v', but got '%v'", expectedResponse, responseStruct)
	}
}

func TestEditMessage(t *testing.T) {
	newMes := UserMessage{
		Message: "Hi",
	}

	requestBody, err := json.Marshal(newMes)
	if err != nil {
		t.Fatalf("Failed to marshal message: %s", err)
	}

	r, err := http.NewRequest("PUT", "/messages/1234/5678/test", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatalf("Failed to create request: %s", err)
	}

	w := httptest.NewRecorder()

	vars := map[string]string{
		"id_1":         "1234",
		"id_2":         "5678",
		"inputMessage": "test",
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
		Sender_ID:   "1234",
		Receiver_ID: "5678",
		Message:     "Hi",
	}

	// CHECK IF THE EXPECTED RESPONSE IS EQUAL TO THE ACTUAL RESPONSE
	if !reflect.DeepEqual(responseStruct, expectedResponse) {
		t.Errorf("Expected the response body '%v', but got '%v'", expectedResponse, responseStruct)
	}
}

func TestDeleteMessage(t *testing.T) {
	delMes := UserMessage{}

	requestBody, err := json.Marshal(delMes)
	if err != nil {
		t.Fatalf("Failed to marshal message: %s", err)
	}

	r, err := http.NewRequest("DELETE", "/messages/1234/5678", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatalf("Failed to create request: %s", err)
	}

	w := httptest.NewRecorder()

	vars := map[string]string{
		"id_1": "1234",
		"id_2": "5678",
	}

	r = mux.SetURLVars(r, vars)

	deleteMessage(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}

	var responseStruct UserMessage

	expectedResponse := UserMessage{
		Model: gorm.Model{
			ID:        responseStruct.ID,
			CreatedAt: responseStruct.CreatedAt,
			UpdatedAt: responseStruct.UpdatedAt,
			DeletedAt: responseStruct.DeletedAt,
		},
		Sender_ID:   "",
		Receiver_ID: "",
		Message:     "",
	}

	// CHECK IF THE EXPECTED RESPONSE IS EQUAL TO THE ACTUAL RESPONSE
	if !reflect.DeepEqual(responseStruct, expectedResponse) {
		t.Errorf("Expected the response body '%v', but got '%v'", expectedResponse, responseStruct)
	}
}

func TestDeleteSpecificMessage(t *testing.T) {
	delMes := UserMessage{}

	requestBody, err := json.Marshal(delMes)
	if err != nil {
		t.Fatalf("Failed to marshal message: %s", err)
	}

	r, err := http.NewRequest("DELETE", "/messages/1234/5678/Hi", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatalf("Failed to create request: %s", err)
	}

	w := httptest.NewRecorder()

	vars := map[string]string{
		"id_1":         "1234",
		"id_2":         "5678",
		"inputMessage": "Hi",
	}

	r = mux.SetURLVars(r, vars)

	deleteSpecificMessage(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}

	var responseStruct UserMessage

	expectedResponse := UserMessage{
		Model: gorm.Model{
			ID:        responseStruct.ID,
			CreatedAt: responseStruct.CreatedAt,
			UpdatedAt: responseStruct.UpdatedAt,
			DeletedAt: responseStruct.DeletedAt,
		},
		Sender_ID:   "",
		Receiver_ID: "",
		Message:     "",
	}

	// CHECK IF THE EXPECTED RESPONSE IS EQUAL TO THE ACTUAL RESPONSE
	if !reflect.DeepEqual(responseStruct, expectedResponse) {
		t.Errorf("Expected the response body '%v', but got '%v'", expectedResponse, responseStruct)
	}
}

func TestCreateUserAccount(t *testing.T) {
	user := UserAccount{
		Username: "user",
		Password: "pass",
		User_ID:  "1234",
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
		Username: "user",
		Password: "pass",
		User_ID:  "1234",
	}

	if !reflect.DeepEqual(responseStruct, expectedResponse) {
		t.Errorf("Expected the response body '%v', but got '%v'", expectedResponse, responseStruct)
	}
}

func TestGetAllUsers(t *testing.T) {

}

func TestAddConversation(t *testing.T) {

}

func TestGetUser(t *testing.T) {

}
