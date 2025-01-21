package repository

import (
	"fmt"
	"github.com/google/uuid"
	"strings"
	"todo-list/todo"
)

func buildListUpdateSet(userID uuid.UUID, listID int, input todo.UpdateListInput) (string, []interface{}) {
	var (
		setValues []string
		args      []interface{}
		argId     = 1
	)

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title = $%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description = $%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	if input.Public != nil {
		setValues = append(setValues, fmt.Sprintf("public = $%d", argId))
		args = append(args, *input.Public)
		argId++
	}

	setValues = append(setValues, "updated_at = NOW()")
	args = append(args, userID, listID)

	return strings.Join(setValues, ", "), args
}

func buildItemUpdateSet(userID uuid.UUID, itemID int, input todo.UpdateItemInput) (string, []interface{}) {
	var (
		setValues []string
		args      []interface{}
		argId     = 1
	)

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description = $%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	if input.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done = $%d", argId))
		args = append(args, *input.Done)
		argId++
	}

	if input.DueDate != nil {
		setValues = append(setValues, fmt.Sprintf("due_date = $%d", argId))
		args = append(args, *input.DueDate)
		argId++
	}

	if input.Priority != nil {
		setValues = append(setValues, fmt.Sprintf("priority = $%d", argId))
		args = append(args, *input.Priority)
		argId++
	}

	setValues = append(setValues, "updated_at=NOW()")
	args = append(args, userID, itemID)

	setQuery := strings.Join(setValues, ", ")

	return setQuery, args
}
