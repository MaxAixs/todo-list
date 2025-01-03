package repository

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"todo-list/todo"
)

func TestTodoListRepository_CreateList(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	r := NewTodoListRepository(db)

	type args struct {
		userID uuid.UUID
		todo   todo.TodoList
	}

	tests := []struct {
		name           string
		mock           func()
		input          args
		expectedResult int
		expectedErr    bool
	}{
		{
			name: "OK",
			mock: func() {
				userID := uuid.MustParse("00000000-0000-0000-0000-000000000001")
				mock.ExpectBegin()
				rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
				mock.ExpectQuery("INSERT INTO todo_lists").
					WithArgs("title", "description", true, userID).WillReturnRows(rows)

				mock.ExpectExec("INSERT INTO user_lists").WithArgs(userID, 1).
					WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectCommit()
			},
			input: args{userID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				todo: todo.TodoList{
					Title:       "title",
					Description: "description",
					Public:      true,
				}},
			expectedResult: 1,
		}, {
			name: "empty fields",
			mock: func() {
				mock.ExpectBegin()

				userID := uuid.MustParse("00000000-0000-0000-0000-000000000001")
				rows := sqlmock.NewRows([]string{"id"})
				mock.ExpectQuery("INSERT INTO todo_lists").
					WithArgs("", "description", true, userID).WillReturnRows(rows)

				mock.ExpectRollback()
			},
			input: args{userID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				todo: todo.TodoList{
					Title:       "",
					Description: "description",
					Public:      true,
				}},
			expectedErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := r.CreateList(tt.input.userID, tt.input.todo)
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

func TestTodoListRepository_DeleteListById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	r := NewTodoListRepository(db)

	type args struct {
		userID uuid.UUID
		listID int
	}

	tests := []struct {
		name        string
		mock        func()
		input       args
		expectedErr bool
	}{
		{
			name: "OK",
			mock: func() {
				mock.ExpectExec("DELETE FROM todo_lists").
					WithArgs(1, uuid.MustParse("00000000-0000-0000-0000-000000000001")).
					WillReturnResult(sqlmock.NewResult(0, 1))
			},
			input: args{userID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				listID: 1},
		}, {
			name: "not found",
			mock: func() {
				mock.ExpectExec("DELETE FROM todo_lists").
					WithArgs(404, uuid.MustParse("00000000-0000-0000-0000-000000000001")).WillReturnError(sql.ErrNoRows)
			},
			input: args{listID: 404,
				userID: uuid.MustParse("00000000-0000-0000-0000-000000000001")},
			expectedErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			err := r.DeleteListById(tt.input.userID, tt.input.listID)
			if tt.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}

}

func TestListItemRepository_UpdateItem(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	r := NewTodoListRepository(db)

	type args struct {
		listID int
		userID uuid.UUID
		input  todo.UpdateListInput
	}

	tests := []struct {
		name        string
		mock        func()
		input       args
		expectedErr bool
	}{
		{
			name: "OK",
			mock: func() {
				mock.ExpectExec("UPDATE todo_lists .* SET .* WHERE .*").
					WithArgs("New title", "New description", uuid.MustParse("00000000-0000-0000-0000-000000000001"), 1).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			input: args{listID: 1,
				userID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				input: todo.UpdateListInput{
					Title:       strPtr("New title"),
					Description: strPtr("New description")}},
		},
		{
			name: "OK_withoutDescription",
			mock: func() {
				mock.ExpectExec("UPDATE todo_lists .* SET .* WHERE .*").
					WithArgs("New title", uuid.MustParse("00000000-0000-0000-0000-000000000001"), 1).
					WillReturnResult(sqlmock.NewResult(0, 1))
			},
			input: args{listID: 1,
				userID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				input:  todo.UpdateListInput{Title: strPtr("New title")}},
		},
		{
			name: "Error updating",
			mock: func() {
				mock.ExpectExec("UPDATE todo_lists .* SET .* WHERE .*").
					WithArgs("New Title", uuid.MustParse("00000000-0000-0000-0000-000000000001"), 1).
					WillReturnError(fmt.Errorf("update error"))
			},
			input: args{
				listID: 1,
				userID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				input: todo.UpdateListInput{
					Title: strPtr("New Title"),
				},
			},
			expectedErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			err := r.UpdateList(tt.input.userID, tt.input.listID, tt.input.input)
			if tt.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func strPtr(s string) *string {
	return &s
}
