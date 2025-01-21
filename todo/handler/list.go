package handler

import (
	"net/http"
	"todo-list/todo"
)

// @Summary      Create a new list
// @Description  Creates a new todo list
// @Tags         lists
// @Accept       json
// @Produce      json
// @Param        input  body      todo.TodoList  true  "Todo List"
// @Success      200    {object}  SuccessResponse
// @Failure      400    {object}  ErrorResponse
// @Failure      401    {object}  ErrorResponse
// @Failure      500    {object}  ErrorResponse
// @Security     ApiKeyAuth
// @Router       /api/lists/ [post]
func (h *Handler) CreateList(w http.ResponseWriter, r *http.Request) {
	if !checkContentType(w, r) {
		return
	}

	id, err := getUserID(r.Context())
	if handleError(w, err, http.StatusUnauthorized, "cant get user id") {
		return
	}

	var input todo.TodoList
	if err := parseJSONBody(w, r, &input); err != nil {
		return
	}

	listID, err := h.services.TodoList.CreateList(id, input)
	if handleError(w, err, http.StatusBadRequest, "failed to create") {
		return
	}

	newSuccessResponse(w, "list create successfully", map[string]interface{}{"listID": listID})
}

// @Summary      Get a list by ID
// @Description  Retrieves a todo list by its ID
// @Tags         lists
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "List ID"
// @Success      200  {object}  SuccessResponse
// @Failure      400  {object}  ErrorResponse
// @Failure      401  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Security     ApiKeyAuth
// @Router       /api/lists/{id} [get]
func (h *Handler) GetList(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserID(r.Context())
	if handleError(w, err, http.StatusUnauthorized, "cant get user id") {
		return
	}

	listID, err := getListID(r)
	if handleError(w, err, http.StatusBadRequest, "list ID is empty") {
		return
	}

	list, err := h.services.TodoList.GetListById(userID, listID)
	if handleError(w, err, http.StatusInternalServerError, "failed to get list") {
		return
	}

	newSuccessResponse(w, "list successfully received", map[string]interface{}{"list": list})
}

// @Summary      Update a list
// @Description  Updates a todo list by its ID
// @Tags         lists
// @Accept       json
// @Produce      json
// @Param        id     path      int                  true  "List ID"
// @Param        input  body      todo.UpdateListInput true  "Update List Data"
// @Success      200    {object}  SuccessResponse
// @Failure      400    {object}  ErrorResponse
// @Failure      401    {object}  ErrorResponse
// @Failure      500    {object}  ErrorResponse
// @Security     ApiKeyAuth
// @Router       /api/lists/{id} [put]
func (h *Handler) UpdateList(w http.ResponseWriter, r *http.Request) {
	if !checkContentType(w, r) {
		return
	}

	userID, err := getUserID(r.Context())
	if handleError(w, err, http.StatusUnauthorized, "cant get user id") {
		return
	}

	listID, err := getListID(r)
	if handleError(w, err, http.StatusBadRequest, "list ID is empty") {
		return
	}

	var input todo.UpdateListInput
	if err := parseJSONBody(w, r, &input); err != nil {
		return
	}

	err = h.services.TodoList.UpdateList(userID, listID, input)
	if handleError(w, err, http.StatusInternalServerError, "failed to update") {
		return
	}

	newSuccessResponse(w, "list update successfully", map[string]interface{}{"Status": "OK"})
}

// @Summary      Delete a list
// @Description  Deletes a todo list by its ID
// @Tags         lists
// @Accept       json
// @Produce      json
// @Param        listID   path      int  true  "List ID"
// @Success      200      {object}  SuccessResponse
// @Failure      400      {object}  ErrorResponse
// @Failure      401      {object}  ErrorResponse
// @Failure      500      {object}  ErrorResponse
// @Security     ApiKeyAuth
// @Router       /api/lists/{id} [delete]
func (h *Handler) DeleteList(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserID(r.Context())
	if handleError(w, err, http.StatusUnauthorized, "cant get user id") {
		return
	}

	listID, err := getListID(r)
	if handleError(w, err, http.StatusBadRequest, "list ID is empty") {
		return
	}

	err = h.services.TodoList.DeleteListById(userID, listID)
	if handleError(w, err, http.StatusBadRequest, "failed to delete list") {
		return
	}

	newSuccessResponse(w, "list delete successfully", map[string]interface{}{"status": "OK"})
}

// @Summary      Get all lists
// @Description  Retrieves all todo lists for the current user
// @Tags         lists
// @Accept       json
// @Produce      json
// @Success      200      {object}  SuccessResponse
// @Failure      400      {object}  ErrorResponse
// @Failure      401      {object}  ErrorResponse
// @Failure      500      {object}  ErrorResponse
// @Security     ApiKeyAuth
// @Router       /api/lists/ [get]
func (h *Handler) GetAllLists(w http.ResponseWriter, r *http.Request) {
	id, err := getUserID(r.Context())
	if handleError(w, err, http.StatusUnauthorized, "cant get user id") {
		return
	}

	lists, err := h.services.TodoList.GetAllLists(id)
	if handleError(w, err, http.StatusBadRequest, "failed to get all lists") {
		return
	}

	newSuccessResponse(w, "all list successfully received", map[string]interface{}{"AllLists": lists})
}
