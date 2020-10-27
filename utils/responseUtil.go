package utils

import (
	"encoding/json"
	"net/http"
)

// Message ... Responds to client with a single message
func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

// Respond ... Responds to client with the json represantation of data
func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
