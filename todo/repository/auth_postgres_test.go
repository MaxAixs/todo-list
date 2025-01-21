package repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
	"todo-list/todo"
)

func TestAuthRepository_CreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	r := NewAuthRepository(db)

	tests := []struct {
		name           string
		mock           func()
		input          todo.User
		expectedResult uuid.UUID
		expectedErr    bool
	}{
		{

			name: "OK",
			mock: func() {
				expectedID := uuid.MustParse("00000000-0000-0000-0000-000000000001")
				rows := sqlmock.NewRows([]string{"id"}).AddRow(expectedID)
				mock.ExpectQuery("INSERT INTO users").WithArgs(sqlmock.AnyArg(), "UserName", "Email", "Password").WillReturnRows(rows)
			},
			input: todo.User{
				Name:     "UserName",
				Email:    "Email",
				Password: "Password",
			},
			expectedResult: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
		},
		{
			name: "empty fields",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"})
				mock.ExpectQuery("INSERT INTO users").WithArgs(sqlmock.AnyArg(), "UserName", "", "Password").WillReturnRows(rows)
			},
			input: todo.User{
				Name:     "UserName",
				Email:    "",
				Password: "Password",
			},
			expectedErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := r.CreateUser(tt.input)
			if tt.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResult, got)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestAuthRepository_GetUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	r := NewAuthRepository(db)

	type args struct {
		Email    string
		Password string
	}

	tests := []struct {
		name           string
		mock           func()
		input          args
		ExpectedResult todo.User
		ExpectedErr    bool
	}{
		{
			name: "OK",
			mock: func() {
				expectedID := uuid.MustParse("00000000-0000-0000-0000-000000000001")
				rows := sqlmock.NewRows([]string{"id"}).AddRow(expectedID)
				mock.ExpectQuery(regexp.QuoteMeta("SELECT id FROM users WHERE email = $1 AND password_hash = $2")).
					WithArgs("email", "password").WillReturnRows(rows)
			},
			input: args{Email: "email", Password: "password"},
			ExpectedResult: todo.User{
				ID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			},
		},
		{
			name: "not found",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"})
				mock.ExpectQuery(regexp.QuoteMeta("SELECT id FROM users WHERE email = $1 AND password_hash = $2")).
					WithArgs("email", "password").WillReturnRows(rows)
			},
			input:       args{Email: "email", Password: "password"},
			ExpectedErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := r.GetUser(tt.input.Email, tt.input.Password)
			if tt.ExpectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.ExpectedResult, got)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}

}
