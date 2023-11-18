package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type JsonRequest struct {
	Id      int       `json:"_id,omitempty"`
	Name    string    `json:"name,omitempty"`
	Message string    `json:"message"`
	Date    time.Time `json:"date"`
}

func (c *Controllers) DataController(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	// Unmarshal the JSON data into a JsonRequest struct
	var requestData JsonRequest
	requestData = JsonRequest{
		Date: time.Now(),
	}
	err = json.Unmarshal(body, &requestData)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	// Print the received JSON data
	fmt.Printf("Received JSON data: %+v\n", requestData)

	// Marshal the JsonRequest struct back to JSON
	responseJSON, err := json.Marshal(requestData)
	if err != nil {
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response back to the client
	w.Write(responseJSON)

}
