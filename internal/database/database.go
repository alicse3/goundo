package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

const (
	statusBackedUp = "BACKED_UP"
)

// Backup is a struct which represents the backups table schema.
type Backup struct {
	ID        int
	SrcPath   string
	DstPath   string
	Status    string
	CreatedAt string
}

// DBHandler manages the SQLite database connection.
type DBHandler struct {
	db *sql.DB
}

// NewDBHandler initializes the database and returns a handler.
func NewDBHandler(dbPath string) (*DBHandler, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	// Initialise the database handler
	dbHandler := &DBHandler{db}

	// Create backups table
	if err := dbHandler.createTable(); err != nil {
		return nil, err
	}

	return dbHandler, nil
}

// createTable creates the necessary table if it doesn't exist.
func (handler *DBHandler) createTable() error {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS backups (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		src_path TEXT NOT NULL,
		dst_path TEXT NOT NULL,
		status VARCHAR(15) NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := handler.db.Exec(createTableSQL)
	if err != nil {
		return err
	}

	return nil
}

// Insert inserts the given data into the backups table.
func (handler *DBHandler) Insert(srcPath, dstPath string) error {
	insertSQL := `INSERT INTO backups (src_path, dst_path, status) VALUES (?, ?, ?)`

	_, err := handler.db.Exec(insertSQL, srcPath, dstPath, statusBackedUp)
	if err != nil {
		return err
	}

	return nil
}

// GetAll returns data from the backups table.
func (handler *DBHandler) GetAll() ([]Backup, error) {
	selectSQL := `SELECT * FROM backups`

	rows, err := handler.db.Query(selectSQL)
	if err != nil {
		return nil, err
	}

	var backups []Backup
	for rows.Next() {
		var backup Backup
		if err := rows.Scan(&backup.ID, &backup.SrcPath, &backup.DstPath, &backup.Status, &backup.CreatedAt); err != nil {
			return nil, err
		}

		backups = append(backups, backup)
	}

	return backups, nil
}
