package handler

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"todo-list/todo"
)

// @Summary      Create a new item
// @Description  Creates a new todo item in a specific list
// @Tags         items
// @Accept       json
// @Produce      json
// @Param        listID   path      int                 true  "List ID"
// @Param        input    body      todo.TodoItem       true  "Todo Item"
// @Success      200      {object}  SuccessResponse
// @Failure      400      {object}  ErrorResponse
// @Failure      401      {object}  ErrorResponse
// @Failure      500      {object}  ErrorResponse
// @Security     ApiKeyAuth
// @Router       /api/lists/{listID}/items [post]
func (h *Handler) CreateItem(w http.ResponseWriter, r *http.Request) {
	if !checkContentType(w, r) {
		return
	}

	userID, err := getUserID(r.Context())
	if handleError(w, err, http.StatusUnauthorized, "cant get user id") {
		return
	}

	listID, err := getListID(r)
	if handleError(w, err, http.StatusBadRequest, "invalid list id param") {
		return
	}

	var input todo.TodoItem
	if err := parseJSONBody(w, r, &input); err != nil {
		return
	}

	logrus.Printf("UserID: %v", userID)
	itemID, err := h.services.TodoItem.CreateItem(userID, listID, input)
	if handleError(w, err, http.StatusInternalServerError, "create item failed") {
		return
	}

	newSuccessResponse(w, "create item successfully", map[string]interface{}{"itemID": itemID})
}

// @Summary      Delete an item
// @Description  Deletes a todo item by its ID
// @Tags         items
// @Accept       json
// @Produce      json
// @Param        itemID   path      int  true  "Item ID"
// @Success      200      {object}  SuccessResponse
// @Failure      400      {object}  ErrorResponse
// @Failure      401      {object}  ErrorResponse
// @Failure      500      {object}  ErrorResponse
// @Security     ApiKeyAuth
// @Router       /api/items/{itemID} [delete]
func (h *Handler) DeleteItem(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserID(r.Context())
	if handleError(w, err, http.StatusUnauthorized, "cant get user id") {
		return
	}

	itemID, err := getItemID(r)
	if handleError(w, err, http.StatusBadRequest, "invalid item id param") {
		return
	}

	err = h.services.TodoItem.DeleteItemById(userID, itemID)
	if handleError(w, err, http.StatusInternalServerError, "cant delete item") {
		return
	}

	newSuccessResponse(w, "delete item successfully", map[string]interface{}{"Status": "OK"})
}

// @Summary      Get all items
// @Description  Retrieves all todo items in a specific list
// @Tags         items
// @Accept       json
// @Produce      json
// @Param        listID   path      int  true  "List ID"
// @Success      200      {object}  SuccessResponse
// @Failure      400      {object}  ErrorResponse
// @Failure      401      {object}  ErrorResponse
// @Failure      500      {object}  ErrorResponse
// @Security     ApiKeyAuth
// @Router       /api/lists/{listID}/items [get]
func (h *Handler) GetItems(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserID(r.Context())
	if handleError(w, err, http.StatusUnauthorized, "cant get user id") {
		return
	}

	listID, err := getListID(r)
	if handleError(w, err, http.StatusBadRequest, "invalid list id param") {
		return
	}

	items, err := h.services.GetAllItems(userID, listID)
	if handleError(w, err, http.StatusInternalServerError, "cant get all items") {
		return
	}

	newSuccessResponse(w, "get all items successfully", map[string]interface{}{"Items": items})
}

// @Summary      Get item by ID
// @Description  Retrieves a todo item by its ID
// @Tags         items
// @Accept       json
// @Produce      json
// @Param        itemID   path      int  true  "Item ID"
// @Success      200      {object}  SuccessResponse
// @Failure      400      {object}  ErrorResponse
// @Failure      401      {object}  ErrorResponse
// @Failure      500      {object}  ErrorResponse
// @Security     ApiKeyAuth
// @Router       /api/items/{itemID} [get]
func (h *Handler) GetItemById(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserID(r.Context())
	if handleError(w, err, http.StatusUnauthorized, "cant get user id") {
		return
	}

	itemID, err := getItemID(r)
	if handleError(w, err, http.StatusBadRequest, "invalid list id params") {
		return
	}

	item, err := h.services.TodoItem.GetItemById(userID, itemID)
	if handleError(w, err, http.StatusInternalServerError, "cant get item by id") {
		return
	}

	newSuccessResponse(w, "item get successfully", map[string]interface{}{"Item": item})
}

// @Summary      Update an item
// @Description  Updates a todo item by its ID
// @Tags         items
// @Accept       json
// @Produce      json
// @Param        itemID   path      int                   true  "Item ID"
// @Param        input    body      todo.UpdateItemInput  true  "Update Item Data"
// @Success      200      {object}  SuccessResponse
// @Failure      400      {object}  ErrorResponse
// @Failure      401      {object}  ErrorResponse
// @Failure      500      {object}  ErrorResponse
// @Security     ApiKeyAuth
// @Router       /api/items/{itemID} [put]
func (h *Handler) UpdateItem(w http.ResponseWriter, r *http.Request) {
	if !checkContentType(w, r) {
		return
	}

	userID, err := getUserID(r.Context())
	if handleError(w, err, http.StatusUnauthorized, "cant get user id") {
		return
	}

	itemID, err := getItemID(r)
	if handleError(w, err, http.StatusBadRequest, "invalid list id param") {
		return
	}

	var input todo.UpdateItemInput
	if err := parseJSONBody(w, r, &input); err != nil {
		return
	}

	err = h.services.TodoItem.UpdateItem(userID, itemID, input)
	if handleError(w, err, http.StatusInternalServerError, "cant update item") {
		return
	}

	newSuccessResponse(w, "update item successfully", map[string]interface{}{"Status": "OK"})
}
