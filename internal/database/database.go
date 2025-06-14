package database

import (
	"database/sql"
	"log"
	"strconv"
	"time"

	"github.com/chandrakantG/lumen/config"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitializeDB() {
	dbConfig := config.GetDBConfig()

	// Establish a connection to the PostgreSQL database
	db, err := sql.Open("postgres", dbConfig.DSN())
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	// Check if the connection is successful
	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging the database: ", err)
	}
	DB = db

	// Set connection pool settings
	maxIdleConns, _ := strconv.Atoi(dbConfig.MaxIdleConns)
	maxOpenConns, _ := strconv.Atoi(dbConfig.MaxOpenConns)
	connMaxLifetime, _ := strconv.Atoi(dbConfig.ConnMaxLifetime)

	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxLifetime(time.Duration(connMaxLifetime) * time.Millisecond)
	db.SetConnMaxIdleTime(5 * time.Minute)
}

// CloseDB closes the database connection
func CloseDB() {
	if DB != nil {
		err := DB.Close()
		if err != nil {
			log.Printf("Error closing the database connection: %v", err)
		} else {
			log.Println("Database connection closed successfully")
		}
	} else {
		log.Println("No database connection to close")
	}
}

// GetDB returns the database connection
func GetDB() *sql.DB {
	if DB == nil {
		log.Fatal("Database connection is not initialized")
	}
	return DB
}
