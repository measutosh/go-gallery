package main

import (
	"encoding/json"
	"net/http"
)

// defining the JSON representation
type jsonResponse struct {
	Error bool `json:"error"`
	Message string `json:"message"`
	Data any `json:"data,omitempty"`
}


// defining one handler which will be declared in routes.go
func(app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	// declare a variable of jsonREsponse and populate it
	payload := jsonResponse {
		Error: false,
		Message: "Hi 😊, you hit the Broker!",
        Data: nil,
	}

	// print out the json
	out, _ := json.MarshalIndent(payload, "", "\t") 
	// Header returns the header map that will be sent by WriteHeader
	w.Header().Set("Content-Type", "application/json")
	// WriteHeader sends an HTTP response header with the provided status code.
	w.WriteHeader(http.StatusAccepted)
	// Write writes the data to the connection as part of an HTTP reply.
	w.Write(out)
	// adding this in routes.go as mux.Broker
}