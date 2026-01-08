package main

import (
	database "github.com/SnackLog/database-config-lib"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {
	err := database.LoadConfig()
	if err != nil {
		panic(err)
	}

	connectionString := database.GetDatabaseConnectionString()
	db, err := connectDatabase(connectionString)
	defer db.Close()
	if err != nil {
		panic(err)
	}
	err = seedDB(db)
	if err != nil {
		panic(err)
	}
}

func connectDatabase(dbConnectionString string) (*sql.DB,error) {
	db, err := sql.Open("postgres", dbConnectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	return db, nil
}
