package handler

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"todo-list/todo/service"
	mock_service "todo-list/todo/service/mocks"
)

func Test_AuthMiddleware(t *testing.T) {
	type mockBehavior func(s *mock_service.MockAuthorization, token string)

	tests := []struct {
		name                 string
		headerName           string
		headerValue          string
		token                string
		mock                 mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:        "OK",
			headerName:  "Authorization",
			headerValue: "Bearer token",
			token:       "token",
			mock: func(s *mock_service.MockAuthorization, token string) {
				s.EXPECT().ParseToken(token).Return(uuid.New(), nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"msg":"success"}`,
		},
		{
			name:        "Invalid token",
			headerName:  "Authorization",
			headerValue: "Bearer token",
			token:       "token",
			mock: func(s *mock_service.MockAuthorization, token string) {
				s.EXPECT().ParseToken(token).Return(uuid.Nil, errors.New("invalid token"))
			},
			expectedStatusCode:   401,
			expectedResponseBody: `{"error":"invalid token","msg":""}`,
		},
		{
			name:                 "Invalid Header Name",
			headerName:           "",
			headerValue:          "Bearer token",
			token:                "token",
			mock:                 func(s *mock_service.MockAuthorization, token string) {},
			expectedStatusCode:   401,
			expectedResponseBody: `{"error":"authorization token is missing","msg":"cant get token from header"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			auth := mock_service.NewMockAuthorization(ctrl)
			tt.mock(auth, tt.token)

			services := &service.Service{Authorization: auth}
			handler := Handler{services}

			finalHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"msg":"success"}`))
			})

			middleware := handler.AuthMiddleware(finalHandler)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/middleware", nil)
			req.Header.Set(tt.headerName, tt.headerValue)

			middleware.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
			assert.Equal(t, tt.expectedResponseBody, strings.TrimSpace(w.Body.String()))
		})
	}
}
