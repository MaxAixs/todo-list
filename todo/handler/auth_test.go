package handler

import (
	"bytes"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"strings"
	"testing"
	"todo-list/todo"
	"todo-list/todo/service"
	mockservice "todo-list/todo/service/mocks"
)

func TestHandler_SignUp(t *testing.T) {
	type mockBehavior func(s *mockservice.MockAuthorization, user todo.User)
	expectUUID := uuid.New()

	tests := []struct {
		name                 string
		mock                 mockBehavior
		inputBody            string
		inputUser            todo.User
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "Successful SignUp",
			mock: func(s *mockservice.MockAuthorization, user todo.User) {
				s.EXPECT().AuthUser(user).Return(expectUUID, nil)
			},
			inputBody: `{"name":"testName","password":"qwerty123","email":"test@test.com"}`,
			inputUser: todo.User{
				Name:     "testName",
				Email:    "test@test.com",
				Password: "qwerty123",
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"msg":"User created successfully","data":{"userID":"` + expectUUID.String() + `"}}`,
		},
		{
			name:                 "Validation Error",
			mock:                 func(s *mockservice.MockAuthorization, user todo.User) {},
			inputBody:            `{"name":"testName","password":"qwerty123"}`,
			inputUser:            todo.User{},
			expectedStatusCode:   400,
			expectedResponseBody: `{"error":"validation failed: Field 'Email' failed validation for 'required' tag","msg":"failed validation user"}`,
		},
		{
			name:                 "Empty Request Body",
			mock:                 func(s *mockservice.MockAuthorization, user todo.User) {},
			inputBody:            ``,
			inputUser:            todo.User{},
			expectedStatusCode:   400,
			expectedResponseBody: `{"error":"EOF","msg":"Request body cannot be empty"}`,
		},
		{
			name: "Server error",
			mock: func(s *mockservice.MockAuthorization, user todo.User) {
				s.EXPECT().AuthUser(user).Return(uuid.Nil, fmt.Errorf("cant create user"))
			},
			inputBody: `{"name":"testName","password":"qwerty123","email":"test@test.com"}`,
			inputUser: todo.User{
				Name:     "testName",
				Email:    "test@test.com",
				Password: "qwerty123",
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"error":"cant create user","msg":"User creation failed due to internal server error"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			auth := mockservice.NewMockAuthorization(ctrl)
			tt.mock(auth, tt.inputUser)

			services := &service.Service{Authorization: auth}
			handler := Handler{services}

			r := mux.NewRouter()
			r.HandleFunc("/auth/sign-up", handler.signUp).Methods("POST")

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/auth/sign-up", bytes.NewBufferString(tt.inputBody))
			req.Header.Set("content-type", "application/json")

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, tt.expectedStatusCode)
			assert.Equal(t, strings.TrimSpace(w.Body.String()), tt.expectedResponseBody)
		})
	}
}

func TestHandler_SignIn(t *testing.T) {
	type mockBehavior func(s *mockservice.MockAuthorization, user todo.User)

	tests := []struct {
		name                 string
		mock                 mockBehavior
		inputBody            string
		inputUser            todo.User
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "Successful SignIn",
			mock: func(s *mockservice.MockAuthorization, user todo.User) {
				s.EXPECT().GenerateToken(user.Email, user.Password).Return("jwt token", nil)
			},
			inputBody: `{"name":"testName","email":"test@test.com","password":"qwerty123"}`,
			inputUser: todo.User{
				Name:     "testName",
				Email:    "test@test.com",
				Password: "qwerty123",
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"msg":"Token created successfully","data":{"token":"jwt token"}}`,
		}, {
			name:                 "Empty Request Body",
			mock:                 func(s *mockservice.MockAuthorization, user todo.User) {},
			inputBody:            ``,
			inputUser:            todo.User{},
			expectedStatusCode:   400,
			expectedResponseBody: `{"error":"EOF","msg":"Request body cannot be empty"}`,
		},
		{
			name: "Failed Validation Error",
			mock: func(s *mockservice.MockAuthorization, user todo.User) {
			},
			inputBody:            `{"name": "testName", "password": "qwerty123"}`,
			inputUser:            todo.User{},
			expectedStatusCode:   400,
			expectedResponseBody: `{"error":"validation failed: Field 'Email' failed validation for 'required' tag","msg":"failed validation user"}`,
		},
		{
			name: "Wrong email or password",
			mock: func(s *mockservice.MockAuthorization, user todo.User) {
				s.EXPECT().GenerateToken(user.Email, user.Password).Return("", fmt.Errorf("cant get user: sql: no rows in result set"))
			},
			inputBody: `{"name": "testName", "email":"testEmail@test.com","password":"qwerty123"}`,
			inputUser: todo.User{
				Name:     "testName",
				Email:    "testEmail@test.com",
				Password: "qwerty123",
			},
			expectedStatusCode:   400,
			expectedResponseBody: `{"error":"cant get user: sql: no rows in result set","msg":"cant create token"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			auth := mockservice.NewMockAuthorization(ctrl)
			tt.mock(auth, tt.inputUser)

			services := &service.Service{Authorization: auth}
			handler := Handler{services}

			r := mux.NewRouter()
			r.HandleFunc("/auth/sign-in", handler.signIn).Methods("POST")

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/auth/sign-in", bytes.NewBufferString(tt.inputBody))
			req.Header.Set("content-type", "application/json")

			r.ServeHTTP(w, req)
			assert.Equal(t, w.Code, tt.expectedStatusCode)
			assert.Equal(t, strings.TrimSpace(w.Body.String()), tt.expectedResponseBody)
		})
	}

}
