package db

import "database/sql"

// CreateTable ensures the URLs table exists
func CreateTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS urls (
id INTEGER PRIMARY KEY AUTOINCREMENT,
original_url TEXT NOT NULL,
short_url TEXT NOT NULL
);`
	_, err := db.Exec(query)
	return err

}

// StoreURL stores the original URL and shortened URL in the database
func StoreURL(db *sql.DB, shortURL string, originalURL string) error {
	query := `INSERT INTO urls (short_url, original_url) VALUES (?, ?)`
	_, err := db.Exec(query, shortURL, originalURL)
	return err
}

// GetOriginalURL retrieves the original URL from the database based on the shortened URL
func GetOriginalURL(db *sql.DB, shortURL string) (string, error) {
	var originalURL string
	query := `SELECT original_url FROM urls WHERE short_url = ? LIMIT 1`
	err := db.QueryRow(query, shortURL).Scan(&originalURL)
	if err != nil {
		return "", nil
	}
	return originalURL, err
}
