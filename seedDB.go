package main

import (
	"database/sql"
	"fmt"
	"os"
)

func seedDB(db *sql.DB) error {
	sqlFileContent, err := readSQLFile("/data/seed.sql")
	if err != nil {
		return fmt.Errorf("failed to read SQL file: %v", err)
	}
	err = executeSQL(sqlFileContent, db)
	if err != nil {
		return fmt.Errorf("failed to execute SQL file: %v", err)
	}
	return nil
}

func readSQLFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("failed to read SQL file: %v", err)
	}
	return string(data), nil
}

func executeSQL(sqlString string, db *sql.DB) error {
	_, err := db.Exec(sqlString)
	if err != nil {
		return fmt.Errorf("failed to execute SQL: %v", err)
	}
	return nil
}
