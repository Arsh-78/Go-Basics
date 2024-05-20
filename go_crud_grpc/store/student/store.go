package store

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

const InsertStudentQuery = `
INSERT INTO studentTwo (name, studentId, class, email, address)
VALUES (?, ?, ?, ?, ?)
`

const ReadStudentQuery = ` SELECT * FROM studentTwo WHERE studentId = ?`

const UpdateStudentQuery = `UPDATE studentTwo SET name = ?, class = ? , email = ? ,address = ? WHERE studentId = ?`

const DeleteStudentQuery = `DELETE FROM studentTwo WHERE studentId = ?`

type Store struct {
	db *sql.DB
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func EstablishDbConnection() (*sql.DB, error) {
	env_err := godotenv.Load()
	if env_err != nil {
		log.Fatal("Error loading .env file")
	}

	// Open up our database connection.
	// I've set up a database on my local machine using mysql and mysql workbench.
	// The database is called test_db

	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	db, err := sql.Open("mysql", connectionString)

	return db, err

}
