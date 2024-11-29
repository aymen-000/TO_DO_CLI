package main

import (
	"log"
)

func main() {
	// Initialize todos
	todos := Todos{}

	// Connect to the database
	db := Connect()

	defer db.Close()

	// Initialize database storage
	storage := NewDBStorage(db)

	// Load todos from the database
	if err := storage.Load(&todos); err != nil {
		log.Fatalf("Failed to load todos: %v", err)
	}

	// Process command-line flags
	cmdFlags := NewCmdFlags()
	cmdFlags.Execute(&todos , storage.db)
}
