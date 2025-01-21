package handler

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
	Msg   string `json:"msg"`
}

func newErrorResponse(w http.ResponseWriter, err string, statusCode int, msg string) {
	logrus.Errorf("error: %s, Messaage: %s", err, msg)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := ErrorResponse{
		Error: err,
		Msg:   msg,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		logrus.Printf("failed to write JSON response: %v", err)
	}
}

func handleError(w http.ResponseWriter, err error, statusCode int, message string) bool {
	if err != nil {
		newErrorResponse(w, err.Error(), statusCode, message)
		return true
	}
	return false
}
