package config

import (
	"fmt"
	"net/url"
	"os"
)

type Database struct {
	Host            string
	Port            string
	User            string
	Password        string
	DbName          string
	SSLMode         string
	MaxIdleConns    string
	MaxOpenConns    string
	ConnMaxLifetime string
}

var dbConfig Database

func LoadDBConfig() Database {
	dbConfig = Database{
		Host:            os.Getenv("DB_HOST"),
		Port:            os.Getenv("DB_PORT"),
		User:            os.Getenv("DB_USER"),
		Password:        os.Getenv("DB_PASSWORD"),
		DbName:          os.Getenv("DB_NAME"),
		SSLMode:         os.Getenv("DB_SSL_MODE"),
		MaxIdleConns:    os.Getenv("DB_MAX_IDLE_CONNS"),
		MaxOpenConns:    os.Getenv("DB_MAX_OPEN_CONNS"),
		ConnMaxLifetime: os.Getenv("DB_CONN_MAX_LIFETIME"),
	}
	return dbConfig
}

func GetDBConfig() Database {
	return dbConfig
}

func (dbConfig Database) DSN() string {
	// Use URL encoding for username and password to handle special characters correctly
	escapedUser := url.QueryEscape(dbConfig.User)
	escapedPassword := url.QueryEscape(dbConfig.Password)
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		escapedUser, escapedPassword, dbConfig.Host, dbConfig.Port, dbConfig.DbName, dbConfig.SSLMode)
}
