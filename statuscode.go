package statuscode

import (
	"encoding/json"
	"net/http"
)

func Return400BadRequest(w http.ResponseWriter, message string) {
	response := make(map[string]interface{})
	response["message"] = message
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(response)
}

func Return401Unauthorized(w http.ResponseWriter, message string) {
	response := make(map[string]interface{})
	response["message"] = message
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(response)
}

func Return403Forbidden(w http.ResponseWriter, message string) {
	response := make(map[string]interface{})
	response["message"] = message
	w.WriteHeader(http.StatusForbidden)
	json.NewEncoder(w).Encode(response)
}

func Return404NotFound(w http.ResponseWriter, message string) {
	response := make(map[string]interface{})
	response["message"] = message
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(response)
}

func Return409Conflict(w http.ResponseWriter, message string) {
	response := make(map[string]interface{})
	response["message"] = message
	w.WriteHeader(http.StatusConflict)
	json.NewEncoder(w).Encode(response)
}

func Return422UnprocessableEntity(w http.ResponseWriter, message string) {
	response := make(map[string]interface{})
	response["message"] = message
	w.WriteHeader(http.StatusUnprocessableEntity)
	json.NewEncoder(w).Encode(response)
}

func Return429TooManyRequests(w http.ResponseWriter, message string) {
	response := make(map[string]interface{})
	response["message"] = message
	w.WriteHeader(http.StatusTooManyRequests)
	json.NewEncoder(w).Encode(response)
}
