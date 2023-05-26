package config

import (
	"fmt"
	"os"

	"github.com/Koakovski/gopportunities/schema"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Database file info
var (
	databaseFolder   = "./db"
	databaseFilename = "main.db"
	databasePath     = fmt.Sprintf("%s/%s", databaseFolder, databaseFilename)
)

func InitializeDatabase() (*gorm.DB, error) {
	logger := GetLogger()

	// Create database file and directory if no existis
	if err := handleDatabaseFileCreation(logger); err != nil {
		return nil, err
	}

	// Create DB and connect
	db, err := gorm.Open(sqlite.Open(databasePath), &gorm.Config{})
	if err != nil {
		logger.Errorf("sqlite opening error: %v", err)
		return nil, err
	}

	// Migrate the Schemas
	if err := db.AutoMigrate(&schema.OpeningSchema{}); err != nil {
		logger.Errorf("database migration error: %v", err)
		return nil, err
	}

	return db, nil
}

func handleDatabaseFileCreation(logger *Logger) error {
	// Check if database file exists
	_, err := os.Stat(databasePath)

	fmt.Println(err)

	if os.IsNotExist(err) {
		logger.Info("database file not found, creating...")
		// Create database file and directory
		if err := os.MkdirAll(databaseFolder, os.ModePerm); err != nil {
			logger.Errorf("database folder creation error: %v", err)
			return err
		}

		file, err := os.Create(databasePath)
		if err != nil {
			logger.Errorf("database file creation error: %v", err)
			return err
		}

		file.Close()
	}

	return nil
}
