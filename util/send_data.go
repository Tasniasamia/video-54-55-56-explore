package util;

import (
	"encoding/json"
	"net/http"
)

func SendResponse(w http.ResponseWriter, data interface{}, status int) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

