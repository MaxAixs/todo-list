package handler

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

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
	err := json.NewDecoder(r.Body).Decode(input)
	if err != nil {
		handleError(w, err, http.StatusBadRequest, "invalid input body")
		return err
	}
	return nil
}
