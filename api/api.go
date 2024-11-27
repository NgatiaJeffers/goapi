package api

import (
	"encoding/json",
	"net/http"
)

/* STRUCT */
// Coin Balance Params
type CoinBalanceParams struct {
	Username string
}

// Coin Balance Response
type CoinBalanceResponse struct {
	// Success Code, usually 200
	Code int

	// Account Balance
	Balance int64
}

// Error Response
type Error struct {
	// Error code
	Code int

	// Error Message
	Message string
}

func writeError(w http.ResponseWriter, message string, code int) {
	resp := Error{
		Code:    code,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHEader(code)

	json.NewEncoder(w).Encode(resp)
}

// Wraapper
var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ReponseWriter) {
		writeError(w, "An Unexpected Error Occurred.", http.StatusInternalServerError)
	}
)