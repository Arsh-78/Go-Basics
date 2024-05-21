package dbMigration

import (
	"database/sql"
	"fmt"
	"log"
)

func Migrations(db *sql.DB, err error) {
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS test_db")
	if err != nil {
		log.Fatalf("Error creating database: %v", err)
	}
	fmt.Println("Database created successfully")

	// Select the database
	_, err = db.Exec("USE test_db")
	if err != nil {
		log.Fatalf("Error selecting database: %v", err)
	}

	// Create a table
	tableCreationQuery := `CREATE TABLE IF NOT EXISTS studentTwo (
		name VARCHAR(50) NOT NULL,
		studentId VARCHAR(50) NOT NULL,
		class INT ,
		email VARCHAR(50) NOT NULL,
		address VARCHAR(50) NOT NULL,
		PRIMARY KEY (StudentId)
	)`

	_, err = db.Exec(tableCreationQuery)
	if err != nil {
		log.Fatalf("Error creating table: %v", err)
	}
	fmt.Println("Table created successfully")

}
