package api

import (
	"encoding/json"
	"net/http"
)

type CoinBalanceParams struct{
	Username string
}

type CoinBalanceResponse struct{
	Code int
	Balance int64
}

type Error struct{
	// Error code
	Code int
	// Error message
	Message string 
}

// this function returns an error response for whos calling it

func writeError(w http.ResponseWriter, message string, statusCode int) {
	resp := Error{
		Code: statusCode,
		Message: message,
	}

	// error structure returned and type setting
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(resp)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error){
		writeError(w, err.Error(), http.StatusBadRequest)

	}
	InternalErrorHandler = func(w http.ResponseWriter){
		writeError(w, "An Unexpected Error Occurred", http.StatusInternalServerError)

	}
)