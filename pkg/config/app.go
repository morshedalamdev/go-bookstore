package config

import (
	"context"
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() error {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		return fmt.Errorf("DATABASE_URL is missing")
	}

	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to the database: %v", err)
	}

	// Get the underlying database/sql connection pool
	sqlDB, err := connection.DB()
	if err != nil {
		return fmt.Errorf("failed to get SQL connection: %v", err)
	}

	// Configure connection pool settings
	sqlDB.SetMaxOpenConns(25)                 // Maximum number of open connections to the database
	sqlDB.SetMaxIdleConns(25)                 // Maximum number of idle connections in the pool
	sqlDB.SetConnMaxLifetime(5 * time.Minute) // Maximum amount of time a connection may be reused

	// Verify database connection
	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()

	if err := sqlDB.PingContext(ctx); err != nil {
		sqlDB.Close() // Close the connection pool if ping fails
		return fmt.Errorf("failed to ping the database: %v", err)
	}

	db = connection
	return nil
}

func GetDB() *gorm.DB {
	return db
}
