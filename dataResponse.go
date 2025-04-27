package main

import (
	"encoding/json"
	"net/http"
)

func dataResponse(w http.ResponseWriter, _ *http.Request, data ...any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Failed to serialize data", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}
