package main

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (todos *Todos) add(title string, db *sql.DB) error {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}

	// Handle the nil CompletedAt when inserting
	result, err := db.Exec("INSERT INTO todo (title, completed, created_at, completed_at) VALUES (?, ?, ?, ?)",
		todo.Title,
		todo.Completed,
		todo.CreatedAt,
		todo.CompletedAt)

	if err != nil {
		return fmt.Errorf("addTodo: %v", err)
	}

	// Check LastInsertId and handle potential errors
	_, err = result.LastInsertId()
	if err != nil {
		return fmt.Errorf("getting last insert ID: %v", err)
	}

	// Optionally, append the new todo to the slice
	*todos = append(*todos, todo)

	return nil
}

func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		return errors.New("invalid index")
	}
	return nil
}

func (todos *Todos) delete(index int, db *sql.DB) error {
	// First, check if the todo exists
	row := db.QueryRow("SELECT id FROM todo WHERE id = ?", index)
	var existingID int
	if err := row.Scan(&existingID); err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("todo with id %d does not exist", index)
		}
		return fmt.Errorf("error checking todo existence: %v", err)
	}

	// Prepare delete query
	query := "DELETE FROM todo WHERE id = ?"

	// Execute delete
	result, err := db.Exec(query, index)
	if err != nil {
		return fmt.Errorf("failed to execute delete query: %v", err)
	}

	// Check the number of rows affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get affected rows: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows were deleted for todo with id %d", index)
	}

	for i, _ := range *todos {
		if i == index {
			*todos = append((*todos)[:i], (*todos)[i+1:]...)
			break
		}
	}

	return nil
}

func (todos *Todos) toggle(index int, db *sql.DB) error {
	// Query to select the specific todo item
	row := db.QueryRow("SELECT id, title, completed, created_at, completed_at FROM todo WHERE id = ?", index)

	var id int
	var title string
	var completed bool
	var createdAtBytes, completedAtBytes []byte

	// Scan the row
	if err := row.Scan(&id, &title, &completed, &createdAtBytes, &completedAtBytes); err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("row %d: no such todo", index)
		}
		return fmt.Errorf("error fetching todo by id %d: %v", index, err)
	}

	// Prepare update query
	query := `UPDATE todo SET completed = ?, completed_at = ? WHERE id = ?`

	// If currently not completed, set completed_at, otherwise set to NULL
	var completedAt interface{}
	if !completed {
		completedAt = time.Now()
	} else {
		completedAt = nil
	}

	// Execute update
	result, err := db.Exec(query, !completed, completedAt, index)
	if err != nil {
		return fmt.Errorf("failed to execute update query: %v", err)
	}

	// Check the number of rows affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get affected rows: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows were updated for todo with id %d", index)
	}

	return nil
}

func (todos *Todos) edit(index int, title string, db *sql.DB) error {

	query := `UPDATE todo SET title = ? WHERE id = ?`

	result, err := db.Exec(query, title, index)

	if err != nil {
		fmt.Errorf("Update title %v", err)
		return err
	}

	// Check the number of rows affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Errorf("failed to get affected rows: %v", err)
		return err
	}

	if rowsAffected == 0 {
		fmt.Errorf("no rows were updated for todo with id %d", index)
		return err
	}

	return nil
}

func (todos *Todos) print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Completed", "CreatedAt", "CompletedAt")
	for index, t := range *todos {
		completed := "❎"
		completedAt := ""

		if t.Completed {
			completed = "✅"
			if t.CompletedAt != nil {
				completedAt = t.CompletedAt.Format(time.RFC1123)
			}
		}

		table.AddRow(strconv.Itoa(index), t.Title, completed, t.CreatedAt.Format(time.RFC1123), completedAt)
	}
	table.Render()
}
