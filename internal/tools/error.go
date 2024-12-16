package tools

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Code    int
	Message string
}

func writeError(w http.ResponseWriter, message string, code int) {
	resp := Error{
		Code:    code,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)
}

var (
	UnAuthorizedHandler = func(w http.ResponseWriter) {
		writeError(w, "Unauthorized", http.StatusUnauthorized)
	}
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHander = func(w http.ResponseWriter) {
		writeError(w, "An Expected Error Occured", http.StatusInternalServerError)
	}
)
