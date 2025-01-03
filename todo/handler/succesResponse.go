package handler

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
)

type SuccessResponse struct {
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func newSuccessResponse(w http.ResponseWriter, msg string, data interface{}) {
	logrus.Printf("status %d, msg %s ", http.StatusOK, msg)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := SuccessResponse{
		Msg:  msg,
		Data: data,
	}

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		logrus.Printf("failed to write JSON response: %v", err)
	}

}
