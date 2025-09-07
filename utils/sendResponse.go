package utils

import (
	"encoding/json"
	"net/http"
)

func SendResponse(w http.ResponseWriter, content interface{}, statusCode int){
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(content)
}