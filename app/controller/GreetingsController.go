package controller

import (
	"encoding/json"
	"net/http"
)

// Greetings ..
func Greetings(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Hello World")
}
