package handler

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"
	"todo-list/todo"
)

// @Summary      Create a user
// @Description  Registers a new user
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input  body      todo.User           true  "User data"
// @Success      200    {object}  SuccessResponse
// @Failure      400    {object}  ErrorResponse
// @Failure      500    {object}  ErrorResponse
// @Router       /auth/sign-up [post]
func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	if !checkContentType(w, r) {
		return
	}

	var input todo.User

	if err := parseJSONBody(w, r, &input); err != nil {
		return
	}

	err := validateUser(input)
	if handleError(w, err, http.StatusBadRequest, "failed validation user") {
		return
	}

	id, err := h.services.Authorization.AuthUser(input)
	if handleError(w, err, http.StatusInternalServerError, "User creation failed due to internal server error") {
		return
	}

	newSuccessResponse(w, "User created successfully", map[string]interface{}{"userID": id})
}

func validateUser(user todo.User) error {
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		errorMessages := handleValidationError(err)
		return fmt.Errorf("validation failed: %s", errorMessages)
	}

	return nil
}

func handleValidationError(err error) string {
	var errorMessages []string
	for _, e := range err.(validator.ValidationErrors) {
		errorMessages = append(errorMessages, fmt.Sprintf("Field '%s' failed validation for '%s' tag", e.Field(), e.Tag()))
	}

	return strings.Join(errorMessages, ", ")
}

// @Summary      SignIn
// @Tags         auth
// @Description  login
// @ID           login
// @Accept       json
// @Produce      json
// @Param        input  body      todo.User           true  "credentials"
// @Success      200    {object}  SuccessResponse
// @Failure      400,404 {object}  ErrorResponse
// @Failure      500    {object}  ErrorResponse
// @Router       /auth/sign-in [post]
func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	if !checkContentType(w, r) {
		return
	}

	var input todo.User

	if err := parseJSONBody(w, r, &input); err != nil {
		return
	}

	err := validateUser(input)
	if handleError(w, err, http.StatusBadRequest, "failed validation user") {
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Email, input.Password)
	if handleError(w, err, http.StatusBadRequest, "cant create token") {
		return
	}

	newSuccessResponse(w, "Token created successfully", map[string]interface{}{"token": token})
}
