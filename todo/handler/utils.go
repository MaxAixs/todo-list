package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strconv"
)

func checkContentType(w http.ResponseWriter, r *http.Request) bool {
	if r.Header.Get("content-type") != "application/json" {
		newErrorResponse(w, "invalid content-type", http.StatusUnsupportedMediaType, "Content-Type must be application json")
		return false
	}

	return true
}

func getUserID(ctx context.Context) (uuid.UUID, error) {
	userID, ok := ctx.Value("userID").(uuid.UUID)
	if !ok || userID == uuid.Nil {
		return uuid.Nil, errors.New("user ID not found in context")
	}

	return userID, nil
}

func getListID(r *http.Request) (int, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok || id == "" {
		return 0, errors.New("list ID not provided in URL")
	}

	listID, err := strconv.Atoi(id)
	if err != nil {
		return 0, errors.New("invalid ID")
	}

	return listID, nil
}

func getItemID(r *http.Request) (int, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok || id == "" {
		return 0, errors.New("item ID not provided in URL")
	}

	itemID, err := strconv.Atoi(id)
	if err != nil {
		return 0, errors.New("invalid item ID")
	}

	return itemID, nil
}

func parseJSONBody(w http.ResponseWriter, r *http.Request, input interface{}) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(input); err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError

		switch {
		case errors.As(err, &syntaxError):
			newErrorResponse(w, err.Error(), http.StatusBadRequest, fmt.Sprintf("Invalid JSON at positon %d", syntaxError.Offset))

		case errors.As(err, &unmarshalTypeError):
			newErrorResponse(w, err.Error(), http.StatusBadRequest, fmt.Sprintf("Invalid value for field %d", unmarshalTypeError.Offset))

		case errors.Is(err, io.EOF):
			newErrorResponse(w, err.Error(), http.StatusBadRequest, "Request body cannot be empty")

		default:
			newErrorResponse(w, err.Error(), http.StatusBadRequest, "Invalid JSON")
		}

		return err
	}

	return nil
}
