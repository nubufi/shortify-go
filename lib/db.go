package lib

import (
	"database/sql"
	"errors"
	"os"

	_ "github.com/lib/pq"
)

// CreateDb creates a new database connection using the DATABASE_URL environment variable.
// Returns:
// - *sql.DB: The database connection.
func CreateDb() (*sql.DB, error) {
	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		return nil, errors.New("DATABASE_URL is not set")
	}

	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		return nil, err
	}

	err = createTable(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// createTable creates the 'urls' table in the database if it does not exist.
// Parameters:
// - db: *sql.DB: The database connection.
// Returns:
// - error: An error if the table creation fails.
func createTable(db *sql.DB) error {
	createTableSQL := `
    CREATE TABLE IF NOT EXISTS urls (
        id SERIAL PRIMARY KEY,
        original_url TEXT NOT NULL,
        short_code TEXT NOT NULL UNIQUE,
        created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
    );`
	_, err := db.Exec(createTableSQL)

	return err
}
