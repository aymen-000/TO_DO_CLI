package main

import (
	"database/sql"
	"fmt"
	"time"
)

// DBStorage holds the database connection
type DBStorage struct {
	db *sql.DB
}

// NewDBStorage initializes a new DBStorage instance
func NewDBStorage(db *sql.DB) *DBStorage {
	return &DBStorage{db: db}
}


// Load retrieves all todos from the database
func (storage *DBStorage) Load(todos *Todos) error {
    query := "SELECT title, completed, created_at, completed_at FROM todo"
    rows, err := storage.db.Query(query)
    if err != nil {
        return fmt.Errorf("failed to query todos: %v", err)
    }
    defer rows.Close()

    for rows.Next() {
        var todo Todo
        var title string
        var completed bool
        var createdAtBytes, completedAtBytes []byte

        // Scan the columns into appropriate types
        if err := rows.Scan(&title, &completed, &createdAtBytes, &completedAtBytes); err != nil {
            return fmt.Errorf("failed to scan row: %v", err)
        }

        todo.Title = title
        todo.Completed = completed

        // Parse created_at timestamp
        if createdAtBytes != nil {
            createdAt, err :=time.Parse("2006-01-02 15:04:05", string(createdAtBytes))
            if err != nil {
                return fmt.Errorf("failed to parse created_at: %v", err)
            }
            todo.CreatedAt = createdAt
        }

        // Parse completed_at timestamp
        if completedAtBytes != nil {
            completedAt, err := time.Parse("2006-01-02 15:04:05", string(completedAtBytes))
            if err != nil {
                return fmt.Errorf("failed to parse completed_at: %v", err)
            }
            todo.CompletedAt = &completedAt
        }

        *todos = append(*todos, todo)
    }

    return nil
}