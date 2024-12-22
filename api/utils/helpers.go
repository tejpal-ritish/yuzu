package utils

import (
	"encoding/json"
	"net/http"
)

func WriteJSONResponse(w http.ResponseWriter, status int, data interface{}) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
