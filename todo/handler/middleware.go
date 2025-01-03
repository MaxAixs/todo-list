package handler

import (
	"context"
	"errors"
	"net/http"
	"strings"
)

func (h *Handler) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")

		token, err := getTokenFromHeader(header)
		if err != nil {
			newErrorResponse(w, err.Error(), http.StatusUnauthorized, "cant get token from header")
			return
		}

		userID, err := h.services.ParseToken(token)
		{
			if err != nil {
				newErrorResponse(w, err.Error(), http.StatusUnauthorized, "")
				return
			}
		}

		ctx := context.WithValue(r.Context(), "userID", userID)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func getTokenFromHeader(header string) (string, error) {
	if header == "" {
		return "", errors.New("authorization token is missing")
	}

	headerParts := strings.Split(header, " ")

	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return "", errors.New("authorization token is invalid")
	}

	if len(headerParts[1]) == 0 {
		return "", errors.New("token is missing")
	}

	return headerParts[1], nil

}
