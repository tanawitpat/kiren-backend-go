package app

import (
	"database/sql"
	"strconv"

	// MySQL driver
	_ "github.com/go-sql-driver/mysql"
)

// Query creates a MySQL connection and execute the provided SQL script using Query function.
func Query(query string) (*sql.Rows, error) {
	// Open up database connection.
	db, err := initMySQLConnection()
	if err != nil {
		return nil, err
	}

	sqlRows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	// Defer the close till after the main function has finished executing
	defer db.Close()
	return sqlRows, nil
}

// QueryRow creates a MySQL connection and execute the provided SQL script using QueryRow function.
func QueryRow(query string) (*sql.Row, error) {
	// Open up database connection.
	db, err := initMySQLConnection()
	if err != nil {
		return nil, err
	}

	sqlRow := db.QueryRow(query)

	// Defer the close till after the main function has finished executing
	defer db.Close()
	return sqlRow, nil
}

func initMySQLConnection() (*sql.DB, error) {
	dbDriver := "mysql"
	dbUser := CFG.DB.Username
	dbPass := CFG.DB.Password
	dbName := CFG.DB.DBName
	dbHost := CFG.DB.Host
	dbPort := strconv.Itoa(CFG.DB.Port)
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbHost+":"+dbPort+")/"+dbName)
	return db, err
}
