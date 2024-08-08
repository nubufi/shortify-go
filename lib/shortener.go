package lib

import (
	"crypto/sha1"
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"
)

type URLShortener struct {
	DB *sql.DB
}

// NewURLShortener creates a new URLShortener instance with the provided database connection.
// Parameters:
//
// - db: The database connection to use for storing and retrieving URLs.
//
// Returns:
//
// - *URLShortener: A new URLShortener instance.
func NewURLShortener(db *sql.DB) *URLShortener {
	return &URLShortener{DB: db}
}

// ShortenURL generates a unique short code for the original URL and stores it in the database.
// Parameters:
//
// - originalURL: The original URL to shorten.
//
// Returns:
//
// - string: The short code generated for the original URL.
// - error: An error if the short code generation or database insertion fails.
func (s *URLShortener) ShortenURL(originalURL string) (string, error) {
	// Generate a unique short code using SHA-1 and Base62 encoding
	hash := sha1.New()
	hash.Write([]byte(originalURL))
	shortCode := base64.StdEncoding.EncodeToString(hash.Sum(nil))[:8]

	// Check if the short code already exists in the database
	existingURL, _ := s.GetOriginalURL(shortCode)
	if existingURL != "" {
		return shortCode, nil // Return existing short code if the URL is already shortened
	}

	// Insert the original URL and short code into the database
	_, err := s.DB.Exec("INSERT INTO urls (original_url, short_code) VALUES ($1, $2)", originalURL, shortCode)
	if err != nil {
		return "", fmt.Errorf("error inserting URL into the database: %w", err)
	}

	return shortCode, nil
}

// GetOriginalURL retrieves the original URL from the database based on the short code.
// Parameters:
//
// - shortCode: The short code to look up in the database.
//
// Returns:
//
// - string: The original URL corresponding to the short code.
// - error: An error if the short code is not found or an error occurs during retrieval.
func (s *URLShortener) GetOriginalURL(shortCode string) (string, error) {
	var originalURL string
	err := s.DB.QueryRow("SELECT original_url FROM urls WHERE short_code = $1", shortCode).Scan(&originalURL)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("short code not found")
		}
		return "", fmt.Errorf("error retrieving original URL: %w", err)
	}

	return originalURL, nil
}
