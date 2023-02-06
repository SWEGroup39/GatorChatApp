package main

// this code is interfaced via Postman

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Text Struct
type Text struct {
	ID      string  `json:"id:"`
	Sender  *Sender `json:"sender"`
	Message string  `json:"message"`
}

// Sender Struct
type Sender struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Init texts var as a slice Text struct
var texts []Text

// Get All Texts
func getTexts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(texts)
}

// Create a New Text
func createText(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var text Text
	_ = json.NewDecoder(r.Body).Decode(&text)
	text.ID = strconv.Itoa(rand.Intn(10000000))
	texts = append(texts, text)
	json.NewEncoder(w).Encode(text)
}

// Edit a Text
func editText(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range texts {
		if item.ID == params["id"] {
			texts = append(texts[:index], texts[index+1:]...)
			var text Text
			_ = json.NewDecoder(r.Body).Decode(&text)
			text.ID = params["id"]
			texts = append(texts, text)
			json.NewEncoder(w).Encode(text)
			return
		}
	}
}

// Delete a Message
func deleteText(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range texts {
		if item.ID == params["id"] {
			texts = append(texts[:index], texts[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(texts)
}

func main() {
	// Init Router
	r := mux.NewRouter()

	// Mock Text Data for testing purposes
	texts = append(texts, Text{ID: "1", Sender: &Sender{Firstname: "Johnny", Lastname: "Struckman"}, Message: "Hello"})
	texts = append(texts, Text{ID: "2", Sender: &Sender{Firstname: "Strohnny", Lastname: "Juckman"}, Message: "Hi"})

	// Text Route Handlers / Endpoints
	r.HandleFunc("/api/texts", getTexts).Methods("GET")
	r.HandleFunc("/api/texts", createText).Methods("POST")
	r.HandleFunc("/api/texts/{id}", editText).Methods("PUT")
	r.HandleFunc("/api/texts/{id}", deleteText).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
