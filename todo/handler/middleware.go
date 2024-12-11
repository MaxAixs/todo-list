package handler

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"net/http"
	"strings"
)

func (h *Handler) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if header == "" {
			newErrorResponse(w, "empty auth header", http.StatusUnauthorized, "Authorization token is missing")
			return
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			newErrorResponse(w, "invalid auth header", http.StatusUnauthorized, "Authorization token is invalid")
			return
		}

		if len(headerParts[1]) == 0 {
			newErrorResponse(w, "token is empty", http.StatusUnauthorized, "Authorization token is missing")
		}

		userID, err := h.services.ParseToken(headerParts[1])
		{
			if err != nil {
				newErrorResponse(w, err.Error(), http.StatusUnauthorized, "")
			}
		}

		ctx := context.WithValue(r.Context(), "userID", userID)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func GetUserID(ctx context.Context) (uuid.UUID, error) {
	userID, ok := ctx.Value("userID").(uuid.UUID)
	if !ok || userID == uuid.Nil {
		return uuid.Nil, errors.New("user ID not found in context")
	}

	return userID, nil
}
