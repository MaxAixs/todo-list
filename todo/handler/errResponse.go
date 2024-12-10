package handler

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
)

type ErrorResponse struct {
	Error      string `json:"error"`
	Msg        string `json:"msg"`
	StatusCode int    `json:"status_code"`
}

func newErrorResponse(w http.ResponseWriter, err string, statusCode int, msg string) {
	logrus.Errorf("Error: %s, Messaage: %s", err, msg)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := ErrorResponse{
		Error:      err,
		Msg:        msg,
		StatusCode: statusCode}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		logrus.Println("failed to write JSON response:", err)
	}
}
